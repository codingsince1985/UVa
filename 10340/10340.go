// UVa 10340 - All in All

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10340.in")
	defer in.Close()
	out, _ := os.Create("10340.out")
	defer out.Close()

	var s, t string
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &s, &t); err != nil {
			break
		}
		idx, count := 0, 0
	here:
		for _, i := range s {
			for p, j := range t[idx:] {
				if i == j {
					idx = idx + p + 1 // p changes every time, use idx+p
					count++
					continue here
				}
			}
			break
		}
		if len(s) > count {
			fmt.Fprintln(out, "No")
		} else {
			fmt.Fprintln(out, "Yes")
		}
	}
}
