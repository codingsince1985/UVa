// UVa 10323 - Factorial! You Must be Kidding!!!

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10323.in")
	defer in.Close()
	out, _ := os.Create("10323.out")
	defer out.Close()

	fact := uint64(1)
	n := 1
	var min, max int
	var facts []uint64
	for {
		fact *= uint64(n)
		if fact >= 10000 && min == 0 {
			min = n
		}
		if fact > 6227020800 && max == 0 {
			max = n - 1
			break
		}
		if min != 0 {
			facts = append(facts, fact)
		}
		n++
	}

	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if n >= 0 {
			switch {
			case n < min:
				fmt.Fprintln(out, "Underflow!")
			case n > max:
				fmt.Fprintln(out, "Overflow!")
			default:
				fmt.Fprintln(out, facts[n-8])
			}
		} else {
			if n%2 == 0 {
				fmt.Fprintln(out, "Underflow!")
			} else {
				fmt.Fprintln(out, "Overflow!")
			}
		}
	}
}
