// UVa 10905 - Children's Game

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func solve(nums []string) string {
	var oldNum, newNum big.Int
here:
	for found := true; found; {
		found = false
		for i := 0; i < len(nums)-1; i++ {
			for j := 1; j < len(nums); j++ {
				oldNum.SetString(strings.Join(nums, ""), 10)
				nums[i], nums[j] = nums[j], nums[i]
				newNum.SetString(strings.Join(nums, ""), 10)
				if newNum.Cmp(&oldNum) > 0 {
					found = true
					continue here
				}
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return strings.Join(nums, "")
}

func main() {
	in, _ := os.Open("10905.in")
	defer in.Close()
	out, _ := os.Create("10905.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		nums := make([]string, n)
		for i := range nums {
			fmt.Fscanf(in, "%s", &nums[i])
		}
		fmt.Fprintln(out, solve(nums))
	}
}
