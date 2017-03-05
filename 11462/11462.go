// UVa 11462 - Age Sort

package main

import (
	"fmt"
	"os"
)

var out *os.File

func output(ages [100]int) {
	first := true
	for i, v := range ages {
		for ; v > 0; v-- {
			if first {
				first = false
			} else {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, i)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("11462.in")
	defer in.Close()
	out, _ = os.Create("11462.out")
	defer out.Close()

	var n, age int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var ages [100]int
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d", &age)
			ages[age]++
		}
		output(ages)
	}
}
