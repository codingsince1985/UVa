// UVa 10019 - Funny Encryption Method

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func onesInBinary(n int64) int { return strings.Count(strconv.FormatInt(n, 2), "1") }

func hexaDec(n int64) int64 {
	var base int64 = 1
	var ret int64 = 0
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

	var n, dec int64
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &dec)
		fmt.Fprintln(out, onesInBinary(dec), onesInBinary(hexaDec(dec)))
	}
}
