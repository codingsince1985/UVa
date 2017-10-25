// UVa 408 - Uniform Generator

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("408.in")
	defer in.Close()
	out, _ := os.Create("488.out")
	defer out.Close()

	var step, mod int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &step, &mod); err != nil {
			break
		}
		appeared := make([]bool, mod)
		appeared[0] = true
		seed, cnt := 0, 1
		for {
			seed = (seed + step) % mod
			if appeared[seed] {
				break
			}
			cnt++
			appeared[seed] = true
		}
		if fmt.Fprintf(out, "%10d%10d    ", step, mod); cnt == mod {
			fmt.Fprintln(out, "Good Choice")
		} else {
			fmt.Fprintln(out, "Bad Choice")
		}
		fmt.Fprintln(out)
	}
}
