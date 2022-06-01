package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFocusNothing(t *testing.T) {
	t.Parallel()
	for _, x := range []bool{false, true} {
		assert.Equal(t, "123", Focus("123", 2, 0, x))
	}
}

func TestFocusEqual(t *testing.T) {
	t.Parallel()
	for _, x := range []bool{false, true} {
		assert.Equal(t, " 123", Focus("123", 3, 0, x))
		assert.Equal(t, "  x", Focus("x", 3, 0, x))
	}
}

func TestFocusLeft(t *testing.T) {
	t.Parallel()
	assert.Equal(t, " xx", Focus("xx", 2, 0, true))
}

func TestFocusRight(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "xx", Focus("xx", 2, 0, false))
}

func TestAlignRight(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "     hello", AlignRight("hello", 10))
}

func TestFocusLineM1_5Wrap0(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"12345\n" +
		"1234\n" +
		"123456\n")
	out, err := FocusReader(in, 3, 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"  0",
		"12345",
		" 1234",
		"123456",
	},
		out)
}

func TestFocusLineM1_4Wrap6(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"123456\n" +
		"12345678\n")
	out, err := FocusReader(in, 4, 6, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"   0",
		"123456",
		"123456", "   78",
	},
		out)
}
