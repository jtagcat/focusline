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
			if fWrap <= focus {
				return fmt.Errorf("flag \"wrap\" (%d) must be a larger value than focus (%d)", fWrap, focus)
			} // TODO: errors might be better upstream?

			out, err := FocusReader(os.Stdin, uint(focus), uint(fWrap), -2)
			fmt.Println(out)
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
func FocusReader(r io.Reader, focus, wrap uint, mode int) (w []string, _ error) {
	var prefBool, focusAll bool
	switch mode {
	case -1, 1:
		prefBool = true
	case 0, 2:
		prefBool = false
	default:
		return w, fmt.Errorf("runtime enum: FocusReader: mode %d is invalid", mode)
	}
	if mode == -1 || mode == 0 {
		focusAll = true
	}

	scanner := bufio.NewScanner(r)
	for i := 1; scanner.Scan(); i++ {

		line := StringsSplitEveryN(scanner.Text(), wrap)
		for _, l := range line {
			if focusAll {
				w = append(w, Focus(l, focus, wrap, prefBool))
			} else {
				w = append(w, Align(l, wrap, prefBool))
			}
		}
		if !focusAll {
			last := len(w) - 1
			w[last] = Focus(w[last], focus, wrap, prefBool)
		}
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
// maxLen is a soft limit for whitespaces  |focus| |focus|
// "      hello" when forced unsymetrical, |  on | | on  |
// "       tere" text shifted to right or  | left| |rigt |
// target:  ^^    left (preferLeft=true)   |  /  | |  \  |
//          LR                             | /   | |   \ |
func Focus(text string, focus, maxLen uint, preferLeft bool) string {
	// TODO: check input sanity; is maxLen right place?
	// TODO: swap right/left
	textBeforeFocus := (len(text) / 2) + 1
	if (len(text)%2) == 0 && preferLeft {
		textBeforeFocus -= 1
	}

	whitespace := int(focus) - textBeforeFocus

	if whitespace > 0 {
		if maxLen > 0 && (len(text)+whitespace) >= int(maxLen) {
			whitespace = int(maxLen) - len(text)
		}
		return strings.Repeat(" ", whitespace) + text
	}
	return text
}

func Align(s string, maxLen uint, toRight bool) string {
	whitelen := int(maxLen) - len(s)
	if whitelen < 1 {
		return s
	}
	whitespace := strings.Repeat(" ", whitelen)

	if toRight {
		return whitespace + s
	}
	return s + whitespace
}
