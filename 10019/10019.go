// UVa 10019 - Funny Encryption Method

package main

import (
	"fmt"
	"os"
)

func binary(n int) int {
	cnt := 0
	for n > 0 {
		if n%2 == 1 {
			cnt++
		}
		n /= 2
	}
	return cnt
}

func hexaDec(n int) int {
	base := 1
	ret := 0
	for n > 0 {
		ret += (n % 10) * base
		n /= 10
		base *= 16
	}
	return ret
}

func main() {
	in, _ := os.Open("10019.in")
	defer in.Close()
	out, _ := os.Create("10019.out")
	defer out.Close()

	var n, dec int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &dec)
		fmt.Fprintln(out, binary(dec), binary(hexaDec(dec)))
	}
}
