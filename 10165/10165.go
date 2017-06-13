// UVa 10165 - Stone Game

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10165.in")
	defer in.Close()
	out, _ := os.Create("10165.out")
	defer out.Close()

	var n, pile int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var sum int
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d", &pile)
			sum ^= pile
		}
		if sum != 0 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}
