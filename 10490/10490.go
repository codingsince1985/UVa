// UVa 10490 - Mr. Azad and his Son!!!!!

package main

import (
	"fmt"
	"math"
	"os"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10490.in")
	defer in.Close()
	out, _ := os.Create("10490.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		if isPrime(n) {
			if mersenne := int(math.Pow(2, float64(n))) - 1; isPrime(mersenne) {
				fmt.Fprintf(out, "Perfect: %d!\n", mersenne*int(math.Pow(2, float64(n-1))))
			} else {
				fmt.Fprintln(out, "Given number is prime. But, NO perfect number is available.")
			}
		} else {
			fmt.Fprintln(out, "Given number is NOT prime! NO perfect number is available.")
		}
	}
}
