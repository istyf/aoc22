package tuningtrouble

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestPartOne(t *testing.T) {

	tune := func(input string, want string) func(t *testing.T) {
		return func(t *testing.T) {
			is := is.New(t)
			result, err := PartOne(bytes.NewBuffer([]byte(input)))
			is.NoErr(err)
			is.Equal(result, want)
		}
	}

	t.Run("first", tune("mjqjpqmgbljsphdztnvjfqwrcgsmlb", "7"))
	t.Run("second", tune("bvwbjplbgvbhsrlpgdmjqwftvncz", "5"))
	t.Run("third", tune("nppdvjthqldpwncqszvftbrmjlhg", "6"))
	t.Run("fourth", tune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "10"))
	t.Run("fifth", tune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11"))
}
