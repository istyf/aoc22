package tuningtrouble

import (
	"errors"
	"io"
	"strconv"
)

type MarkerFinderFunc func(r rune) bool

func PartOne(input io.Reader) (string, error) {
	const StartOfPacketMarkerSize int = 4
	return findMarker(input, newMarkerFinder(StartOfPacketMarkerSize))
}

func PartTwo(input io.Reader) (string, error) {
	const StartOfMessageMarkerSize int = 14
	return findMarker(input, newMarkerFinder(StartOfMessageMarkerSize))
}

func findMarker(input io.Reader, isEndOfMarker MarkerFinderFunc) (string, error) {
	b, err := io.ReadAll(input)
	if err != nil {
		return "", err
	}

	for idx, c := range string(b) {
		if isEndOfMarker(c) {
			return strconv.FormatInt(int64(idx+1), 10), nil
		}
	}

	return "", errors.New("no marker found")
}

func newMarkerFinder(size int) MarkerFinderFunc {

	lookupIndex := func(r rune) int { return int(r - 'A') }
	maxLookupIndex := lookupIndex('z')
	numberOfRunesWithinWindow := make([]int, maxLookupIndex+1)

	memory := make([]rune, size)
	numberOfRunesRead := 0

	return func(r rune) bool {
		insertPosition := numberOfRunesRead % size
		numberOfRunesRead++

		leastRecentlySeen := memory[insertPosition]
		memory[insertPosition] = r
		numberOfRunesWithinWindow[lookupIndex(r)]++

		if numberOfRunesRead < 4 {
			return false
		}

		if leastRecentlySeen != 0 {
			numberOfRunesWithinWindow[lookupIndex(leastRecentlySeen)]--
		}

		for x := 0; x < size-1; x++ {
			if numberOfRunesWithinWindow[lookupIndex(memory[x])] > 1 {
				return false
			}
		}

		return true
	}
}
