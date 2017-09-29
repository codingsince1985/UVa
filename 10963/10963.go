// UVa 10963 - The Swallowing Ground

package main

import (
	"fmt"
	"os"
)

func solve(w int, gaps [][2]int) bool {
	diff := gaps[0][0] - gaps[0][1]
	for i := 2; i < w; i++ {
		if diff != gaps[i][0]-gaps[i][1] {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10963.in")
	defer in.Close()
	out, _ := os.Create("10963.out")
	defer out.Close()

	var kase, w int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d", &w)
		gaps := make([][2]int, w)
		for i := range gaps {
			fmt.Fscanf(in, "%d%d", &gaps[i][0], &gaps[i][1])
		}
		if solve(w, gaps) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
