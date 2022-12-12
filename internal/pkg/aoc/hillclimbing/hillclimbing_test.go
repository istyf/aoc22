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

const input string = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`
