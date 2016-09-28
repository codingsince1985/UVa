// UVa 11059 - Maximum Product

package main

import (
	"fmt"
	"os"
)

func solve(nums []int64) int64 {
	var max int64
	for i := range nums {
		var product int64 = 1
		for j := i; j < len(nums); j++ {
			product *= nums[j]
			if product > max {
				max = product
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("11059.in")
	defer in.Close()
	out, _ := os.Create("11059.out")
	defer out.Close()

	var n, kase int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		nums := make([]int64, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		kase++
		max := solve(nums)
		if max < 0 {
			max = 0
		}
		fmt.Fprintf(out, "Case #%d: The maximum product is %d.\n\n", kase, max)
		fmt.Fscanln(in)
	}
}
