// UVa 640 - Self Numbers

package main

import (
	"fmt"
	"os"
)

const max = 1000000

func digitadition(n int) int {
	total := n
	for n > 0 {
		total += n % 10
		n /= 10
	}
	return total
}

func main() {
	out, _ := os.Create("640.out")
	defer out.Close()

	notSelfNumbers := make(map[int]bool)

	for i := 1; i <= max; i++ {
		sum := digitadition(i)
		if sum <= max && !notSelfNumbers[sum] {
			notSelfNumbers[sum] = true
		}
		if !notSelfNumbers[i] {
			fmt.Fprintln(out, i)
		}
	}
}
