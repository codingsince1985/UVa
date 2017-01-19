// UVa 10508 - Word Morphing

package main

import (
	"fmt"
	"os"
)

func getDistance(w1, w2 string) int {
	var count int
	for i := range w1 {
		if w1[i] != w2[i] {
			count++
		}
	}
	return count
}

func solve(words []string) map[int]int {
	distanceMap := make(map[int]int)
	for i := 2; i < len(words); i++ {
		distanceMap[getDistance(words[0], words[i])] = i
	}
	return distanceMap
}

func main() {
	in, _ := os.Open("10508.in")
	defer in.Close()
	out, _ := os.Create("10508.out")
	defer out.Close()

	var ws, ls int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &ws, &ls); err != nil {
			break
		}
		words := make([]string, ws)
		for i := range words {
			fmt.Fscanf(in, "%s", &words[i])
		}
		fmt.Fprintln(out, words[0])
		distanceMap := solve(words)
		for i := 1; i < ls; i++ {
			fmt.Fprintln(out, words[distanceMap[i]])
		}
		fmt.Fprintln(out, words[1])
	}
}
