// UVa 10469 - To Carry or not to Carry

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10469.in")
	defer in.Close()
	out, _ := os.Create("10469.out")
	defer out.Close()

	var a, b uint32
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &a, &b); err != nil {
			break
		}
		fmt.Fprintln(out, a^b)
	}
}
