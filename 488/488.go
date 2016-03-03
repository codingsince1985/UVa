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
	fmt.Fscanf(in, "%d\n", &n)
	for n > 0 {
		if _, err := fmt.Fscanf(in, "\n%d\n%d", &a, &f); err != nil {
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
		for f > 0 {
			if !first {
				fmt.Fprintln(out)
			} else {
				first = false
			}
			fmt.Fprint(out, s)
			f--
		}
		n--
	}
}
