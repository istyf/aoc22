package treetoptreehouse

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestPartOne(t *testing.T) {
	is := is.New(t)

	result, err := PartOne(bytes.NewBuffer([]byte(input)))
	is.NoErr(err)
	is.Equal(result, "21")
}

func TestPartTwo(t *testing.T) {
	is := is.New(t)

	result, err := PartTwo(bytes.NewBuffer([]byte(input)))
	is.NoErr(err)
	is.Equal(result, "8")
}

const input string = `30373
25512
65332
33549
35390`
