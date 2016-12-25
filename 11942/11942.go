// UVa 11942 - Lumberjack Sequencing

package main

import (
	"fmt"
	"os"
)

var beards [10]int

func ordered() bool {
	diff := beards[1] - beards[0]
	for i := 2; i < 10; i++ {
		if (beards[i]-beards[i-1])*diff < 0 {
			return false
		}
		diff = beards[i] - beards[i-1]
	}
	return true
}

func main() {
	in, _ := os.Open("11942.in")
	defer in.Close()
	out, _ := os.Create("11942.out")
	defer out.Close()

	fmt.Fprintln(out, "Lumberjacks:")
	var n int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		for i := range beards {
			fmt.Fscanf(in, "%d", &beards[i])
		}
		if ordered() {
			fmt.Fprintln(out, "Ordered")
		} else {
			fmt.Fprintln(out, "Unordered")
		}
	}
}
