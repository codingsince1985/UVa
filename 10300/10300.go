// UVa 10300 - Ecological Premium

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10300.in")
	defer in.Close()
	out, _ := os.Create("10300.out")
	defer out.Close()

	var n, f int
	var s, a, e, t uint64
	fmt.Fscanf(in, "%d", &n)
	for i := 0; i < n; i++ {
		t = 0
		fmt.Fscanf(in, "%d", &f)
		for j := 0; j < f; j++ {
			fmt.Fscanf(in, "%d%d%d", &s, &a, &e)
			t += s * e
		}
		fmt.Fprintln(out, t)
	}
}
