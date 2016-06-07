// UVa 10931 - Parity

package main

import (
	"fmt"
	"os"
	"strconv"
)

func binary(n int) (string, int) {
	str, cnt := "", 0
	for n > 0 {
		if n%2 == 1 {
			cnt++
		}
		str = strconv.Itoa(n%2) + str
		n /= 2
	}
	return str, cnt
}

func main() {
	in, _ := os.Open("10931.in")
	defer in.Close()
	out, _ := os.Create("10931.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		b, o := binary(n)
		fmt.Fprintf(out, "The parity of %s is %d (mod 2).\n", b, o)
	}
}
