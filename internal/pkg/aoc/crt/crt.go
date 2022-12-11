package crt

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartOne(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	accumulator := 0
	clock := 0
	register := 1

	accumulate := func(clock, accumulator, register int) int {
		if clock == 20 || (clock-20)%40 == 0 {
			return accumulator + (clock * register)
		}

		return accumulator
	}

	for scanner.Scan() {
		instruction := scanner.Text()
		clock++

		if instruction == "noop" {
			accumulator = accumulate(clock, accumulator, register)
		} else if strings.HasPrefix(instruction, "addx") {
			accumulator = accumulate(clock, accumulator, register)
			accumulator = accumulate(clock+1, accumulator, register)
			clock++

			amount, _ := strconv.ParseInt(strings.Split(instruction, " ")[1], 10, 64)
			register += int(amount)
		} else {
			panic("cpu halted; illegal instruction")
		}
	}

	return strconv.FormatInt(int64(accumulator), 10), nil
}

func PartTwo(input io.Reader) (string, error) {
	return "not implemented", nil
}
