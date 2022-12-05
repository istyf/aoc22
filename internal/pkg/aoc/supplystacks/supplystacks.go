package supplystacks

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CrateMoverFunc func(quantity int, from, to []rune) ([]rune, []rune)

func PartOne(input io.Reader) (string, error) {
	const OneAtATime int = 1
	CrateMover9000 := func(quantity int, from, to []rune) ([]rune, []rune) {
		return batchMove(quantity, OneAtATime, from, to)
	}
	return findTopOfStacks(input, CrateMover9000), nil
}

func PartTwo(input io.Reader) (string, error) {
	const SingleMove int = 1
	CrateMover9001 := func(batchSize int, from, to []rune) ([]rune, []rune) {
		return batchMove(SingleMove, batchSize, from, to)
	}
	return findTopOfStacks(input, CrateMover9001), nil
}

func findTopOfStacks(input io.Reader, move CrateMoverFunc) string {
	scanner := bufio.NewScanner(input)
	stacks := parseStacks(scanner)

	for scanner.Scan() {
		quantity, from, to := parseNextRearrangement(scanner.Text())
		stacks[from], stacks[to] = move(quantity, stacks[from], stacks[to])
	}

	result := ""

	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}

	return result
}

func batchMove(numberOfMoves, batchSize int, from, to []rune) ([]rune, []rune) {
	for i := 0; i < numberOfMoves; i++ {
		// get the size of the origin stack
		srcStackLength := len(from)

		// pop the last (topmost) batch size number of crates from that stack
		crates := from[srcStackLength-batchSize:]
		from = from[:srcStackLength-batchSize]

		// and push them to the top of the target stack
		to = append(to, crates...)
	}

	return from, to
}

func parseNextRearrangement(line string) (quantity int, from int, to int) {
	// read arrangement instruction from line
	fmt.Sscanf(line, "move %d from %d to %d", &quantity, &from, &to)
	// adjust for zero based indexing
	from -= 1
	to -= 1
	return
}

func parseStacks(scanner *bufio.Scanner) [][]rune {
	scanner.Scan()
	line := scanner.Text()

	numberOfStacks := (len(line) + 1) / 4
	stacks := make([][]rune, numberOfStacks)

	// process each line until we find the "index line"
	for !strings.HasPrefix(line, " 1 ") {
		for stackIdx := 0; stackIdx < numberOfStacks; stackIdx++ {
			stackStartPos := stackIdx * 4
			crate := line[stackStartPos : stackStartPos+3]
			if crate[1] != ' ' {
				// insert the crate at the bottom of the stack
				stacks[stackIdx] = append(
					[]rune{rune(crate[1])},
					stacks[stackIdx]...,
				)
			}
		}
		scanner.Scan()
		line = scanner.Text()
	}

	// Forward past blank line
	scanner.Scan()

	return stacks
}
