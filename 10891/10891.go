// UVa 10891 - Game of Sum

package main

import (
	"fmt"
	"os"
)

var (
	sum     []int
	dp      [][]int
	visited [][]bool
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find(l, r int) int {
	if !visited[l][r] {
		visited[l][r] = true
		var b int
		for k := l + 1; k <= r; k++ {
			b = min(b, find(k, r))
		}
		for k := r - 1; k >= l; k-- {
			b = min(b, find(l, k))
		}
		dp[l][r] = sum[r] - sum[l-1] - b
	}
	return dp[l][r]
}

func solve(n int, nums []int) int {
	sum = make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	visited = make([][]bool, n+1)
	dp = make([][]int, n+1)
	for i := range visited {
		visited[i] = make([]bool, n+1)
		dp[i] = make([]int, n+1)
	}
	return 2*find(1, n) - sum[n]
}

func main() {
	in, _ := os.Open("10891.in")
	defer in.Close()
	out, _ := os.Create("10891.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		fmt.Fprintln(out, solve(n, nums))
	}
}
