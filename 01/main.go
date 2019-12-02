package main

import (
	"fmt"

	"github.com/vdobler/aoc19/aoc"
)

type Mass struct {
	M int
}

func fuel(m int) int {
	return m/3 - 2
}

func totalfuel(m int) int {
	t := 0
	f := fuel(m)
	for f > 0 {
		t += f
		f = fuel(f)
	}
	return t
}

func main() {
	mass := []Mass{
		Mass{12},
		Mass{14},
		Mass{1969},
		Mass{100756},
	}
	aoc.ReadInput(&mass, " ")

	sum := 0
	for i := range mass {
		f := fuel(mass[i].M)
		// fmt.Println(mass[i].M, " --> ", f)
		sum += f
	}
	fmt.Println("Day 01, Part 1: ", sum)

	total := 0
	for i := range mass {
		t := totalfuel(mass[i].M)
		// fmt.Println(mass[i].M, " --> ", t)
		total += t
	}
	fmt.Println("Day 01, Part 2: ", total)

}
