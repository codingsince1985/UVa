// UVa 100 - The 3n + 1 problem

package main

import (
	"fmt"
	"os"
)

var cache = map[int]int{1: 1}

func calculate(i int) int {
	if _, ok := cache[i]; !ok {
		if i%2 == 0 {
			cache[i] = 1 + calculate(i/2)
		} else {
			cache[i] = 1 + calculate(i*3+1)
		}
	}
	return cache[i]
}

func solve(i, j int) int {
	max := 0
	for k := i; k <= j; k++ {
		if m := calculate(k); m > max {
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
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &i, &j); err != nil {
			break
		}
		fmt.Fprintln(out, i, j, solve(i, j))
	}
}
