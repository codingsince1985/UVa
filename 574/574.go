// UVa 574 - Sum It Up

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	t       int
	nums    []int
	answers [][]string
)

func sameAsLast(selected []string) bool {
	if len(answers) == 0 {
		return false
	}
	for i := range selected {
		if selected[i] != answers[len(answers)-1][i] {
			return false
		}
	}
	return true
}

func backtracking(sum, idx int, selected []string) {
	if sum == t {
		if !sameAsLast(selected) {
			answers = append(answers, selected)
		}
		return
	}
	if idx == len(nums) || sum > t || len(selected) == len(nums) {
		return
	}
	backtracking(sum+nums[idx], idx+1, append(selected, strconv.Itoa(nums[idx])))
	backtracking(sum, idx+1, selected)
}

func main() {
	in, _ := os.Open("574.in")
	defer in.Close()
	out, _ := os.Create("574.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d%d", &t, &n); t == 0 && n == 0 {
			break
		}
		nums = make([]int, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		answers = nil
		backtracking(0, 0, nil)
		if fmt.Fprintf(out, "Sums of %d:\n", t); len(answers) == 0 {
			fmt.Fprintln(out, "NONE")
		} else {
			for _, answer := range answers {
				fmt.Fprintln(out, strings.Join(answer, "+"))
			}
		}
	}
}
