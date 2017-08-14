// UVa 10014 - Simple calculations

package main

import (
	"fmt"
	"os"
)

var (
	n      int
	af, al float64
)

func solve(c []float64) float64 {
	// a[1] = (n*a[0] + a[n+1] - 2*(n*c[1] + (n-1)*c[2] + ... + 1*c[n])) / (n + 1)
	var sum float64
	for i, cv := range c {
		sum += float64(n-i) * cv
	}
	return (float64(n)*af + al - 2*sum) / float64(n+1)
}

func main() {
	in, _ := os.Open("10014.in")
	defer in.Close()
	out, _ := os.Create("10014.out")
	defer out.Close()

	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		fmt.Fscanf(in, "%f", &af)
		fmt.Fscanf(in, "%f", &al)
		c := make([]float64, n)
		for i := range c {
			fmt.Fscanf(in, "%f", &c[i])
		}
		fmt.Fprintf(out, "%.2f\n", solve(c))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
