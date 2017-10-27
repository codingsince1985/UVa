// UVa 12468 - Zapping

package main

import (
	"fmt"
	"os"
)

func solve(a, b int) int {
	fmt.Println(a, b)
	var count int
	if a < b {
		count = b - a
	} else {
		count = a - b
	}
	if count > 50 {
		count = 100 - count
	}
	return count
}

func main() {
	in, _ := os.Open("12468.in")
	defer in.Close()
	out, _ := os.Create("12468.out")
	defer out.Close()

	var a, b int
	for {
		if fmt.Fscanf(in, "%d%d", &a, &b); a == -1 && b == -1 {
			break
		}
		fmt.Fprintln(out, solve(a, b))
	}
}
