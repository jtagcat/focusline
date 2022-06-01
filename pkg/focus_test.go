package pkg_test

import (
	"strings"
	"testing"

	"github.com/jtagcat/focusline/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFocusNothing(t *testing.T) {
	t.Parallel()
	for _, x := range []bool{false, true} {
		assert.Equal(t, "123", pkg.Focus("123", 2, 0, x))
	}
}

func TestFocusEqual(t *testing.T) {
	t.Parallel()
	for _, x := range []bool{false, true} {
		assert.Equal(t, " 123", pkg.Focus("123", 3, 0, x))
		assert.Equal(t, "  x", pkg.Focus("x", 3, 0, x))
	}
}

func TestFocusLeft(t *testing.T) {
	t.Parallel()
	assert.Equal(t, " xx", pkg.Focus("xx", 2, 0, false))
}

func TestFocusRight(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "xx", pkg.Focus("xx", 2, 0, true))
}

func TestAlignRight(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "     hello", pkg.AlignRight("hello", 10))
}

func TestFocusLineM1_5Wrap0(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"12345\n" +
		"1234\n" +
		"123456\n")
	out, err := pkg.FocusReader(in, 3, 0, -1)
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
	out, err := pkg.FocusReader(in, 4, 6, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"   0",
		"123456",
		"123456", "   78",
	},
		out)
}

func TestFocusLine0_5Wrap0(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"12345\n" +
		"1234\n" +
		"123456\n")
	out, err := pkg.FocusReader(in, 3, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"  0",
		"12345",
		"1234",
		"123456",
	},
		out)
}

func TestFocusLine0_4Wrap6(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"123456\n" +
		"12345678\n")
	out, err := pkg.FocusReader(in, 4, 6, 0)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"   0",
		"123456",
		"123456", "  78",
	},
		out)
}

func TestFocusLine1_5Wrap0(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"12345\n" +
		"1234\n" +
		"123456\n")
	out, err := pkg.FocusReader(in, 3, 0, 1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"0",
		"12345",
		"1234",
		"123456",
	},
		out)
}

func TestFocusLine1_4Wrap6(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"123456\n" +
		"12345678\n")
	out, err := pkg.FocusReader(in, 4, 6, 1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"0",
		"123456",
		"123456", "  78",
	},
		out)
}

func TestFocusLine2_5Wrap0(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"12345\n" +
		"1234\n" +
		"123456\n" +
		"0\n")
	out, err := pkg.FocusReader(in, 3, 0, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"0",
		"12345",
		"1234",
		"123456",
		"  0",
	},
		out)
}

func TestFocusLine2_4Wrap6(t *testing.T) {
	t.Parallel()

	in := strings.NewReader("0\n" +
		"123456\n" +
		"12345678\n" +
		"0\n")
	out, err := pkg.FocusReader(in, 4, 6, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"     0",
		"123456",
		"123456", "    78",
		"   0",
	},
		out)
}

//

func TestReadmeCenter(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"lorem ipsum dolor\n")
	out, err := pkg.FocusReader(in, 16, 24, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"       lorem ipsum dolor",
	}, out)
}

func TestReadmeEOLine(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"lorem ipsum dolor sit\n")
	out, err := pkg.FocusReader(in, 16, 24, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"   lorem ipsum dolor sit",
	}, out)
}

func TestReadmeFocusLeft(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"focus\non\nleft\n")
	out, err := pkg.FocusReader(in, 3, 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"focus", "  on", " left",
	}, out)
}

func TestReadmeFocusRight(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"focus\non\nrigt\n")
	out, err := pkg.FocusReader(in, 3, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"focus", " on", "rigt",
	}, out)
}

func TestReadmeFocusCenter(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"focus\nand\nbreathe\n")
	out, err := pkg.FocusReader(in, 4, 0, -1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		" focus", "  and", "breathe",
	}, out)
}

func TestReadmeAlignLeft(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"iseenda eest on\n" +
			"kõige raskem\n" +
			"põgeneda\n")
	out, err := pkg.FocusReader(in, 10, 17, 1)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"iseenda eest on",
		"kõige raskem",
		"     põgeneda",
	}, out)
}

func TestReadmeAlignRight(t *testing.T) {
	t.Parallel()

	in := strings.NewReader(
		"iseenda eest on\n" +
			"kõige raskem\n" +
			"põgeneda\n")
	out, err := pkg.FocusReader(in, 12, 17, 2)
	assert.Nil(t, err)
	assert.Equal(t, []string{
		"  iseenda eest on",
		"     kõige raskem",
		"        põgeneda",
	}, out)
}
