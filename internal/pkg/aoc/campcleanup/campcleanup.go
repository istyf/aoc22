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

	scanner := bufio.NewScanner(rd)
	overlappingSections := 0

	var s1Start, s1End int
	var s2Start, s2End int

	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &s1Start, &s1End, &s2Start, &s2End)
		if err != nil {
			return "", err
		}

		if sectionsAreOverlapping(s1Start, s1End, s2Start, s2End) {
			overlappingSections += 1
		}
	}

	return strconv.FormatInt(int64(overlappingSections), 10), nil
}

func oneSectionIsContainedByTheOther(fsStart, fsEnd, ssStart, ssEnd int) bool {
	return (ssStart >= fsStart && ssEnd <= fsEnd) || (fsStart >= ssStart && fsEnd <= ssEnd)
}

func sectionsAreOverlapping(fsStart, fsEnd, ssStart, ssEnd int) bool {
	// If we make sure that the second section does not start before the first ...
	if fsStart > ssStart {
		return sectionsAreOverlapping(ssStart, ssEnd, fsStart, fsEnd)
	}

	// ... we only need to check if it starts before the first one ends
	return ssStart <= fsEnd
}
