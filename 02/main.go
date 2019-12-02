package main

import (
	"fmt"
)

var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 2, 9, 19, 23, 1, 9, 23, 27, 2, 27, 9, 31, 1, 31, 5, 35, 2, 35, 9, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 6, 51, 2, 51, 10, 55, 1, 9, 55, 59, 2, 6, 59, 63, 1, 63, 6, 67, 1, 67, 10, 71, 1, 71, 10, 75, 2, 9, 75, 79, 1, 5, 79, 83, 2, 9, 83, 87, 1, 87, 9, 91, 2, 91, 13, 95, 1, 95, 9, 99, 1, 99, 6, 103, 2, 103, 6, 107, 1, 107, 5, 111, 1, 13, 111, 115, 2, 115, 6, 119, 1, 119, 5, 123, 1, 2, 123, 127, 1, 6, 127, 0, 99, 2, 14, 0, 0}

func prt(p []int) {
	for i := 0; i < len(p); {
		opcode := p[i]
		switch opcode {
		case 1, 2:
			fmt.Printf("%d,%d,%d,%d\n", p[i], p[i+1], p[i+2], p[i+3])
			i += 4
		case 99:
			fmt.Println("99")
			i += 1
		default:
			for ; i < len(p); i++ {
				fmt.Printf("%d,", p[i])
			}
		}
	}
	fmt.Println()
}

func step(p []int, pc int) int {
	opcode := p[pc]
	switch opcode {
	case 1:
		s := p[p[pc+1]] + p[p[pc+2]]
		p[p[pc+3]] = s
		return pc + 4
	case 2:
		d := p[p[pc+1]] * p[p[pc+2]]
		p[p[pc+3]] = d
		return pc + 4
	case 99:
		return -1
	}
	panic(opcode)
}

func main() {
	prog := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	pc := 0

	startover := func() {
		prog = make([]int, len(input))
		copy(prog, input)
		prog[1] = 12
		prog[2] = 2
		pc = 0
	}

	startover()
	for pc >= 0 {
		pc = step(prog, pc)
	}
	fmt.Println("Day 01, Part 1: ", prog[0])

	answer := 0
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			startover()
			prog[1] = noun
			prog[2] = verb
			for pc >= 0 {
				pc = step(prog, pc)
			}
			if prog[0] == 19690720 {
				answer = 100*noun + verb
			}
		}
	}
	fmt.Println("Day 01, Part 2: ", answer)
}
