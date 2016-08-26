// UVa 10049 - Self-describing Sequence

package main

import (
	"fmt"
	"os"
)

const max = 2000000000

var dp []int

func f(n int) int {
	l, r := 0, len(dp)-1
	for l <= r {
		mid := (l + r) / 2
		if dp[mid] < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	in, _ := os.Open("10049.in")
	defer in.Close()
	out, _ := os.Create("10049.out")
	defer out.Close()

	dp = append(dp, []int{0, 1, 3}...)
	for i := 3; dp[i-1] <= max; i++ {
		dp = append(dp, dp[i-1]+f(i))
	}

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, f(n))
	}
}
