// UVa 10931 - Parity

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binary(n int64) (string, int) {
	bin := strconv.FormatInt(n, 2)
	return bin, strings.Count(bin, "1")
}

func main() {
	in, _ := os.Open("10931.in")
	defer in.Close()
	out, _ := os.Create("10931.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		b, o := binary(n)
		fmt.Fprintf(out, "The parity of %s is %d (mod 2).\n", b, o)
	}
}
