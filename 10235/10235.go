// UVa 10235 - Simply Emirp

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverse(n int) int {
	r := 0
	for n != 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func main() {
	in, _ := os.Open("10235.in")
	defer in.Close()
	out, _ := os.Create("10235.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if fmt.Fprintf(out, "%d ", n); isPrime(n) {
			if r := reverse(n); n != r && isPrime(r) {
				fmt.Fprintln(out, "is emirp.")
			} else {
				fmt.Fprintln(out, "is prime.")
			}
		} else {
			fmt.Fprintln(out, "is not prime.")
		}
	}
}
