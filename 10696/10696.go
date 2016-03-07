// UVa 10696 - f91

package main

import (
	"fmt"
	"os"
)

func f91(n int) int {
	if n <= 100 {
		return f91(f91(n + 11))
	}
	return n - 10
}

func main() {
	in, _ := os.Open("10696.in")
	defer in.Close()
	out, _ := os.Create("10696.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "f91(%d) = %d\n", n, f91(n))
	}
}
