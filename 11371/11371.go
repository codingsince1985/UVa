// UVa 11371 - Number Theory for Newbies

package main

import (
	"fmt"
	"os"
	"sort"
)

func toNumber(digits []int) int {
	var n int
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}

func a(digits []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(digits)))
	return toNumber(digits)
}

func b(digits []int) int {
	sort.Ints(digits)
	if digits[0] == 0 {
		for i, d := range digits {
			if d != 0 {
				digits[0], digits[i] = digits[i], 0
				break
			}
		}
	}
	return toNumber(digits)
}

func solve(n int) (int, int) {
	var digits []int
	for ; n > 0; n /= 10 {
		digits = append(digits, n%10)
	}
	return a(digits), b(digits)
}

func main() {
	in, _ := os.Open("11371.in")
	defer in.Close()
	out, _ := os.Create("11371.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		a, b := solve(n)
		fmt.Fprintf(out, "%d - %d = %d = 9 * %d\n", a, b, a-b, (a-b)/9)
	}
}
