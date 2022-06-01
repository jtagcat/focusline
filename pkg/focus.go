package pkg

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

var ErrModeInvalid = fmt.Errorf("runtime enum: FocusReader: mode is invalid")

// Focuses lines of text.
//
// mode: -1: focus all lines, prefer left
// mode:  0: focus all lines, prefer right
// mode:  1: left-align all lines, focus last line, prefer left
// mode:  2: right-align all lines, focus last line, prefer right
// TODO: StringsSplitEveryN: try to split at spaces
func FocusReader(r io.Reader, focus, lenTarget uint, mode int) (w []string, _ error) {
	var prefRight, focusAll bool

	switch mode {
	case -1, 1:
		prefRight = false
	case 0, 2:
		prefRight = true
	default:
		return w, fmt.Errorf("%d: %w", mode, ErrModeInvalid)
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
		prefRight = !prefRight // mirror last for better astetiks
		w[last] = Focus(strings.TrimSpace(w[last]), focus, lenTarget, prefRight)
	}

	if err := scanner.Err(); err != nil {
		return w, fmt.Errorf("while reading input: %w", err)
	}
	return w, nil
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
	tlen := utf8.RuneCountInString(text)

	textBeforeFocus := (tlen / 2) + 1
	if (tlen%2) == 0 && !preferRight {
		textBeforeFocus--
	}

	whiteLen := int(focus) - textBeforeFocus

	if whiteLen > 0 {
		if lenTarget != 0 && (tlen+whiteLen) >= int(lenTarget) {
			whiteLen = int(lenTarget) - tlen
		}
		return strings.Repeat(" ", whiteLen) + text
	}

	return text
}
