// UVa 10338 - Mischievous Children

package main

import (
	"fmt"
	"os"
)

func factCache() map[int]uint64 {
	cache := make(map[int]uint64)
	var t, i uint64 = 1, 1
	for ; i <= 20; i++ {
		t *= i
		cache[int(i)] = t
	}
	return cache
}

func main() {
	in, _ := os.Open("10338.in")
	defer in.Close()
	out, _ := os.Create("10338.out")
	defer out.Close()

	cache := factCache()
	var n int
	var word string
	fmt.Fscanf(in, "%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%s", &word)
		var charMap = make(map[rune]int)
		for _, c := range word {
			charMap[c]++
		}
		t := cache[len(word)]
		for _, v := range charMap {
			if v > 1 {
				t /= cache[v]
			}
		}
		fmt.Fprintf(out, "Data set %d: %d\n", i, t)
	}
}
