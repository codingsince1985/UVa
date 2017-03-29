// UVa 10057 - A mid-summer night's dream.

package main

import (
	"fmt"
	"os"
	"sort"
)

func count(nums []int, f int) int {
	var cnt int
	idx := sort.SearchInts(nums, f)
	for ; idx < len(nums) && nums[idx] == f; idx++ {
		cnt++
	}
	return cnt
}

func solve(nums []int) (int, int, int) {
	sort.Ints(nums)
	var median, n1, n2 int
	if size := len(nums); size%2 == 0 {
		median = nums[size/2-1]
		n1 = count(nums, median)
		if median != nums[size/2] {
			n1 += count(nums, nums[size/2])
		}
		n2 = nums[size/2] - median + 1
	} else {
		median = nums[size/2]
		n1 = count(nums, median)
		n2 = 1
	}
	return median, n1, n2
}

func main() {
	in, _ := os.Open("10057.in")
	defer in.Close()
	out, _ := os.Create("10057.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		median, n1, n2 := solve(nums)
		fmt.Fprintf(out, "%d %d %d\n", median, n1, n2)
	}
}
