// UVa 10177 - (2/3/4)-D Sqr/Rects/Cubes/Boxes?

package main

import (
	"fmt"
	"io"
	"os"
)

func solve(out io.Writer, n int) {
	s2 := n * (n + 1) * (2*n + 1) / 6
	r2 := (n+1)*n/2*(n+1)*n/2 - s2
	s3 := n * (n + 1) / 2 * n * (n + 1) / 2
	r3 := (n+1)*n/2*(n+1)*n/2*(n+1)*n/2 - s3
	var s4 int
	for i := 1; i <= n; i++ {
		s4 += i * i * i * i
	}
	r4 := (n+1)*n/2*(n+1)*n/2*(n+1)*n/2*(n+1)*n/2 - s4
	fmt.Fprintln(out, s2, r2, s3, r3, s4, r4)
}

func main() {
	in, _ := os.Open("10177.in")
	defer in.Close()
	out, _ := os.Create("10177.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		solve(out, n)
	}
}
