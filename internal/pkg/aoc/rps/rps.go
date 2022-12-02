package rps

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func PartOne(rd io.Reader) (string, error) {

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

	score := map[string]map[string]int{
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

	parseTurn := func(turn string) (string, string, error) {
		hands := strings.Split(turn, " ")
		if len(hands) != 2 {
			return "", "", errors.New("erroneous input")
		}

		return hands[0], hands[1], nil
	}

	scanner := bufio.NewScanner(rd)
	totalScore := 0

	for scanner.Scan() {
		elfHand, myHand, err := parseTurn(scanner.Text())

		if err != nil {
			return "", err
		}

		totalScore += score[elfHand][myHand]
	}

	return strconv.FormatInt(int64(totalScore), 10), nil
}
