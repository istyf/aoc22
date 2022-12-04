package campcleanup

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func PartOne(rd io.Reader) (string, error) {

	scanner := bufio.NewScanner(rd)
	fullyContainingAssignments := 0

	var s1Start, s1End int
	var s2Start, s2End int

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &s1Start, &s1End, &s2Start, &s2End)
		if err != nil {
			return "", err
		}

		if oneSectionIsContainedByTheOther(s1Start, s1End, s2Start, s2End) {
			fullyContainingAssignments += 1
		}
	}

	return strconv.FormatInt(int64(fullyContainingAssignments), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {
	return "", nil
}

func oneSectionIsContainedByTheOther(fsStart, fsEnd, ssStart, ssEnd int) bool {
	return (ssStart >= fsStart && ssEnd <= fsEnd) || (fsStart >= ssStart && fsEnd <= ssEnd)
}
