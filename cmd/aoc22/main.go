package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/istyf/aoc22/internal/pkg/aoc/campcleanup"
)

func main() {

	inputfile, err := os.Open("./input.dat")
	if err != nil {
		panic("unable to open input file")
	}

	defer inputfile.Close()

	data, err := io.ReadAll(inputfile)
	if err != nil {
		panic(err.Error())
	}

	result1, err := campcleanup.PartOne(bytes.NewBuffer(data))
	if err != nil {
		panic(err.Error())
	}

	result2, err := campcleanup.PartTwo(bytes.NewBuffer(data))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Result; part one = %s, part two = %s", result1, result2)
}
