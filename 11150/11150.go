// UVa 11150 - Cola

package main

import (
	"fmt"
	"os"
)

func solve(n int) (int, int) {
	cnt, empty := n, n
	for empty >= 3 {
		cnt += empty / 3
		empty = empty/3 + empty%3
	}
	return cnt, empty
}

func main() {
	in, _ := os.Open("11150.in")
	defer in.Close()
	out, _ := os.Create("11150.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		cnt, empty := solve(n)
		if empty == 2 {
			cnt++
		}
		fmt.Fprintln(out, cnt)
	}
}
