package supplystacks

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestSupplyStacksPartOne(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartOne(rd)

	is.NoErr(err)
	is.Equal(result, "CMZ")
}

func TestSupplyStacksPartTwo(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := PartTwo(rd)

	is.NoErr(err)
	is.Equal(result, "MCD")
}

const input string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
