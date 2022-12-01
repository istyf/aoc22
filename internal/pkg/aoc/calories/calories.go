package calories

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func PartOne(rd io.Reader) (string, error) {
	calorieTotals, err := parseElfInventory(rd)
	if err != nil {
		return "", err
	}

	nrofElves := len(calorieTotals)

	return strconv.FormatInt(int64(calorieTotals[nrofElves-1]), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {
	calorieTotals, err := parseElfInventory(rd)
	if err != nil {
		return "", err
	}

	nrofElves := len(calorieTotals)
	minimumPartySize := 3
	if nrofElves < minimumPartySize {
		return "", fmt.Errorf("there must be at least %d elves", minimumPartySize)
	}

	sumOfCalories := 0

	for idx := nrofElves - minimumPartySize; idx < nrofElves; idx++ {
		sumOfCalories += calorieTotals[idx]
	}

	return strconv.FormatInt(int64(sumOfCalories), 10), nil
}

func parseElfInventory(rd io.Reader) ([]int, error) {
	var inventories []int
	var currentTotal int

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			inventories = append(inventories, currentTotal)
			currentTotal = 0
			continue
		}

		calories, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("bad input data: %w", err)
		}

		currentTotal += int(calories)
	}

	if currentTotal > 0 {
		inventories = append(inventories, currentTotal)
	}

	if len(inventories) == 0 {
		return nil, fmt.Errorf("no input data")
	}

	sort.Ints(inventories)

	return inventories, nil
}
