// UVa 484 - The Department of Redundancy Department

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("484.in")
	defer in.Close()
	out, _ := os.Create("484.out")
	defer out.Close()

	var n int
	var lst []int
	dict := make(map[int]int)
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if _, ok := dict[n]; !ok {
			lst = append(lst, n)
		}
		dict[n] += 1
	}
	for _, v := range lst {
		fmt.Fprintln(out, v, dict[v])
	}
}
