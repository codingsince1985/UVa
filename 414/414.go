// UVa 414 - Machined Surfaces

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("414.in")
	defer in.Close()
	out, _ := os.Create("414.out")
	defer out.Close()

	var n int
	var ch byte
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		nums := make([]int, n)
		max := 0
		for ; n > 0; n-- {
			num := 0
			for i := 1; i <= 25; i++ {
				if fmt.Fscanf(in, "%c", &ch); ch == 'X' {
					num++
				}
			}
			fmt.Fscanln(in)
			if num > max {
				max = num
			}
			nums[n-1] = num
		}
		total := 0
		for _, v := range nums {
			total += max - v
		}
		fmt.Fprintln(out, total)
	}
}
