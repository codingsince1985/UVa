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
		seed, cnt := 0, 1
		appeared[0] = true
		for {
			seed = (seed + step) % mod
			if appeared[seed] {
				break
			}
			cnt++
			appeared[seed] = true
		}
		fmt.Fprintf(out, "%10d%10d    ", step, mod)
		if cnt == mod {
			fmt.Fprintf(out, "Good Choice\n\n")
		} else {
			fmt.Fprintf(out, "Bad Choice\n\n")
		}
	}
}
