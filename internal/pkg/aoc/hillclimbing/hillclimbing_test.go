package hillclimbing

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestHillClimbingPartOne(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartOne(rd)

	is.NoErr(err)
	is.Equal(result, "31")
}

func TestHillClimbingPartTwo(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartTwo(rd)

	is.NoErr(err)
	is.Equal(result, "29")
}

const input string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
