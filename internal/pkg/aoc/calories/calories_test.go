package calories

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestCalories(t *testing.T) {
	is := is.New(t)
	rd := bytes.NewReader([]byte(input))

	result, err := Solve(rd)

	is.NoErr(err)
	is.Equal(result, "24000")
}

const input string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
