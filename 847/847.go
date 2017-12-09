// UVa 847 - A Multiplication Game

package main

import (
	"fmt"
	"os"
)

func stanWins(n int64) bool {
	win := true
	for min, max := int64(1), int64(9); ; {
		if n >= min && n <= max {
			break
		}
		min = max + 1
		if win {
			max *= 2
		} else {
			max *= 9
		}
		win = !win
	}
	return win
}

func main() {
	in, _ := os.Open("847.in")
	defer in.Close()
	out, _ := os.Create("847.out")
	defer out.Close()

	var n int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if stanWins(n) {
			fmt.Fprintln(out, "Stan wins.")
		} else {
			fmt.Fprintln(out, "Ollie wins.")
		}
	}
}
