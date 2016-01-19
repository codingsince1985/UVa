// UVa 100 - The 3n + 1 problem

package main

import (
	"fmt"
	"os"
)

var cache map[int]int

func calculate(i int) int {
	v := cache[i]
	switch  {
	case v != 0:
		return v
	case i % 2 == 0:
		v = 1 + calculate(i / 2)
	default:
		v = 1 + calculate(i * 3 + 1)
	}
	cache[i] = v
	return v
}

func solve(i, j int) int {
	max := 0
	for k := i; k <= j; k++ {
		m := calculate(k)
		if m > max {
			max = m
		}
	}
	return max
}

func main() {
	in, _ := os.Open("100.in")
	defer in.Close()
	out, _ := os.Create("100.out")
	defer out.Close()

	var i, j int
	cache = make(map[int]int)
	cache[1] = 1
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &i, &j); err != nil {
			break
		}
		fmt.Fprintln(out, i, j, solve(i, j))
	}
}
