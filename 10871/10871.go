// UVa 10871 - Primed Subsequence

package main

import (
	"fmt"
	"os"
)

const max = 46340 // √(2^31-1)

var primes = sieve()

func sieve() []bool {
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func isPrime(n int) bool {
	if n <= max {
		return !primes[n]
	}
	for i := range primes {
		if i*i > n {
			break
		}
		if !primes[i] && n%i == 0 {
			return false
		}
	}
	return true
}

func solve(n int, sequence []int) (int, string) {
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + sequence[i-1]
	}
	for length := 2; length <= n; length++ {
		for i := length; i <= n; i++ {
			if isPrime(preSum[i] - preSum[i-length]) {
				str := fmt.Sprint(sequence[i-length : i])
				return length, str[1 : len(str)-1]
			}
		}
	}
	return 0, ""
}

func main() {
	in, _ := os.Open("10871.in")
	defer in.Close()
	out, _ := os.Create("10871.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		sequence := make([]int, n)
		for i := range sequence {
			fmt.Fscanf(in, "%d", &sequence[i])
		}
		if length, subSequence := solve(n, sequence); length == 0 {
			fmt.Fprintln(out, "This sequence is anti-primed.")
		} else {
			fmt.Fprintf(out, "Shortest primed subsequence is length %d: %s\n", length, subSequence)
		}
	}
}
