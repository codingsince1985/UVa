// UVa 10534 - Wavio Sequence

package main

import (
	"fmt"
	"math"
	"os"
)

var n int

func binarySearch(num int, stack []int) int {
	low, high := 0, len(stack)-1
	for low <= high {
		if mid := (low + high) / 2; num > stack[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func lis(num []int) []int {
	dp := make([]int, n)
	stack := []int{math.MinInt32}
	for i := range num {
		if num[i] > stack[len(stack)-1] {
			stack = append(stack, num[i])
		} else {
			stack[binarySearch(num[i], stack)] = num[i]
		}
		dp[i] = len(stack) - 1
	}
	return dp
}

func lds(num []int) []int {
	dp := make([]int, n)
	stack := []int{math.MinInt32}
	for i := n - 1; i >= 0; i-- {
		if num[i] > stack[len(stack)-1] {
			stack = append(stack, num[i])
		} else {
			stack[binarySearch(num[i], stack)] = num[i]
		}
		dp[i] = len(stack) - 1
	}
	return dp
}

func solve(num []int) int {
	dp1, dp2 := lis(num), lds(num)
	var length int
	for i := range num {
		if tmp := min(dp1[i], dp2[i]); length < tmp {
			length = tmp
		}
	}
	return 2*length - 1
}

func main() {
	in, _ := os.Open("10534.in")
	defer in.Close()
	out, _ := os.Create("10534.out")
	defer out.Close()

	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		fmt.Fprintln(out, solve(num))
	}
}
