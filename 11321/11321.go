// UVa 11321 - Sort! Sort!! and Sort!!!

package main

import (
	"fmt"
	"os"
	"sort"
)

func isOdd(n int) bool { return n%2 == 1 }

func cmp(n1, n2, m int) bool {
	if m1, m2 := n1%m, n2%m; m1 != m2 {
		return m1 < m2
	}
	odd1, odd2 := isOdd(n1), isOdd(n2)
	switch {
	case odd1 && odd2:
		return n1 > n2
	case !odd1 && !odd2:
		return n1 < n2
	default:
		return odd1
	}
}

func main() {
	in, _ := os.Open("11321.in")
	defer in.Close()
	out, _ := os.Create("11321.out")
	defer out.Close()

	var n, m int
	for {
		fmt.Fscanf(in, "%d%d", &n, &m)
		fmt.Fprintln(out, n, m)
		if n == 0 && m == 0 {
			break
		}
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		sort.Slice(nums, func(i, j int) bool { return cmp(nums[i], nums[j], m) })
		for _, num := range nums {
			fmt.Fprintln(out, num)
		}
	}
}
