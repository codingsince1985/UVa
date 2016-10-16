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
	a1 := (float64(n)*af + al - 2*sum) / float64(n+1)
	return a1
}

func main() {
	in, _ := os.Open("10014.in")
	defer in.Close()
	out, _ := os.Create("10014.out")
	defer out.Close()

	var kase int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d", &n)
		fmt.Fscanf(in, "%f", &af)
		fmt.Fscanf(in, "%f", &al)
		c := make([]float64, n)
		for i := range c {
			fmt.Fscanf(in, "%f", &c[i])
		}
		fmt.Fprintf(out, "%.2f\n", solve(c))
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
