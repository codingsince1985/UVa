// UVa 10705 - The Fun Number System

package main

import (
	"fmt"
	"os"
)

func solve(n int64, np string) string {
	var funk string
	for i := len(np) - 1; i >= 0; i-- {
		lastBit := n & 1
		funk = fmt.Sprint(lastBit) + funk
		n /= 2
		switch {
		case lastBit == 1 && np[i] == 'n' && n > 0:
			n++
		case lastBit == 1 && np[i] == 'p' && n < 0:
			n--
		}
	}
	if n != 0 {
		return "Impossible"
	}
	return funk
}

func main() {
	in, _ := os.Open("10705.in")
	defer in.Close()
	out, _ := os.Create("10705.out")
	defer out.Close()

	var kase, k, n int64
	var np string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d\n%s\n%d", &k, &np, &n)
		fmt.Fprintln(out, solve(n, np))
	}
}
