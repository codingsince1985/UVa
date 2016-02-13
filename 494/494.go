// UVa 494 - Kindergarten Counting Game

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("494.in")
	defer in.Close()
	out, _ := os.Create("494.out")
	defer out.Close()

	var c byte
	var count int
	var word bool
	for {
		if _, err := fmt.Fscanf(in, "%c", &c); err != nil {
			break
		}
		if c == '\n' {
			fmt.Fprintln(out, count)
			count = 0
		}
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			if !word {
				word = true
				count++
			}
		} else {
			word = false
		}
	}
}
