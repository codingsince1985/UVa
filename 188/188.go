// UVa 188 - Perfect Hash

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calc(w []int) int {
	c, l := w[0], len(w)
	for {
		ok := true
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				if c/w[i]%l == c/w[j]%l {
					c = min((c/w[i]+1)*w[i], (c/w[j]+1)*w[j])
					ok = false
				}
			}
		}
		if ok {
			break
		}
	}
	return c
}

func toNumber(word string) int {
	var num int
	l := len(word)
	for i := range word {
		num += int(math.Pow(2, float64((l-i-1)*5))) * int(word[i]-'a'+1)
	}
	return num
}

func solve(line string) int {
	words := strings.Fields(line)
	w := make([]int, len(words))
	for i, word := range words {
		w[i] = toNumber(word)
	}
	sort.Ints(w)
	return calc(w)
}

func main() {
	in, _ := os.Open("188.in")
	defer in.Close()
	out, _ := os.Create("188.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		line = s.Text()
		fmt.Fprintln(out, line)
		fmt.Fprintf(out, "%d\n\n", solve(line))
	}
}
