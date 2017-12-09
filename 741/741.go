// UVa 741 - Burrows Wheeler Decoder

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(line string, p int) string {
	matrix := make([]string, len(line))
	for range matrix {
		for j := range matrix {
			matrix[j] = string(line[j]) + matrix[j]
		}
		sort.Strings(matrix)
	}
	return matrix[p-1]
}

func main() {
	in, _ := os.Open("741.in")
	defer in.Close()
	out, _ := os.Create("741.out")
	defer out.Close()

	var line string
	var p int
	for first := true; ; {
		if fmt.Fscanf(in, "%s\n%d", &line, &p); line == "END" && p == 0 {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, solve(line, p))
	}
}
