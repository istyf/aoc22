package crt

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartOne(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	accumulator := 0
	clock := 0
	register := 1

	accumulate := func(clock, accumulator, register int) int {
		if clock == 20 || (clock-20)%40 == 0 {
			return accumulator + (clock * register)
		}

		return accumulator
	}

	for scanner.Scan() {
		instruction := scanner.Text()
		clock++

		if instruction == "noop" {
			accumulator = accumulate(clock, accumulator, register)
		} else if strings.HasPrefix(instruction, "addx") {
			accumulator = accumulate(clock, accumulator, register)
			accumulator = accumulate(clock+1, accumulator, register)
			clock++

			amount, _ := strconv.ParseInt(strings.Split(instruction, " ")[1], 10, 64)
			register += int(amount)
		} else {
			panic("cpu halted; illegal instruction")
		}
	}

	return strconv.FormatInt(int64(accumulator), 10), nil
}

func PartTwo(input io.Reader) (string, error) {

	scanner := bufio.NewScanner(input)
	instructions := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}

		return scanner.Text(), true
	}

	cpu := newCPU(instructions)

	const BitmapHeight int = 6
	const BitmapWidth int = 40

	bitmap := newBitmap(BitmapWidth, BitmapHeight)

	for cycle := 0; cycle < len(bitmap); cycle++ {
		xpos := cycle % BitmapWidth
		if xpos >= cpu.register-1 && xpos <= cpu.register+1 {
			bitmap[cycle] = '#'
		}
		cpu.tick()
	}

	var lines []string
	for line := 0; line < BitmapHeight; line++ {
		lines = append(lines, string(bitmap[line*BitmapWidth:(line+1)*BitmapWidth]))
	}

	return "\n" + strings.Join(lines, "\n") + "\n", nil
}

func newBitmap(width, height int) []byte {
	bitmap := make([]byte, width*height)
	for b := 0; b < (width * height); b++ {
		bitmap[b] = '.'
	}
	return bitmap
}

func newCPU(instructions func() (string, bool)) *cpu {
	return &cpu{
		pop:      instructions,
		register: 1,
	}
}

type cpu struct {
	pop      func() (string, bool)
	register int
	adding   bool
	amount   int
}

func (c *cpu) tick() {
	if c.adding {
		c.register += c.amount
		c.adding = false
		return
	}

	instr, ok := c.pop()
	if !ok {
		return
	}

	if strings.HasPrefix(instr, "addx") {
		amt, _ := strconv.ParseInt(strings.Split(instr, " ")[1], 10, 64)
		c.amount = int(amt)
		c.adding = true
	}
}
