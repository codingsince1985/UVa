// UVa 10017 - The Never Ending Towers of Hanoi

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	out     io.WriteCloser
	step, m int
	pegs    [][]int
)

func output() {
	for i, peg := range pegs {
		fmt.Fprintf(out, "%c=>", 'A'+i)
		if len(peg) > 0 {
			fmt.Fprint(out, "  ")
			for _, disk := range peg {
				fmt.Fprintf(out, " %d", disk)
			}
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out)
}

func hanoi(n, curr, via, target int) {
	if n > 1 {
		hanoi(n-1, curr, target, via)
	}
	pegs[target] = append(pegs[target], pegs[curr][len(pegs[curr])-1])
	pegs[curr] = pegs[curr][:len(pegs[curr])-1]
	if step++; step > m {
		return
	}
	output()
	if n > 1 {
		hanoi(n-1, via, curr, target)
	}
}

func solve(n int) {
	pegs = make([][]int, 3)
	for i := n; i > 0; i-- {
		pegs[0] = append(pegs[0], i)
	}
	step = 0
	output()
	hanoi(n, 0, 1, 2)
}

func main() {
	in, _ := os.Open("10017.in")
	defer in.Close()
	out, _ = os.Create("10017.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		fmt.Fprintf(out, "Problem #%d\n\n", kase)
		solve(n)
	}
}
