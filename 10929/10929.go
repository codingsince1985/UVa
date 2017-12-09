// UVa 10929 - You can say 11

package main

import (
	"fmt"
	"os"
)

func add(n string) int {
	var sum int
	odd := true
	for i := range n {
		if odd {
			sum += int(n[i] - '0')
		} else {
			sum -= int(n[i] - '0')
		}
		odd = !odd
	}
	return sum
}

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
		if fmt.Fprintf(out, "%s ", n); add(n)%11 == 0 {
			fmt.Fprintln(out, "is a multiple of 11.")
		} else {
			fmt.Fprintln(out, "is not a multiple of 11.")
		}
	}
}
