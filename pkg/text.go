package pkg

import (
	"strings"
	"unicode/utf8"
)

// Splits string to string slice after every n characters
func StringsSplitEveryN(s string, n uint) (o []string) {
	if n == 0 {
		return []string{s}
	}

	u, r := int(n), []rune(s)

	for i := 0; i < len(r); i += u {

		if u > len(r)-i {
			// last (unfinished) piece
			u = len(r) - i
		}

		o = append(o,
			string(r[i:i+u]),
		)
	}
	return o
}

// Align to lenTarget by padding with spaces
func AlignRight(s string, lenTarget uint) string {
	whiteLen := int(lenTarget) - utf8.RuneCountInString(s)
	if whiteLen < 1 {
		return s
	}

	return strings.Repeat(" ", whiteLen) + s
}
