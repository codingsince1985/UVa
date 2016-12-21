// UVa 10679 - I Love Strings!!

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("10679.in")
	defer in.Close()
	out, _ := os.Create("10679.out")
	defer out.Close()

	var kase, n int
	var s, t string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &s)
		for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
			fmt.Fscanf(in, "%s", &t)
			// may need Aho–Corasick algorithm, but ah well
			if strings.Contains(s, t) {
				fmt.Fprintln(out, "y")
			} else {
				fmt.Fprintln(out, "n")
			}
		}
	}
}
