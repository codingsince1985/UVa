// UVa 130 - Roman Roulette

package main

import (
	"fmt"
	"os"
)

func findNext(circle []int, next, k int) int {
	for k > 0 {
		next = (next + 1) % len(circle)
		if circle[next] != 0 {
			k--
		}
	}
	return next
}

func main() {
	in, _ := os.Open("130.in")
	defer in.Close()
	out, _ := os.Create("130.out")
	defer out.Close()

	var n, k int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &k); n == 0 && k == 0 {
			break
		}
		circle := make([]int, n)
		for i := range circle {
			circle[i] = i + 1
		}
		this := (k - 1) % n
		for i := 0; i < n-1; i++ {
			circle[this] = 0
			next := findNext(circle, this, k)
			circle[this], circle[next] = circle[next], circle[this]
			this = findNext(circle, this, k)
		}
		fmt.Fprintln(out, circle[this])
	}
}
