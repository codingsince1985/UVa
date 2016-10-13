// UVa 10170 - The Hotel with Infinite Rooms

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10170.in")
	defer in.Close()
	out, _ := os.Create("10170.out")
	defer out.Close()

	var s, d int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &s, &d); err != nil {
			break
		}
		for day := s; day < d; day += s {
			s++
		}
		fmt.Fprintln(out, s)
	}
}
