package main

import (
	"fmt"
	"os"

	"github.com/istyf/aoc22/internal/pkg/aoc/calories"
)

func main() {

	inputfile, err := os.Open("./input.dat")
	if err != nil {
		panic("unable to open input file")
	}

	defer inputfile.Close()

	result, err := calories.Solve(inputfile)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The result is", result)
}
