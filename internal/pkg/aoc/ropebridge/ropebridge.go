package ropebridge

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

func PartOne(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	head := position{}
	tail := position{}

	visitedPositions := map[string]bool{tail.pos(): true}

	for scanner.Scan() {
		var direction string
		var distance int

		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &distance)

		for move := 0; move < distance; move++ {
			head = head.move(direction)
			tail = tail.chase(head)
			visitedPositions[tail.pos()] = true
		}
	}

	numVisitedPositions := len(visitedPositions)
	return strconv.FormatInt(int64(numVisitedPositions), 10), nil
}

func PartTwo(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	head := position{}
	knots := make([]position, 9)
	tailpos := func() string { return knots[len(knots)-1].pos() }

	visitedPositions := map[string]bool{tailpos(): true}

	for scanner.Scan() {
		var direction string
		var distance int

		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &distance)

		for move := 0; move < distance; move++ {
			head = head.move(direction)
			knots[0] = knots[0].chase(head)

			for k := 1; k < len(knots); k++ {
				knots[k] = knots[k].chase(knots[k-1])
			}

			visitedPositions[tailpos()] = true
		}
	}

	numVisitedPositions := len(visitedPositions)
	return strconv.FormatInt(int64(numVisitedPositions), 10), nil
}

type position struct {
	x int
	y int
}

func (p position) chase(other position) position {
	if isTouching(p, other) {
		return p
	}

	if p.x == other.x {
		if other.y > p.y {
			return position{p.x, p.y + 1}
		}

		return position{p.x, p.y - 1}
	} else if p.y == other.y {
		if other.x > p.x {
			return position{p.x + 1, p.y}
		}

		return position{p.x - 1, p.y}
	} else if other.x > p.x {
		if other.y > p.y {
			return position{p.x + 1, p.y + 1}
		}

		return position{p.x + 1, p.y - 1}
	} else if other.x < p.x {
		if other.y > p.y {
			return position{p.x - 1, p.y + 1}
		}

		return position{p.x - 1, p.y - 1}
	}

	panic("target lost captain!")
}

func (p position) move(direction string) position {
	switch direction {
	case "U":
		return position{p.x, p.y - 1}
	case "D":
		return position{p.x, p.y + 1}
	case "R":
		return position{p.x + 1, p.y}
	case "L":
		return position{p.x - 1, p.y}
	}

	panic("unknown direction: " + direction)
}

func (p position) pos() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func isTouching(lhs, rhs position) bool {
	xdiff := int(math.Abs(float64(lhs.x - rhs.x)))
	ydiff := int(math.Abs(float64(lhs.y - rhs.y)))

	return xdiff <= 1 && ydiff <= 1
}
