// UVa 392 - Polynomial Showdown

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

func output(out *os.File, coefficients []int) {
	first := true
	for i := range coefficients {
		if coefficients[i] == 0 {
			continue
		}
		if !first {
			if coefficients[i] > 0 {
				fmt.Fprint(out, " + ")
			} else {
				fmt.Fprint(out, " - ")
			}
		}
		if first && coefficients[i] < 0 {
			fmt.Fprint(out, coefficients[i])
		} else if abs(coefficients[i]) != 1 {
			fmt.Fprint(out, abs(coefficients[i]))
		}
		switch {
		case i == 7:
			fmt.Fprint(out, "x")
		case i == 8:
			fmt.Fprint(out, abs(coefficients[i]))
		default:
			fmt.Fprint(out, "x^", 8-i)
		}
		first = false
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("392.in")
	defer in.Close()
	out, _ := os.Create("392.out")
	defer out.Close()

here:
	for {
		coefficients := make([]int, 9)
		for i := range coefficients {
			if _, err := fmt.Fscanf(in, "%d", &coefficients[i]); err != nil {
				break here
			}
		}
		output(out, coefficients)
	}
}
