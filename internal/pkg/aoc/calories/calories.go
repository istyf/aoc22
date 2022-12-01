package calories

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Solve(rd io.Reader) (string, error) {
	var maxCalories uint64 = 0
	var currentTotal uint64 = 0

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentTotal > maxCalories {
				maxCalories = currentTotal
			}
			currentTotal = 0
			continue
		}

		calories, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return "", fmt.Errorf("bad input data: %w", err)
		}

		currentTotal += calories
	}

	return strconv.FormatUint(maxCalories, 10), nil
}
