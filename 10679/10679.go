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
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "%s", &s)
		fmt.Fscanf(in, "%d", &n)
		for n > 0 {
			fmt.Fscanf(in, "%s", &t)
			// may need Aho–Corasick algorithm, but ah well
			if strings.Contains(s, t) {
				fmt.Fprintln(out, "y")
			} else {
				fmt.Fprintln(out, "n")
			}
			n--
		}
		kase--
	}
}
