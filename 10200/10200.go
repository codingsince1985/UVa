// UVa 10200 - Prime Time

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

func main() {
	in, _ := os.Open("10200.in")
	defer in.Close()
	out, _ := os.Create("10200.out")
	defer out.Close()

	var a, b int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &b); err != nil {
			break
		}
		cnt := 0
		for i := a; i <= b; i++ {
			if isPrime(i*i + i + 41) {
				cnt++
			}
		}
		fmt.Fprintf(out, "%.2f\n", float64(cnt)*100/float64(b-a+1))
	}
}
