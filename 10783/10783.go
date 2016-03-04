// UVa 10783 - Odd Sum

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10783.in")
	defer in.Close()
	out, _ := os.Create("10783.out")
	defer out.Close()

	var kase, n, l, h int
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		fmt.Fscanf(in, "%d\n%d", &l, &h)
		total := 0
		for i := l; i <= h; i++ {
			if i%2 != 0 {
				total += i
			}
		}
		kase++
		fmt.Fprintf(out, "Case %d: %d\n", kase, total)
		n--
	}
}
