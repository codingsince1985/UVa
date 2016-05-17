// UVa 417 - Word Index

package main

import (
	"fmt"
	"os"
)

func dp() [27][6]int {
	var cache [27][6]int
	for i := range cache {
		cache[i][0] = 1
		if i == 0 {
			continue
		}
		for j := 1; j < len(cache[i]); j++ {
			cache[i][j] = cache[i-1][j-1] + cache[i-1][j]
		}
	}
	return cache
}

func main() {
	in, _ := os.Open("417.in")
	defer in.Close()
	out, _ := os.Create("417.out")
	defer out.Close()

	cache := dp()
	var word string
here:
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		for i := 0; i < len(word)-1; i++ {
			if word[i] >= word[i+1] {
				fmt.Fprintln(out, 0)
				continue here
			}
		}
		total := 1
		for i := 1; i < len(word); i++ {
			total += cache[26][i]
		}
		var ch byte = 'a'
		for i := 0; i < len(word); i++ {
			for ; ch < word[i]; ch++ {
				total += cache['z'-ch][len(word)-i-1]
			}
			ch++
		}
		fmt.Fprintln(out, total)
	}
}
