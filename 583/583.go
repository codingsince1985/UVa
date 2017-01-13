// UVa 583 - Prime Factors

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func factorize(n int) []string {
	var p []string

	if n < 0 {
		p = append(p, "-1")
		n = -n
	}

	t := 2
here:
	for {
		for n%t == 0 {
			p = append(p, strconv.Itoa(t))
			if n /= t; n == 1 {
				break here
			}
		}

		if t++; t*t > n {
			p = append(p, strconv.Itoa(n))
			break
		}
	}
	return p
}

func main() {
	in, _ := os.Open("583.in")
	defer in.Close()
	out, _ := os.Create("583.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		p := factorize(n)
		fmt.Fprintf(out, "%d = ", n)
		fmt.Fprintln(out, strings.Join(p, " x "))
	}
}
