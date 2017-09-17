// UVa 10527 - Persistent Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

func solve(line string) string {
	if len(line) == 1 {
		return "1" + line
	}
	var n, tmp big.Int
	n.SetString(line, 10)
	pos := make([]int, 10)
	for i := 9; i > 1; i-- {
		for d := big.NewInt(int64(i)); ; pos[i]++ {
			tmp.Set(&n)
			if tmp.Mod(&tmp, d); tmp.Cmp(zero) != 0 {
				break
			}
			n.Div(&n, d)
		}
	}
	if n.Cmp(one) != 0 {
		return "There is no such number."
	}
	var ans string
	for i, d := range pos {
		if d != 0 {
			ans = ans + strings.Repeat(string(i+'0'), d)
		}
	}
	return ans
}

func main() {
	in, _ := os.Open("10527.in")
	defer in.Close()
	out, _ := os.Create("10527.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "-1" {
			break
		}
		fmt.Fprintln(out, solve(line))
	}
}
