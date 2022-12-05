package supplystacks

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func PartOne(rd io.Reader) (string, error) {
	scanner := bufio.NewScanner(rd)
	stacks := parseStacks(scanner)

	const BatchSize int = 1

	for scanner.Scan() {
		quantity, from, to := parseNextRearrangement(scanner.Text())
		stacks[from], stacks[to] = move(quantity, BatchSize, stacks[from], stacks[to])
	}

	result := ""

	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}

	return result, nil
}

func PartTwo(rd io.Reader) (string, error) {
	scanner := bufio.NewScanner(rd)
	stacks := parseStacks(scanner)

	const SingleMove int = 1

	for scanner.Scan() {
		batchSize, from, to := parseNextRearrangement(scanner.Text())
		stacks[from], stacks[to] = move(SingleMove, batchSize, stacks[from], stacks[to])
	}

	result := ""

	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}

	return result, nil
}

func move(amount, batch int, from, to []rune) ([]rune, []rune) {
	for i := 0; i < amount; i++ {
		// get the size of the origin stack
		srcStackLength := len(from)
		// pop the last (topmost) batch size number of crates from that stack
		crates := from[srcStackLength-batch:]

		from = from[:srcStackLength-batch]
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
