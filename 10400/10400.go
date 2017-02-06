// UVa 10400 - Game Show Math

package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	p, target  int
	expression string
	nums       []int
	visited    []map[int]bool
)

func inRange(n int) bool { return n > -32000 && n < 32000 }

func dfs(level, result int, exp string) bool {
	if level == p {
		if result == target {
			expression = exp
			return true
		}
		return false
	}
	if visited[level][result] {
		return false
	}
	if num, str := nums[level], strconv.Itoa(nums[level]); inRange(result+num) && dfs(level+1, result+num, exp+"+"+str) ||
		inRange(result-num) && dfs(level+1, result-num, exp+"-"+str) ||
		inRange(result*num) && dfs(level+1, result*num, exp+"*"+str) ||
		result%num == 0 && inRange(result*num) && dfs(level+1, result/num, exp+"/"+str) {
		return true
	}
	visited[level][result] = true
	return false
}

func main() {
	in, _ := os.Open("10400.in")
	defer in.Close()
	out, _ := os.Create("10400.out")
	defer out.Close()

	var n int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &p)
		nums = make([]int, p)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		fmt.Fscanf(in, "%d", &target)
		visited = make([]map[int]bool, p)
		for i := range visited {
			visited[i] = make(map[int]bool)
		}
		if dfs(1, nums[0], strconv.Itoa(nums[0])) {
			fmt.Fprintf(out, "%s=%d\n", expression, target)
		} else {
			fmt.Fprintln(out, "NO EXPRESSION")
		}
	}
}
