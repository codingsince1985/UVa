// UVa 353 - Pesky Palindromes

package main

import (
	"fmt"
	"os"
)

func isPalindrome(word string) bool {
	for i := 0; i <= len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func solve(line string) int {
	cache := make(map[string]bool)
	var count int
	for i := range line {
		for j := i + 1; j <= len(line); j++ {
			if !cache[line[i:j]] && isPalindrome(line[i:j]) {
				cache[line[i:j]] = true
				count++
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("353.in")
	defer in.Close()
	out, _ := os.Create("353.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		fmt.Fprintf(out, "The string '%s' contains %d palindromes.\n", line, solve(line))
	}
}
