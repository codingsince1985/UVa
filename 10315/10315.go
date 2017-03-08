// UVa 10315 - Poker Hands

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type card struct {
	value int
	suit  byte
}

var valueMap = map[byte]int{'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

func straightFlush(c []card) bool { return straight(c) && flush(c) }

func fourOfAKind(c []card) bool { return c[0].value == c[3].value || c[1].value == c[4].value }

func fullHouse(c []card) bool { return threeOfAKind(c) && getValues(c) == 2 }

func flush(c []card) bool {
	for i := 1; i < 5; i++ {
		if c[i].suit != c[0].suit {
			return false
		}
	}
	return true
}

func straight(c []card) bool {
	for i := 1; i < 5; i++ {
		if c[i].value != c[0].value-i {
			return false
		}
	}
	return true
}

func threeOfAKind(c []card) bool {
	return c[0].value == c[2].value || c[1].value == c[3].value || c[2].value == c[4].value
}

func twoPairs(c []card) bool { return getValues(c) == 3 }

func pair(c []card) bool { return getValues(c) == 4 }

func getValues(c []card) int {
	valueMap := make(map[int]bool)
	for _, vi := range c {
		valueMap[vi.value] = true
	}
	return len(valueMap)
}

func buildCard(c string) card {
	if v, ok := valueMap[c[0]]; ok {
		return card{v, c[1]}
	}
	return card{int(c[0] - '0'), c[1]}
}

func score(c []card) int {
	sort.Slice(c, func(i, j int) bool { return c[i].value > c[j].value })
	// order matters
	switch {
	case straightFlush(c):
		return 9
	case fourOfAKind(c):
		return 8
	case fullHouse(c):
		return 7
	case flush(c):
		return 6
	case straight(c):
		return 5
	case threeOfAKind(c):
		return 4
	case twoPairs(c):
		return 3
	case pair(c):
		return 2
	default:
		return 1
	}
}

func highCard(blacks, whites []card) (int, int) {
	for i := range blacks {
		if blacks[i].value != whites[i].value {
			return blacks[i].value, whites[i].value
		}
	}
	return 0, 0
}

func scoreForPair(c []card) []int {
	switch {
	case c[0].value == c[1].value:
		return []int{c[0].value, c[2].value, c[3].value, c[4].value}
	case c[1].value == c[2].value:
		return []int{c[1].value, c[0].value, c[3].value, c[4].value}
	case c[2].value == c[3].value:
		return []int{c[2].value, c[0].value, c[1].value, c[4].value}
	default:
		return []int{c[3].value, c[0].value, c[1].value, c[2].value}
	}
}

func scoreForTwoPairs(c []card) []int {
	switch {
	case c[0].value != c[1].value:
		return []int{c[1].value, c[3].value, c[0].value}
	case c[1].value != c[2].value:
		return []int{c[0].value, c[3].value, c[2].value}
	default:
		return []int{c[0].value, c[2].value, c[4].value}
	}
}

func compareSameScore(score int, blacks, whites []card) (int, int) {
	switch score {
	case 9, 5:
		return blacks[0].value, whites[0].value
	case 8:
		return blacks[1].value, whites[1].value
	case 7, 4:
		return blacks[2].value, whites[2].value
	case 6, 1:
		return highCard(blacks, whites)
	case 3, 2:
		var bs, ws []int
		if score == 3 {
			bs, ws = scoreForTwoPairs(blacks), scoreForTwoPairs(whites)
		} else {
			bs, ws = scoreForPair(blacks), scoreForPair(whites)
		}
		for i := range bs {
			if bs[i] != ws[i] {
				return bs[i], ws[i]
			}
		}
		fallthrough
	default:
		return 0, 0
	}
}

func compare(line string) string {
	c := strings.Split(line, " ")
	blacks, whites := make([]card, 5), make([]card, 5)
	for i := 0; i < 5; i++ {
		blacks[i], whites[i] = buildCard(c[i]), buildCard(c[i+5])
	}
	blackScore, whiteScore := score(blacks), score(whites)
	if blackScore == whiteScore {
		blackScore, whiteScore = compareSameScore(whiteScore, blacks, whites)
	}
	switch {
	case blackScore > whiteScore:
		return "Black wins."
	case whiteScore > blackScore:
		return "White wins."
	default:
		return "Tie."
	}
}

func main() {
	in, _ := os.Open("10315.in")
	defer in.Close()
	out, _ := os.Create("10315.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		fmt.Fprintln(out, compare(line))
	}
}
