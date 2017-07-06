// UVa 11526 - H(n)

package main

import (
	"fmt"
	"math"
	"os"
)

func h(n int64) int64 {
	var i, result int64
	sqrt := int64(math.Sqrt(float64(n)))
	for i = 1; i <= sqrt; i++ {
		result += i * (n/i - n/(i+1))
	}
	for i = n / (sqrt + 1); i > 0; i-- {
		result += n / i
	}
	return result
}

func main() {
	in, _ := os.Open("11526.in")
	defer in.Close()
	out, _ := os.Create("11526.out")
	defer out.Close()

	var t, n int64
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, h(n))
	}
}
