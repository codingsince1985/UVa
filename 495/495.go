// UVa 495 - Fibonacci Freeze

package main

import (
	"fmt"
	"os"
)

const MAX_DIGITS = 1050

var f [5001][MAX_DIGITS]int

func prepare() {
	f[1][0] = 1;
	for i := 2; i <= 5000; i ++ {
		for j := 0; j < MAX_DIGITS - 2; j ++ {
			f[i][j] += f[i - 1][j] + f[i - 2][j]
			f[i][j + 1] += f[i][j] / 10
			f[i][j] %= 10
		}
	}
}

func main() {
	prepare()

	in, _ := os.Open("495.in")
	out, _ := os.Create("495.out")
	var n int
	for {
		_, err := fmt.Fscanf(in, "%d", &n)
		if err != nil {
			break;
		}
		var skip = true
		fmt.Fprintf(out, "The Fibonacci number for %d is ", n)
		for j := MAX_DIGITS - 1; j >= 0; j -- {
			tmp := f[n][j]
			if tmp != 0 {
				skip = false
			}
			if !skip {
				fmt.Fprintf(out, "%d", tmp)
			}
		}
		fmt.Fprintln(out)
	}
}
