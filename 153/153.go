// UVa 153 - Permalex

package main

import (
	"fmt"
	"math/big"
	"os"
)

func factorial(a int) *big.Int {
	f := big.NewInt(1)
	for ; a > 1; a-- {
		f.Mul(f, big.NewInt(int64(a)))
	}
	return f
}

func calc(a, b int) *big.Int {
	f := factorial(a)
	f.Div(f, factorial(b))
	return f
}

func solve(word string) *big.Int {
	if len(word) <= 1 {
		return big.NewInt(1)
	}

	smallerMap := make(map[byte]bool)
	for i := 1; i < len(word); i++ {
		if word[0] > word[i] {
			smallerMap[word[i]] = true
		}
	}
	charMap := make(map[byte]int)
	for i := 1; i < len(word); i++ {
		charMap[word[i]]++
	}

	total := big.NewInt(0)
	for i := range smallerMap {
		dup := 0
		for j, v := range charMap {
			if j == i {
				v--
			}
			if v > 1 {
				dup += v
			}
		}
		total.Add(total, calc(len(word)-1, dup))
	}
	return total.Add(total, solve(word[1:]))
}

func main() {
	in, _ := os.Open("153.in")
	defer in.Close()
	out, _ := os.Create("153.out")
	defer out.Close()

	var word string
	for {
		if fmt.Fscanf(in, "%s", &word); word == "#" {
			break
		}
		fmt.Fprintf(out, "%10v\n", solve(word))
	}
}
