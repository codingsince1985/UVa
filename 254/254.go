// UVa 254 - Towers of Hanoi

package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	pegs [3]int
	bin  string
)

func solve(n, from, to int) {
	if n < 0 {
		return
	}
	tmp := 3 - from - to
	if len(bin) >= n+1 && bin[len(bin)-n-1] == '1' {
		pegs[from] -= (n + 1)
		pegs[tmp] += n
		pegs[to]++
		solve(n-1, tmp, to)
	} else {
		solve(n-1, from, tmp)
	}
}

func main() {
	in, _ := os.Open("254.in")
	defer in.Close()
	out, _ := os.Create("254.out")
	defer out.Close()

	var n int
	var m big.Int
	var s string
	for {
		if fmt.Fscanf(in, "%d %s", &n, &s); n == 0 && s == "0" {
			break
		}
		m.SetString(s, 10)
		bin = fmt.Sprintf("%b", &m)
		pegs = [3]int{n, 0, 0}
		if n%2 == 0 {
			solve(n-1, 0, 2)
		} else {
			solve(n-1, 0, 1)
		}
		for i, peg := range pegs {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, peg)
		}
		fmt.Fprintln(out)
	}
}
