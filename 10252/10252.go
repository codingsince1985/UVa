// UVa 10252 - Common Permutation

package main

import (
	"fmt"
	"os"
)

func charMap(s string) map[byte]int {
	cm := make(map[byte]int)
	for i := range s {
		cm[s[i]]++
	}
	return cm
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonPermutation(a, b string) map[byte]int {
	cb := charMap(b)
	cm := make(map[byte]int)
	for k, va := range charMap(a) {
		if vb, ok := cb[k]; ok {
			cm[k] = min(va, vb)
		}
	}
	return cm
}

func main() {
	in, _ := os.Open("10252.in")
	defer in.Close()
	out, _ := os.Create("10252.out")
	defer out.Close()

	var a, b string
	for {
		if _, err := fmt.Fscanf(in, "%s\n%s", &a, &b); err != nil {
			break
		}
		cm := commonPermutation(a, b)
		var i byte
		for i = 'a'; i <= 'z'; i++ {
			if v, ok := cm[i]; ok {
				for ; v > 0; v-- {
					fmt.Fprintf(out, "%c", i)
				}
			}
		}
		fmt.Fprintln(out)
	}
}
