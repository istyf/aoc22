package rucksack

import (
	"bufio"
	"io"
	"strconv"
)

func PartOne(rd io.Reader) (string, error) {

	scanner := bufio.NewScanner(rd)
	sumOfPriorities := 0

	for scanner.Scan() {
		contents := scanner.Text()
		first, second := checkCompartments(contents)
		missplacedItem := findError(first, second)
		sumOfPriorities += priority(missplacedItem)
	}

	return strconv.FormatInt(int64(sumOfPriorities), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {
	return "not implemented", nil
}

func checkCompartments(contents string) (string, string) {
	numberOfItems := len(contents)
	return contents[0 : numberOfItems/2], contents[numberOfItems/2:]
}

func findError(first, second string) rune {
	const LookupSize int = int('z'-'A') + 1
	isItemInFirstCompartment := make([]bool, LookupSize)

	for _, item := range first {
		isItemInFirstCompartment[item-'A'] = true
	}

	for _, item := range second {
		if isItemInFirstCompartment[item-'A'] {
			return item
		}
	}

	panic("this rucksack contains a bomb!")
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}

	return int(item-'A') + 27
}
