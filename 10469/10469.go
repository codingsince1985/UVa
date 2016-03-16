// UVa 10469 - To Carry or not to Carry

package main

import (
	"fmt"
	"os"
)

func add(a, b uint32) uint32 {
	var total, base uint32
	base = 1
	for a != 0 || b != 0 {
		a1, b1 := a%2, b%2
		if a1 != b1 {
			total += base
		}
		base *= 2
		a /= 2
		b /= 2
	}
	return total
}

func main() {
	in, _ := os.Open("10469.in")
	defer in.Close()
	out, _ := os.Create("10469.out")
	defer out.Close()

	var a, b uint32
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &b); err != nil {
			break
		}
		fmt.Fprintln(out, a^b)
	}
}
