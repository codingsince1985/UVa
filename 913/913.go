// UVa 913 - Joana and the Odd Numbers

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("913.in")
	defer in.Close()
	out, _ := os.Create("913.out")
	defer out.Close()

	var n uint64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, (n+1)*(n+1)*3/2-9)
	}
}
