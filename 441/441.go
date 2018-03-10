// UVa 441 - Lotto

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.WriteCloser

func output(nums []int, selected []bool) {
	var tokens []string
	for i, v := range selected {
		if v {
			tokens = append(tokens, strconv.Itoa(nums[i]))
		}
	}
	fmt.Fprintln(out, strings.Join(tokens, " "))
}

func backtracking(nums []int, selected []bool) {
	cnt, largest := 0, -1
	for i, v := range selected {
		if v {
			cnt++
			largest = i
		}
	}
	if cnt == 6 {
		output(nums, selected)
		return
	}
	for i := range nums {
		if !selected[i] && i > largest {
			selected[i] = true
			backtracking(nums, selected)
			selected[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("441.in")
	defer in.Close()
	out, _ = os.Create("441.out")
	defer out.Close()

	var k int
	for first := true; ; {
		if fmt.Fscanf(in, "%d", &k); k == 0 {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		nums := make([]int, k)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		backtracking(nums, make([]bool, k))
	}
}
