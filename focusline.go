package main

import (
	"bufio"
	"fmt"
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
			&cli.IntFlag{
				Name:     "first",
				Aliases:  []string{"f"},
				Usage:    "character limit on first line (focus is relative to 1st)",
				Required: false,
			},
			&cli.IntFlag{
				Name:     "shift",
				Aliases:  []string{"start", "s"},
				Usage:    "shift focus by n on first line",
				Required: false,
			},
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

			fShift := c.Int("shift")
			fFirst := c.Int("first")
			fWrap := c.Int("wrap")
			if fShift > 0 && fWrap+fFirst == 0 {
				return fmt.Errorf("using shift without wrapping is ineffective;\n" +
					"to use shift, use a character limit (wrap or first)")
			}

			// scanner, once := bufio.NewScanner(os.Stdin), false // TODO:
			xyz, _ := os.Open("test")
			scanner, once := bufio.NewScanner(xyz), false
			for i := 1; scanner.Scan(); i++ {
				txt := scanner.Text()
				var line []string

				doFirstWrap := i == 1 && fFirst != 0
				if doFirstWrap {
					line = append(line, txt[:fFirst])
					txt = txt[fFirst:]
				}

				if fWrap != 0 {
					line = append(line, StringsSplitEveryN(txt, fWrap)...)
				}
				for i, l := range line {
					if i+1 == len(line) {
						output(FocusText(l, focus))
					} else {
						output(l)
					}

					if i == 1 && fShift != 0 && !once {
						focus = focus - fShift
						once = true // sync.Once got way too complex and unreadable
					}
				}
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("while reading standard input: %w", err)
			}

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
	}
}

func output(line string) {
	fmt.Println(line)
}

func StringsSplitEveryN(s string, n int) (o []string) {
	for i := 0; i < len(s); i += n {
		if n > len(s)-i {
			// last (unfinished) piece
			n = len(s) - i
		}
		o = append(o, s[i:i+n])
	}
	return
}

// Centers string by prepending whitespace
//
// "      hello" when forced unsymetrical,
// "       tere" text shifted to right
// target:  ^
func FocusText(text string, focus int) string {
	textBeforeFocus := (len(text) / 2) + (len(text) % 2)
	whitespace := focus - textBeforeFocus

	if whitespace > 0 {
		return strings.Repeat(" ", whitespace) + text
	}
	return text
}
