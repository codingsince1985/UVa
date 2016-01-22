// UVa 371 - Ackermann Functions

package main

import (
	"fmt"
	"os"
)

var cache map[int]int

func calculate(i int) int {
	v := cache[i]
	if v != 0 {
		return v
	}

	if i % 2 == 0 {
		if i / 2 == 1 { // so that f(1)=3 as requested in problem
			return 1
		} else {
			v = 1 + calculate(i / 2)
		}
	} else {
		v = 1 + calculate(3 * i + 1)
	}
	cache[i] = v
	return v
}

func solve(l, h int) (int, int) {
	num, max := 0, 0
	for i := l; i <= h; i++ {
		n := calculate(i)
		if n > max {
			num, max = i, n
		}
	}
	return num, max
}

func main() {
	in, _ := os.Open("371.in")
	defer in.Close()
	out, _ := os.Create("371.out")
	defer out.Close()

	var l, h int
	cache = make(map[int]int)
	for {
		fmt.Fscanf(in, "%d%d", &l, &h)
		if l == 0 && h == 0 {
			break
		}
		num, max := solve(l, h)
		fmt.Fprintf(out, "Between %d and %d, %d generates the longest sequence of %d values.\n", l, h, num, max)
	}
}
