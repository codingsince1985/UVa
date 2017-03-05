// UVa 488 - Triangle Wave

package main

import (
	"fmt"
	"os"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	in, _ := os.Open("488.in")
	defer in.Close()
	out, _ := os.Create("488.out")
	defer out.Close()

	var n, a, f int
	first := true
	for fmt.Fscanf(in, "%d\n", &n); n > 0; n-- {
		fmt.Fscanln(in)
		if _, err := fmt.Fscanf(in, "%d", &a); err != nil {
			break
		}
		var s string
		for i := 0; i < 2*a-1; i++ {
			toPrint := a - abs(a-i-1)
			for j := 0; j < toPrint; j++ {
				s += fmt.Sprint(toPrint)
			}
			s += "\n"
		}
		for fmt.Fscanf(in, "%d", &f); f > 0; f-- {
			if first {
				first = false
			} else {
				fmt.Fprintln(out)
			}
			fmt.Fprint(out, s)
		}
	}
}
