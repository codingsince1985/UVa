// UVa 847 - A Multiplication Game

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("847.in")
	defer in.Close()
	out, _ := os.Create("847.out")
	defer out.Close()

	var n int64
	var min, max int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		stanWin := true
		min, max = 1, 9
		for {
			if n >= min && n <= max {
				break
			}
			min = max + 1
			if stanWin {
				max *= 2
			} else {
				max *= 9
			}
			stanWin = !stanWin
		}
		if stanWin {
			fmt.Fprintln(out, "Stan wins.")
		} else {
			fmt.Fprintln(out, "Ollie wins.")
		}
	}
}
