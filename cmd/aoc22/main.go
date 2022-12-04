package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/istyf/aoc22/internal/pkg/aoc/calories"
	"github.com/istyf/aoc22/internal/pkg/aoc/campcleanup"
	"github.com/istyf/aoc22/internal/pkg/aoc/rps"
	"github.com/istyf/aoc22/internal/pkg/aoc/rucksack"
)

var day string
var inputFilePath string

type AoCSolver func(io.Reader) (string, error)

func main() {

	flag.StringVar(&day, "day", "4", "")
	flag.StringVar(&inputFilePath, "input", "./input.dat", "The path to the inout file")
	flag.Parse()

	solutions := map[string][]AoCSolver{
		"1": {calories.PartOne, calories.PartTwo},
		"2": {rps.PartOne, rps.PartTwo},
		"3": {rucksack.PartOne, rucksack.PartTwo},
		"4": {campcleanup.PartOne, campcleanup.PartTwo},
	}

	inputfile, err := os.Open(inputFilePath)
	if err != nil {
		panic("unable to open input file")
	}

	defer inputfile.Close()

	data, err := io.ReadAll(inputfile)
	if err != nil {
		panic(err.Error())
	}

	solvers, ok := solutions[day]
	if !ok {
		panic("no solution found for that day")
	}

	for part, solve := range solvers {
		result, err := solve(bytes.NewBuffer(data))
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Result for day %s part %d is %s\n", day, part+1, result)
	}
}
