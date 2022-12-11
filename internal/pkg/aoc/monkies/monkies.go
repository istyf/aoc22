package monkies

import (
	"io"
	"strconv"
)

func PartOne(input io.Reader) (string, error) {
	monkies, err := parseMonkies(input)
	if err != nil {
		return "", err
	}

	for round := 0; round < 20; round++ {
		for idx := range monkies {
			monkies[idx].inspectAndThrow(func(i item, receiver int) {
				monkies[receiver].items = append(monkies[receiver].items, i)
			})
		}
	}

	var mostActiceThrows, runnerUpThrows int

	for idx := range monkies {
		if monkies[idx].numberOfThrows > runnerUpThrows {
			runnerUpThrows = monkies[idx].numberOfThrows
		}

		if runnerUpThrows > mostActiceThrows {
			tmp := mostActiceThrows
			mostActiceThrows = runnerUpThrows
			runnerUpThrows = tmp
		}
	}

	return strconv.FormatInt(int64(mostActiceThrows*runnerUpThrows), 10), nil
}

type item struct {
	worryLevel int
}

type OpFunc func(int) int
type TestFunc func(int) bool

type monkey struct {
	items          []item
	operation      OpFunc
	test           TestFunc
	receivers      map[bool]int
	numberOfThrows int
}

func (m *monkey) inspectAndThrow(throw func(item, int)) {
	for _, i := range m.items {
		i.worryLevel = (m.operation(i.worryLevel) / 3)
		testResult := m.test(i.worryLevel)
		throw(i, m.receivers[testResult])
		m.numberOfThrows++
	}

	m.items = []item{}
}

func parseMonkies(input io.Reader) ([]monkey, error) {
	monkies := []monkey{
		{
			items:     []item{{79}, {98}},
			operation: mul(19),
			test:      divisibleBy(23),
			receivers: map[bool]int{true: 2, false: 3},
		},
		{
			items:     []item{{54}, {65}, {75}, {74}},
			operation: add(6),
			test:      divisibleBy(19),
			receivers: map[bool]int{true: 2, false: 0},
		},
		{
			items:     []item{{79}, {60}, {97}},
			operation: square(),
			test:      divisibleBy(13),
			receivers: map[bool]int{true: 1, false: 3},
		},
		{
			items:     []item{{74}},
			operation: add(3),
			test:      divisibleBy(17),
			receivers: map[bool]int{true: 0, false: 1},
		},
	}
	return monkies, nil
}

func add(amount int) OpFunc {
	return func(old int) int {
		return old + amount
	}
}

func mul(factor int) OpFunc {
	return func(old int) int {
		return old * factor
	}
}

func square() OpFunc {
	return func(old int) int {
		return old * old
	}
}

func divisibleBy(divider int) func(int) bool {
	return func(value int) bool {
		return (value % divider) == 0
	}
}
