// UVa 11805 - Bafana Bafana

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11805.in")
	defer in.Close()
	out, _ := os.Create("11805.out")
	defer out.Close()

	var kase, n, k, p int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d%d%d", &n, &k, &p)
		end := (k + p) % n
		if end == 0 {
			end = n
		}
		fmt.Fprintf(out, "Case %d: %d\n", i, end)
	}
}
