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
		missplacedItem := findCommonItem(first, second)
		sumOfPriorities += priority(missplacedItem)
	}

	return strconv.FormatInt(int64(sumOfPriorities), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {

	scanner := bufio.NewScanner(rd)
	sumOfPriorities := 0

	for scanner.Scan() {
		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()

		badge := findCommonItem(first, second, third)
		sumOfPriorities += priority(badge)
	}

	return strconv.FormatInt(int64(sumOfPriorities), 10), nil
}

func checkCompartments(contents string) (string, string) {
	numberOfItems := len(contents)
	return contents[0 : numberOfItems/2], contents[numberOfItems/2:]
}

func findCommonItem(first string, theRest ...string) rune {
	const LookupSize int = int('z'-'A') + 1
	numberOfCarriersOfItemType := make([]int, LookupSize)

	for _, item := range uniqueItems(first) {
		numberOfCarriersOfItemType[item-'A'] += 1
	}

	maxNumberOfCarriersOfAnItem := len(theRest)

	for _, another := range theRest {
		for _, item := range uniqueItems(another) {
			numberOfCarriersOfItemType[item-'A'] += 1
			if numberOfCarriersOfItemType[item-'A'] > maxNumberOfCarriersOfAnItem {
				return item
			}
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

func uniqueItems(contents string) string {
	hasSeen := make(map[rune]bool)
	items := make([]rune, 0, len(contents))

	for _, item := range contents {
		if _, ok := hasSeen[item]; !ok {
			hasSeen[item] = true
			items = append(items, item)
		}
	}

	return string(items)
}
