// UVa 151 - Power Crisis

package main

import (
	"fmt"
	"os"
)

func josephus(n, k int) int {
	if n == 1 {
		return 1
	}
	return (josephus(n-1, k)+k-1)%n + 1
}

func main() {
	in, _ := os.Open("151.in")
	defer in.Close()
	out, _ := os.Create("151.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		for k := 1; ; k++ {
			// region 1 is blacked out first, so count it out and consider n-1 regions with 12th is the last
			if josephus(n-1, k) == 12 {
				fmt.Fprintln(out, k)
				break
			}
		}
	}
}
