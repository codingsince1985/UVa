// UVa 616 - Coconuts, Revisited

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(n int) int {
here:
	for people := int(math.Sqrt(float64(n)-1)) + 1; people >= 2; people-- {
		coconuts := n
		for j := 0; j < people; j++ {
			if coconuts--; coconuts < 0 || coconuts%people != 0 {
				continue here
			}
			coconuts -= coconuts / people
		}
		if coconuts >= 0 && coconuts%people == 0 {
			return people
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("616.in")
	defer in.Close()
	out, _ := os.Create("616.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		fmt.Fprintf(out, "%d coconuts, ", n)
		if people := solve(n); people == -1 {
			fmt.Fprintln(out, "no solution")
		} else {
			fmt.Fprintf(out, "%d people and 1 monkey\n", people)
		}
	}
}
