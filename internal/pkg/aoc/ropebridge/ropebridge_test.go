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

const input string = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
