package monkies

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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

func newItem(worry int64) *item {
	return &item{worryLevel: int(worry)}
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

func newMonkey() *monkey {
	return &monkey{
		items:     []item{},
		receivers: map[bool]int{},
	}
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
	monkies := []monkey{}

	scanner := bufio.NewScanner(input)
	currentMonkey := newMonkey()

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			monkies = append(monkies, *currentMonkey)
			currentMonkey = newMonkey()
			continue
		}

		colonPos := strings.Index(line, ":")

		if strings.Contains(line, "Starting items:") {
			for _, item := range strings.Split(line[colonPos+2:], ", ") {
				worry, _ := strconv.ParseInt(item, 10, 64)
				currentMonkey.items = append(currentMonkey.items, *newItem(worry))
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			line = line[len("  Operation: new = "):]

			parts := strings.Split(line, " ")
			if parts[0] != "old" {
				panic("that was unexpected")
			}

			if parts[1] == "*" {
				if parts[2] == "old" {
					currentMonkey.operation = square()
				} else {
					factor, _ := strconv.ParseInt(parts[2], 10, 64)
					currentMonkey.operation = mul(int(factor))
				}
			} else if parts[1] == "+" {
				amount, _ := strconv.ParseInt(parts[2], 10, 64)
				currentMonkey.operation = add(int(amount))
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			line = line[len("  Test: divisible by "):]
			divisor, _ := strconv.ParseInt(line, 10, 64)
			currentMonkey.test = divisibleBy(int(divisor))
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			line = line[len("    If true: throw to monkey "):]
			receiver, _ := strconv.ParseInt(line, 10, 64)
			currentMonkey.receivers[true] = int(receiver)
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			line = line[len("    If false: throw to monkey "):]
			receiver, _ := strconv.ParseInt(line, 10, 64)
			currentMonkey.receivers[false] = int(receiver)
		}
	}

	monkies = append(monkies, *currentMonkey)

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
