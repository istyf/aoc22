package ropebridge

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestPartOne(t *testing.T) {
	is := is.New(t)

	result, err := PartOne(bytes.NewBuffer([]byte(input)))
	is.NoErr(err)
	is.Equal(result, "13")
}

func TestPartTwo(t *testing.T) {
	is := is.New(t)

	result, err := PartTwo(bytes.NewBuffer([]byte(larger_input)))
	is.NoErr(err)
	is.Equal(result, "36")
}

const input string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

const larger_input string = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
