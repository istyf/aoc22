package treetoptreehouse

import (
	"bufio"
	"io"
	"strconv"
)

func PartOne(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	row := scanner.Text()
	columnCount := len(row)

	rows := append(make([][]tree, 0, columnCount), toarr(row))
	rowIdx := 0

	for idx := range rows[rowIdx] {
		rows[rowIdx][idx].Visible = true
	}

	for scanner.Scan() {
		rowIdx++
		rows = append(rows, toarr(scanner.Text()))
		rows[rowIdx][0].Visible = true
		rows[rowIdx][columnCount-1].Visible = true
	}

	for idx := range rows[rowIdx] {
		rows[rowIdx][idx].Visible = true
	}

	for _, trees := range rows {
		highest := trees[0]
		for i := 1; i < columnCount; i++ {
			if !highest.obscures(trees[i]) {
				trees[i].Visible = true
				highest = trees[i]
			}
		}

		highest = trees[columnCount-1]
		for i := columnCount - 2; i >= 0; i-- {
			if !highest.obscures(trees[i]) {
				trees[i].Visible = true
				highest = trees[i]
			}
		}
	}

	for x := columnCount - 2; x > 0; x-- {
		highest := rows[0][x]
		for y := 1; y < len(rows)-1; y++ {
			if !highest.obscures(rows[y][x]) {
				rows[y][x].Visible = true
				highest = rows[y][x]
			}
		}

		highest = rows[len(rows)-1][x]
		for y := len(rows) - 2; y > 0; y-- {
			if !highest.obscures(rows[y][x]) {
				rows[y][x].Visible = true
				highest = rows[y][x]
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

func toarr(row string) []tree {
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

func (t tree) obscures(other tree) bool {
	return t.Height >= other.Height
}
