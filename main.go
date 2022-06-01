package main

import (
	"fmt"
	"os"

	"github.com/jtagcat/focusline/cmd"
)

func main() {
	err := cmd.App.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
}
