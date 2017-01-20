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

	var n, i, cnt, num int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		for i, cnt = 1, 0; cnt+i < n; i++ {
			cnt += i
		}
		if num = n - cnt; i%2 != 0 {
			num = i + 1 - num
		}
		fmt.Fprintf(out, "TERM %d IS %d/%d\n", n, num, i+1-num)
	}
}
