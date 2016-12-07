// UVa 10633 - Rare Easy Problem

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10633.in")
	defer in.Close()
	out, _ := os.Create("10633.out")
	defer out.Close()

	var nm uint64
	for {
		if fmt.Fscanf(in, "%d", &nm); nm == 0 {
			break
		}
		if n := (nm * 10) / 9; nm%9 == 0 {
			fmt.Fprintln(out, n-1, n)
		} else {
			fmt.Fprintln(out, n)
		}
	}
}
