// UVa 12289 - One-Two-Three

package main

import (
	"fmt"
	"os"
)

const one = "one"

func solve(word string) int {
	if len(word) == 5 {
		return 3
	}
	var diff int
	for i := range one {
		if word[i] != one[i] {
			diff++
		}
	}
	if diff == 1 {
		return 1
	}
	return 2
}

func main() {
	in, _ := os.Open("12289.in")
	defer in.Close()
	out, _ := os.Create("12289.out")
	defer out.Close()

	var kase int
	var word string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &word)
		fmt.Fprintln(out, solve(word))
	}
}
