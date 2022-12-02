package rps

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func PartOne(rd io.Reader) (string, error) {

	scanner := bufio.NewScanner(rd)
	totalScore := 0

	for scanner.Scan() {
		elfHand, myHand, err := parseTurn(scanner.Text())

		if err != nil {
			return "", err
		}

		totalScore += score(elfHand, myHand)
	}

	return strconv.FormatInt(int64(totalScore), 10), nil
}

func PartTwo(rd io.Reader) (string, error) {

	scanner := bufio.NewScanner(rd)
	totalScore := 0

	for scanner.Scan() {
		elfHand, instruction, err := parseTurn(scanner.Text())

		if err != nil {
			return "", err
		}

		myHand := followInstruction(elfHand, instruction)
		totalScore += score(elfHand, myHand)
	}

	return strconv.FormatInt(int64(totalScore), 10), nil
}

func parseTurn(line string) (string, string, error) {
	turn := strings.Split(line, " ")
	if len(turn) != 2 {
		return "", "", errors.New("erroneous input")
	}

	return turn[0], turn[1], nil
}

func followInstruction(elfHand, instruction string) string {
	const (
		VersusRock     string = "A"
		VersusPaper    string = "B"
		VersusScissors string = "C"

		ToWin   string = "Z"
		ToDraw  string = "Y"
		ToLoose string = "X"

		PlayRock     string = "X"
		PlayPaper    string = "Y"
		PlayScissors string = "Z"
	)

	requiredHand := map[string]map[string]string{
		VersusRock: {
			ToWin:   PlayPaper,
			ToDraw:  PlayRock,
			ToLoose: PlayScissors,
		},
		VersusPaper: {
			ToWin:   PlayScissors,
			ToDraw:  PlayPaper,
			ToLoose: PlayRock,
		},
		VersusScissors: {
			ToWin:   PlayRock,
			ToDraw:  PlayScissors,
			ToLoose: PlayPaper,
		},
	}

	return requiredHand[elfHand][instruction]
}

func score(elfHand, myHand string) int {
	const (
		VersusRock     string = "A"
		VersusPaper    string = "B"
		VersusScissors string = "C"

		IfIPlayRock     string = "X"
		IfIPlayPaper    string = "Y"
		IfIPlayScissors string = "Z"

		ChoiceOfRock     int = 1
		ChoiceOfPaper    int = 2
		ChoiceOfScissors int = 3

		Win  int = 6
		Draw int = 3
		Loss int = 0
	)

	compare := map[string]map[string]int{
		VersusRock: {
			IfIPlayRock:     (Draw + ChoiceOfRock),
			IfIPlayPaper:    (Win + ChoiceOfPaper),
			IfIPlayScissors: (Loss + ChoiceOfScissors),
		},
		VersusPaper: {
			IfIPlayRock:     (Loss + ChoiceOfRock),
			IfIPlayPaper:    (Draw + ChoiceOfPaper),
			IfIPlayScissors: (Win + ChoiceOfScissors),
		},
		VersusScissors: {
			IfIPlayRock:     (Win + ChoiceOfRock),
			IfIPlayPaper:    (Loss + ChoiceOfPaper),
			IfIPlayScissors: (Draw + ChoiceOfScissors),
		},
	}

	return compare[elfHand][myHand]
}
