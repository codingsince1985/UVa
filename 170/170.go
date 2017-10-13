// UVa 170 - Clock Patience

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type card struct{ rank, suite byte }

var rankMap = map[byte]int{'A': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'J': 10, 'Q': 11, 'K': 12}

func solve(piles [][]card) (int, card) {
	var last card
	var exposed int
	for idx := 12; len(piles[idx]) > 0; {
		last = piles[idx][0]
		piles[idx] = piles[idx][1:]
		idx = rankMap[last.rank]
		exposed++
	}
	return exposed, last
}

func deal(cards []card) [][]card {
	piles := make([][]card, 13)
	for i := len(cards) - 1; i >= 0; i-- {
		idx := (len(cards) - 1 - i) % 13
		piles[idx] = append([]card{{cards[i].rank, cards[i].suite}}, piles[idx]...)
	}
	return piles
}

func main() {
	in, _ := os.Open("170.in")
	defer in.Close()
	out, _ := os.Create("170.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line, token string
	var cards []card
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		for r := strings.NewReader(line); ; {
			if _, err := fmt.Fscanf(r, "%s", &token); err != nil {
				break
			}
			cards = append(cards, card{token[0], token[1]})
		}
		if len(cards) == 52 {
			exposed, last := solve(deal(cards))
			fmt.Fprintf(out, "%2d,%c%c\n", exposed, last.rank, last.suite)
			cards = nil
		}
	}
}
