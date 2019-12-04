package main

import "fmt"

func valid(n int) bool {
	d6, n := n%10, n/10
	d5, n := n%10, n/10
	if d6 < d5 {
		return false
	}
	d4, n := n%10, n/10
	if d5 < d4 {
		return false
	}
	d3, n := n%10, n/10
	if d4 < d3 {
		return false
	}
	d2, n := n%10, n/10
	if d3 < d2 {
		return false
	}
	d1, n := n%10, n/10
	if d2 < d1 {
		return false
	}

	return d1 == d2 || d2 == d3 || d3 == d4 || d4 == d5 || d5 == d6
}

func valid2(n int) bool {
	d6, n := n%10, n/10
	d5, n := n%10, n/10
	if d6 < d5 {
		return false
	}
	d4, n := n%10, n/10
	if d5 < d4 {
		return false
	}
	d3, n := n%10, n/10
	if d4 < d3 {
		return false
	}
	d2, n := n%10, n/10
	if d3 < d2 {
		return false
	}
	d1, n := n%10, n/10
	if d2 < d1 {
		return false
	}

	if d1 == d2 && d2 != d3 {
		return true
	}
	if d2 == d3 && d1 != d2 && d3 != d4 {
		return true
	}
	if d3 == d4 && d2 != d3 && d4 != d5 {
		return true
	}
	if d4 == d5 && d3 != d4 && d5 != d6 {
		return true
	}
	if d5 == d6 && d4 != d5 {
		return true
	}

	return false
}

func main() {
	count := 0
	for n := 124075; n <= 580769; n++ {
		if valid(n) {
			count++
		}
	}
	fmt.Println("Day 04, Part 1:", count)

	count2 := 0
	for n := 124075; n <= 580769; n++ {
		if valid2(n) {
			count2++
		}
	}
	fmt.Println("Day 04, Part 2:", count2)

}
