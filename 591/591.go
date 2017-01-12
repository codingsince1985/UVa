// UVa 591 - Box of Bricks

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("591.in")
	defer in.Close()
	out, _ := os.Create("591.out")
	defer out.Close()

	var n, t, avg, cnt int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		stacks := make([]int, n)
		t = 0
		for i := range stacks {
			fmt.Fscanf(in, "%d", &stacks[i])
			t += stacks[i]
		}
		cnt, avg = 0, t/n
		for _, v := range stacks {
			if v > avg {
				cnt += v - avg
			}
		}
		fmt.Fprintf(out, "Set #%d\n", kase)
		fmt.Fprintf(out, "The minimum number of moves is %d.\n", cnt)
	}
}
