// UVa 357 - Let Me Count The Ways

package main

import (
	"fmt"
	"os"
)

var cs = []int{1, 5, 10, 25, 50}

func main() {
	in, _ := os.Open("357.in")
	defer in.Close()
	out, _ := os.Create("357.out")
	defer out.Close()

	var v int
	for {
		if _, err := fmt.Fscanf(in, "%d", &v); err != nil {
			break
		}

		dp := make([]int64, v+1)
		dp[0] = 1
		for i := range cs {
			c := cs[i]
			for j := c; j <= v; j++ {
				dp[j] += dp[j-c]
			}
		}
		fmt.Fprintf(out, "There are %v ways to produce %d cents change.\n", dp[v], v)
	}
}
