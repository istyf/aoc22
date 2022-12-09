package nospaceleft

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestPartOne(t *testing.T) {
	is := is.New(t)

	result, err := PartOne(bytes.NewBuffer([]byte(input)))
	is.NoErr(err)
	is.Equal(result, "95437")
}

func TestPartTwo(t *testing.T) {
	is := is.New(t)

	result, err := PartTwo(bytes.NewBuffer([]byte(input)))
	is.NoErr(err)
	is.Equal(result, "24933642")
}

const input string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`