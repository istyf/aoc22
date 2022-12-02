package rps

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestRockPaperScissorsPartOne(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartOne(rd)

	is.NoErr(err)
	is.Equal(result, "15")
}

func TestRockPaperScissorsPartTwo(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartTwo(rd)

	is.NoErr(err)
	is.Equal(result, "12")
}

const input string = `A Y
B X
C Z`
