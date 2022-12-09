package treetoptreehouse

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

func PartOne(input io.Reader) (string, error) {

	rows := loadForest(input)
	columnCount := len(rows[0])

	for y := len(rows) - 1; y >= 0; y-- {
		highest := -1
		for x := 0; x < columnCount-1; x++ {
			if highest < rows[y][x].Height {
				rows[y][x].Visible = true
				highest = rows[y][x].Height
			}
		}

		highest = -1
		for x := columnCount - 1; x > 0; x-- {
			if highest < rows[y][x].Height {
				rows[y][x].Visible = true
				highest = rows[y][x].Height
			}
		}
	}

	for x := columnCount - 2; x >= 0; x-- {
		highest := -1
		for y := 0; y < len(rows)-1; y++ {
			if highest < rows[y][x].Height {
				rows[y][x].Visible = true
				highest = rows[y][x].Height
			}
		}

		highest = -1
		for y := len(rows) - 1; y > 0; y-- {
			if highest < rows[y][x].Height {
				rows[y][x].Visible = true
				highest = rows[y][x].Height
			}
		}
	}

	numVisible := 0

	for x := 0; x < columnCount; x++ {
		for y := 0; y < len(rows); y++ {
			if rows[x][y].Visible {
				numVisible++
			}
		}
	}

	return strconv.FormatInt(int64(numVisible), 10), nil
}

func PartTwo(input io.Reader) (string, error) {

	rows := loadForest(input)
	columnCount := len(rows[0])

	boundsCheck := func(x, y int) (int, int, error) {
		if x < 0 || y < 0 || x >= columnCount || y >= len(rows) {
			return 0, 0, errors.New("out of bounds")
		}

		return x, y, nil
	}

	viewingDistanceFrom := func(x, y int, move func(int, int) (int, int)) int {
		candidateTree := rows[y][x]
		var err error

		if x, y, err = boundsCheck(move(x, y)); err != nil {
			return 0
		}

		distance := 0

		for err == nil {
			distance++

			nextTree := rows[y][x]
			if nextTree.Height >= candidateTree.Height {
				return distance
			}

			x, y, err = boundsCheck(move(x, y))
		}

		return distance
	}

	calculateScenicScore := func(x, y int) int {
		a := viewingDistanceFrom(x, y, func(x_, y_ int) (int, int) { return x_ - 1, y_ })
		b := viewingDistanceFrom(x, y, func(x_, y_ int) (int, int) { return x_ + 1, y_ })
		c := viewingDistanceFrom(x, y, func(x_, y_ int) (int, int) { return x_, y_ - 1 })
		d := viewingDistanceFrom(x, y, func(x_, y_ int) (int, int) { return x_, y_ + 1 })

		return a * b * c * d
	}

	maxScenicScore := 0

	for y := 1; y < len(rows)-1; y++ {
		for x := 1; x < columnCount-1; x++ {
			score := calculateScenicScore(x, y)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	return strconv.FormatInt(int64(maxScenicScore), 10), nil
}

func loadForest(input io.Reader) [][]tree {
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	row := scanner.Text()
	columnCount := len(row)

	rows := append(make([][]tree, 0, columnCount), toTreeLine(row))
	rowIdx := 0

	for scanner.Scan() {
		rowIdx++
		rows = append(rows, toTreeLine(scanner.Text()))
	}

	return rows
}

func toTreeLine(row string) []tree {
	arr := make([]tree, len(row))
	for idx, tree := range row {
		arr[idx] = *newTree(int(tree - '0'))
	}
	return arr
}

type tree struct {
	Height  int
	Visible bool
}

func newTree(height int) *tree {
	return &tree{
		Height: height,
	}
}
