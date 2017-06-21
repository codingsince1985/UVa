// UVa 10029 - Edit Step Ladders

package main

import (
	"fmt"
	"os"
)

var dp = make(map[int]map[string]int)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(word string) int {
	var longest int
	for j := 0; j <= len(word); j++ {
		for k := 'a'; k <= 'z'; k++ {
			if newWord := word[:j] + string(k) + word[j:]; dp[len(newWord)][newWord] > 0 && word > newWord {
				longest = max(longest, dp[len(newWord)][newWord])
			}
		}
	}
	return longest
}

func remove(word string) int {
	var longest int
	for j := 0; j < len(word); j++ {
		if newWord := word[:j] + word[j+1:]; dp[len(newWord)][newWord] > 0 && word > newWord {
			longest = max(longest, dp[len(newWord)][newWord])
		}
	}
	return longest
}

func replace(word string) int {
	var longest int
	for j := 0; j < len(word); j++ {
		for k := 'a'; k <= 'z'; k++ {
			if newWord := word[:j] + string(k) + word[j+1:]; dp[len(newWord)][newWord] > 0 && word > newWord {
				longest = max(longest, dp[len(newWord)][newWord])
			}
		}
	}
	return longest
}

func solve(words []string) int {
	var longest int
	for _, word := range words {
		dp[len(word)][word] = max(max(add(word), remove(word)), replace(word)) + 1
		longest = max(longest, dp[len(word)][word])
	}
	return longest
}

func main() {
	in, _ := os.Open("10029.in")
	defer in.Close()
	out, _ := os.Create("10029.out")
	defer out.Close()

	var word string
	var words []string
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		words = append(words, word)
		if dp[len(word)] == nil {
			dp[len(word)] = make(map[string]int)
		}
		dp[len(word)][word] = 1
	}
	fmt.Fprintln(out, solve(words))
}
