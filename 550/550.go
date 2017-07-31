// UVa 550 - Multiplying by Rotation

package main

import (
	"fmt"
	"os"
)

func solve(base, first, second int) int {
	length := 1
	for curr, carry := first, 0; ; length++ {
		tmp := curr*second + carry
		if carry, curr = tmp/base, tmp%base; carry == 0 && curr == first {
			break
		}
	}
	return length
}

func main() {
	in, _ := os.Open("550.in")
	defer in.Close()
	out, _ := os.Create("550.out")
	defer out.Close()

	var base, first, second int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &base, &first, &second); err != nil {
			break
		}
		fmt.Fprintln(out, solve(base, first, second))
	}
}
