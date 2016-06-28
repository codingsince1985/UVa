// UVa 455 - Periodic Strings

package main

import (
	"fmt"
	"os"
)

func factors(n int) []int {
	var f []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func periodic(line string, k int) bool {
	token := make(map[string]bool)
	n := len(line) / k
	for i := 0; i < n; i++ {
		token[line[i*k:i*k+k]] = true
		if len(token) > 1 {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("455.in")
	defer in.Close()
	out, _ := os.Create("455.out")
	defer out.Close()

	var n int
	var line string
	first := true
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		n--
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%s", &line)
		ks := factors(len(line))
		for _, k := range ks {
			if periodic(line, k) {
				fmt.Fprintln(out, k)
				break
			}
		}
	}
}
