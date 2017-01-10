// UVa 406 - Prime Cuts

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
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

func getPrimes() []int {
	primes := make([]int, 2)
	primes[0] = 1
	primes[1] = 2
	for i := 3; i < 1000; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func output(out *os.File, lst []int) {
	for _, v := range lst {
		fmt.Fprintf(out, " %d", v)
	}
	fmt.Fprintf(out, "\n\n")
}

func main() {
	in, _ := os.Open("406.in")
	defer in.Close()
	out, _ := os.Create("406.out")
	defer out.Close()

	primes := getPrimes()
	var n, c, i int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &c); err != nil {
			break
		}
		for i = 0; i < len(primes); i++ {
			if primes[i] > n {
				break
			}
		}
		fmt.Fprintf(out, "%d %d:", n, c)
		var toPrint []int
		if size := 2*c - i%2; size >= i {
			toPrint = primes[0:i]
		} else {
			edge := (i - size) / 2
			toPrint = primes[edge : edge+size]
		}
		output(out, toPrint)
	}
}
