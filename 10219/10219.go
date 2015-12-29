// UVa 10219 - Find the ways !

// c(n, k) = n! / (k!(n-k)!)
// log(n!) = log(n) + log(n-1) + log(n-2) + ... + log(1)

package main

import (
	"fmt"
	"math"
	"os"
)

func digitsOfFactorial(n int) float64 {
	var t float64
	for i := 1; i <= n; i++ {
		t += math.Log10(float64(i))
	}
	return t
}

func main() {
	in, _ := os.Open("10219.in")
	defer in.Close()
	out, _ := os.Create("10219.out")
	defer out.Close()

	var n, k int
	for {
		_, err := fmt.Fscanf(in, "%d %d", &n, &k)
		if err != nil {
			break
		}
		result := digitsOfFactorial(n) - (digitsOfFactorial(k) + digitsOfFactorial(n - k))
		fmt.Fprintln(out, int(result) + 1)
	}
}
