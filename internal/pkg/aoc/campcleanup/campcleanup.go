package campcleanup

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func PartOne(rd io.Reader) (string, error) {
	answer, err := countMatchingInputLines(rd, oneSectionIsContainedByTheOther)

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(answer), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {
	answer, err := countMatchingInputLines(rd, overlappingSections)

	if err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(answer), 10), nil
}

func countMatchingInputLines(rd io.Reader, shouldCountSection func(int, int, int, int) bool) (int, error) {

	scanner := bufio.NewScanner(rd)
	count := 0

	var s1Start, s1End int
	var s2Start, s2End int

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &s1Start, &s1End, &s2Start, &s2End)
		if err != nil {
			return 0, err
		}

		if shouldCountSection(s1Start, s1End, s2Start, s2End) {
			count += 1
		}
	}

	return count, nil
}

func oneSectionIsContainedByTheOther(fsStart, fsEnd, ssStart, ssEnd int) bool {
	return (ssStart >= fsStart && ssEnd <= fsEnd) || (fsStart >= ssStart && fsEnd <= ssEnd)
}

func overlappingSections(fsStart, fsEnd, ssStart, ssEnd int) bool {
	// If we make sure that the second section does not start before the first ...
	if fsStart > ssStart {
		return overlappingSections(ssStart, ssEnd, fsStart, fsEnd)
	}

	// ... we only need to check if it starts before the first one ends
	return ssStart <= fsEnd
}
