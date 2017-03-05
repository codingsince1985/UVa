// UVa 10940 - Throwing cards away II

package main

import (
	"fmt"
	"os"
)

const max = 500000

var cache = func() []int {
	cache := make([]int, max+1)
	cache[0], cache[1] = 0, 1
	for n, i := 1, 2; i <= max; i++ {
		cache[i] = n * 2
		if n*2 == i {
			n = 1
		} else {
			n++
		}
	}
	return cache
}()

func main() {
	in, _ := os.Open("10940.in")
	defer in.Close()
	out, _ := os.Create("10940.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, cache[n])
	}
}
