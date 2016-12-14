// UVa 457 - Linear Cellular Automata

package main

import (
	"fmt"
	"os"
)

var (
	symbol = [4]byte{' ', '.', 'x', 'W'}
	out    *os.File
)

func output(dish []int) {
	for _, vi := range dish {
		fmt.Fprintf(out, "%c", symbol[vi])
	}
	fmt.Fprintln(out)
}

func solve(dna []int) {
	dish := make([]int, 42)
	dish[20] = 1
	for i := 0; i < 50; i++ {
		output(dish[1:41])
		next := make([]int, 42)
		for i := 1; i <= 40; i++ {
			next[i] = dna[dish[i-1]+dish[i]+dish[i+1]]
		}
		copy(dish, next)
	}
}

func main() {
	in, _ := os.Open("457.in")
	defer in.Close()
	out, _ = os.Create("457.out")
	defer out.Close()

	var kase int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanln(in)
		dna := make([]int, 10)
		for i := range dna {
			fmt.Fscanf(in, "%d", &dna[i])
		}
		solve(dna)
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
