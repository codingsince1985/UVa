// UVa 264 - Count on Cantor

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("264.in")
	defer in.Close()
	out, _ := os.Create("264.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		i, cnt := 1, 0
		for cnt+i < n {
			cnt += i
			i++
		}
		num := n - cnt
		if i%2 != 0 {
			num = i + 1 - num
		}
		fmt.Fprintf(out, "TERM %d IS %d/%d\n", n, num, i+1-num)
	}
}
