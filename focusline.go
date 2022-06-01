package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "focusline",
		Version: "0.1.0",
		Authors: []*cli.Author{{Name: "jtagcat"}}, // TODO: Email: ""

		Description: "Center text aiming at nth character.",
		Usage:       "focusline <focus> <stdin",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "wrap",
				Aliases:  []string{"w"},
				Usage:    "per-line character limit",
				Required: false,
			},
			// center; except last: left-align, right-align
		},
		EnableBashCompletion: true,

		Action: func(c *cli.Context) error {
			args := c.Args()
			if args.Len() != 1 {
				return fmt.Errorf("expected exactly 1 argument, got %d", args.Len())
			}

			focusStr := args.First()
			focus, err := strconv.Atoi(focusStr)
			if err != nil {
				return fmt.Errorf("argument focusChar must be an integer")
			}

			fWrap := c.Int("wrap")
			if fWrap > 0 && fWrap <= focus {
				return fmt.Errorf("flag \"wrap\" (%d) must be a larger value than focus (%d)", fWrap, focus)
			} // TODO: errors might be better upstream?

			// out, err := FocusReader(os.Stdin, uint(focus), uint(fWrap), 2)
			r, _ := os.Open("test")
			out, err := FocusReader(r, uint(focus), uint(fWrap), 2)
			for _, o := range out {
				fmt.Println(o)
			}
			return err
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
}

// TODO:
//
// mode: -1: focus all lines, prefer left
// mode:  0: focus all lines, prefer right
// mode:  1: left-align all lines, focus last line, prefer left
// mode:  2: right-align all lines, focus last line, prefer right
func FocusReader(r io.Reader, focus, lenTarget uint, mode int) (w []string, _ error) {
	var prefRight, focusAll bool
	switch mode {
	case -1, 1:
		prefRight = false
	case 0, 2:
		prefRight = true
	default:
		return w, fmt.Errorf("runtime enum: FocusReader: mode %d is invalid", mode)
	}
	if mode == -1 || mode == 0 {
		focusAll = true
	}
	if lenTarget != 0 && focus > lenTarget {
		// TODO: or upstream?
	}

	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {
		line := StringsSplitEveryN(scanner.Text(), lenTarget)
		for _, l := range line {
			if focusAll {
				w = append(w, Focus(l, focus, lenTarget, prefRight))
			} else {
				if prefRight {
					w = append(w, AlignRight(l, lenTarget))
					continue
				}
				w = append(w, l)
			}
		}
	}
	if !focusAll {
		last := len(w) - 1
		w[last] = Focus(strings.TrimSpace(w[last]), focus, lenTarget, prefRight)
	}

	if err := scanner.Err(); err != nil {
		return w, fmt.Errorf("while reading input: %w", err)
	}
	return w, nil
}

// Splits string to string slice after every n characters.
func StringsSplitEveryN(s string, n uint) (o []string) {
	if n == 0 {
		return []string{s}
	}
	u := int(n)
	for i := 0; i < len(s); i += u {
		if u > len(s)-i {
			// last (unfinished) piece
			u = len(s) - i
		}
		o = append(o, s[i:i+u])
	}
	return
}

// Centers string by prepending whitespace |  ↓  | |  ↓  |
// lenTarget: soft limit for prewhitespace |focus| |focus|
// "      hello" when forced unsymetrical, |  on | | on  |
// "       tere" text shifted to right or  | left| |rigt |
// target:  ^^    left (preferLeft=true)   |  /  | |  \  |
//          LR                             | /   | |   \ |
func Focus(text string, focus, lenTarget uint, preferRight bool) string {
	if lenTarget != 0 && focus > lenTarget {
		// lenTarget unreasonable, err?
	}

	textBeforeFocus := (len(text) / 2) + 1
	if (len(text)%2) == 0 && !preferRight {
		textBeforeFocus -= 1
	}

	whiteLen := int(focus) - textBeforeFocus

	if whiteLen > 0 {
		if lenTarget != 0 && (len(text)+whiteLen) >= int(lenTarget) {
			whiteLen = int(lenTarget) - len(text)
		}
		return strings.Repeat(" ", whiteLen) + text
	}
	return text
}

func AlignRight(s string, lenTarget uint) string {
	whiteLen := int(lenTarget) - len(s)
	if whiteLen < 1 {
		return s
	}

	return strings.Repeat(" ", whiteLen) + s
}
