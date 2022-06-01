package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jtagcat/focusline/pkg"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:    "focusline",
	Version: "1.0.0",
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
		&cli.BoolFlag{
			Name:     "right",
			Aliases:  []string{"r"},
			Usage:    "prefer right instead of left",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "last",
			Aliases:  []string{"l"},
			Usage:    "focus only last, align others to wrap",
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

		var mode int
		if c.Bool("last") {
			mode++
		}
		if c.Bool("right") {
			mode++
		}

		// TODO: errors might be better upstream?
		fWrap := c.Int("wrap")
		if fWrap > 0 && fWrap <= focus {
			return fmt.Errorf("flag \"wrap\" (%d) must be a larger value than focus (%d)", fWrap, focus)
		}

		out, err := pkg.FocusReader(os.Stdin, uint(focus), uint(fWrap), mode)
		for _, o := range out {
			fmt.Println(o)
		}
		return err
	},
}
