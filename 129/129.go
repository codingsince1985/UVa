// UVa 129 - Krypton Factor

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	n, l, count int
	out         io.WriteCloser
)

func isEasy(level int, sequence []string) bool {
here:
	for i := 1; 2*i <= level+1; i++ {
		for j := 0; j < i; j++ {
			if sequence[level-j] != sequence[level-j-i] {
				continue here
			}
		}
		return true
	}
	return false
}

func output(sequence []string) {
	for i := range sequence {
		switch {
		case i > 0 && i%64 == 0:
			fmt.Fprintln(out)
		case i > 0 && i%4 == 0:
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, sequence[i])
	}
	fmt.Fprintf(out, "\n%d\n", len(sequence))
}

func dfs(level int, sequence []string) bool {
	if count == n {
		output(sequence)
		return true
	}
	count++
	for i := 0; i < l; i++ {
		if !isEasy(level, append(sequence, string('A'+i))) {
			if dfs(level+1, append(sequence, string('A'+i))) {
				return true
			}
		}
	}
	return false
}

func main() {
	in, _ := os.Open("129.in")
	defer in.Close()
	out, _ = os.Create("129.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%d%d", &n, &l); n == 0 && l == 0 {
			break
		}
		count = 0
		dfs(0, nil)
	}
}
