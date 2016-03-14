// UVa 10346 - Peter's Smokes

package main

import (
	"fmt"
	"os"
)

func solve(n, k int) int {
	cnt := n
	for n >= k {
		newCigar := n / k
		cnt += newCigar
		n = newCigar + n%k
	}
	return cnt
}

func main() {
	in, _ := os.Open("10346.in")
	defer in.Close()
	out, _ := os.Create("10346.out")
	defer out.Close()

	var n, k int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &k); err != nil {
			break
		}
		fmt.Fprintln(out, solve(n, k))
	}
}
