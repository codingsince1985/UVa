// UVa 10190 - Divide, But Not Quite Conquer!

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("10190.in")
	defer in.Close()
	out, _ := os.Create("10190.out")
	defer out.Close()

	var n, m int
here:
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		if m <= 1 {
			fmt.Fprintln(out, "Boring!")
			continue
		}
		var nums []string
		for n >= 1 {
			nums = append(nums, strconv.Itoa(n))
			if n == 1 {
				break
			}
			if n%m == 0 {
				n /= m
			} else {
				fmt.Fprintln(out, "Boring!")
				continue here
			}
		}
		fmt.Fprintln(out, strings.Join(nums, " "))
	}
}
