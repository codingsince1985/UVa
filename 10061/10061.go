// UVa 10061 - How many zero's and how many digits ?

package main

import (
	"fmt"
	"math"
	"os"
)

const zero = 0.00001

func digits(n, b int) int {
	var d float64
	for i := 1; i <= n; i++ {
		d += math.Log10(float64(i))
		d /= math.Log10(float64(b))
	}
	return int(d+zero) + 1
}

func factorize(n, b int) []int {
	factor := make([]int, b+1)
	for i := 2; i <= n; i++ {
		for j, temp := 2, i; j <= temp && j <= b; j++ {
			for temp%j == 0 {
				factor[j]++
				temp /= j
			}
		}
	}
	return factor
}

func zeros(n, b int) int {
	factor := factorize(n, b)
	var count int
	for ; ; count++ {
		temp := b
		for i := 2; i <= b; i++ {
			for factor[i] > 0 && temp%i == 0 {
				factor[i]--
				temp /= i
			}
		}
		if temp != 1 {
			break
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10061.in")
	defer in.Close()
	out, _ := os.Create("10061.out")
	defer out.Close()

	var n, b int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &b); err != nil {
			break
		}
		fmt.Fprintln(out, zeros(n, b), digits(n, b))
	}
}
