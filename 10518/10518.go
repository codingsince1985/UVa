// UVa 10518 - How Many Calls?

package main

import (
	"fmt"
	"os"
)

const max = 1000000

func solve(n, b int64) int64 {
	d := make([]int64, max)
	d[0], d[1] = 1, 1
	for m := 2; m < max; m++ {
		if d[m] = (d[m-1] + d[m-2] + 1) % b; d[m] == 1 && d[m-1] == 1 {
			return d[n%int64(m-1)]
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10518.in")
	defer in.Close()
	out, _ := os.Create("10518.out")
	defer out.Close()

	var n, b int64
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &b); n == 0 && b == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: %d %d %d\n", kase, n, b, solve(n, b))
	}
}
