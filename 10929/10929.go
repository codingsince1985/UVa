// UVa 10929 - You can say 11

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10929.in")
	defer in.Close()
	out, _ := os.Create("10929.out")
	defer out.Close()

	var n string
	for {
		if fmt.Fscanf(in, "%s", &n); n == "0" {
			break
		}
		odd := true
		addsum := 0
		for i := range n {
			if odd {
				addsum += int(n[i] - '0')
			} else {
				addsum -= int(n[i] - '0')
			}
			odd = !odd
		}
		if addsum%11 == 0 {
			fmt.Fprintf(out, "%s is a multiple of 11.\n", n)
		} else {
			fmt.Fprintf(out, "%s is not a multiple of 11.\n", n)
		}
	}
}
