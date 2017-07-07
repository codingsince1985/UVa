// UVa 11005 - Cheapest Base

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func convert(n, base int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%base)
		n /= base
	}
	return digits
}

func solve(costs []int, n int) []string {
	lowest := math.MaxInt32
	var bases []string
here:
	for base := 2; base <= 36; base++ {
		var total int
		for _, digit := range convert(n, base) {
			if total += costs[digit]; total > lowest {
				continue here
			}
		}
		if total <= lowest {
			if total < lowest {
				lowest = total
				bases = nil
			}
			bases = append(bases, strconv.Itoa(base))
		}
	}
	return bases
}

func main() {
	in, _ := os.Open("11005.in")
	defer in.Close()
	out, _ := os.Create("11005.out")
	defer out.Close()

	var k, q, n int
	fmt.Fscanf(in, "%d", &k)
	for kase := 1; kase <= k; kase++ {
		costs := make([]int, 36)
		for i := range costs {
			fmt.Fscanf(in, "%d", &costs[i])
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case %d:\n", kase)
		for fmt.Fscanf(in, "%d", &q); q > 0; q-- {
			fmt.Fscanf(in, "%d", &n)
			fmt.Fprintf(out, "Cheapest base(s) for number %d: %s\n", n, strings.Join(solve(costs, n), " "))
		}
	}
}
