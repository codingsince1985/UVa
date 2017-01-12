// UVa 127 - "Accordian" Patience

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct{ rank, suit byte }

func matched(c1, c2 card) bool { return c1.rank == c2.rank || c1.suit == c2.suit }

func cleanPiles(piles [][]card, i int) [][]card {
	if len(piles[i]) == 0 {
		if i == len(piles)-1 {
			return piles[:i]
		}
		return append(piles[:i], piles[i+1:]...)
	}
	return piles
}

func deal(cs []string) (int, []string) {
	piles := make([][]card, len(cs))
	for i, c := range cs {
		piles[i] = []card{{c[0], c[1]}}
	}
	for i := 0; i < len(piles); i++ {
		if i-3 >= 0 && matched(piles[i-3][len(piles[i-3])-1], piles[i][len(piles[i])-1]) {
			piles[i-3] = append(piles[i-3], piles[i][len(piles[i])-1])
			piles[i] = piles[i][:len(piles[i])-1]
			piles = cleanPiles(piles, i)
			i = 0
		} else if i-1 >= 0 && matched(piles[i-1][len(piles[i-1])-1], piles[i][len(piles[i])-1]) {
			piles[i-1] = append(piles[i-1], piles[i][len(piles[i])-1])
			piles[i] = piles[i][:len(piles[i])-1]
			piles = cleanPiles(piles, i)
			i = 0
		}
	}
	height := make([]string, len(piles))
	for i := range piles {
		height[i] = strconv.Itoa(len(piles[i]))
	}
	return len(piles), height
}

func main() {
	in, _ := os.Open("127.in")
	defer in.Close()
	out, _ := os.Create("127.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		s.Scan()
		num, height := deal(strings.Split(line+" "+s.Text(), " "))
		fmt.Fprintf(out, "%d piles remaining: %s\n", num, strings.Join(height, " "))
	}
}
