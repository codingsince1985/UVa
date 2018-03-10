// UVa 131 - The Psychic Poker Player

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const cardNumber = 5

type card struct {
	value int
	suit  byte
}

var (
	valueMap   = map[byte]int{'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	hand, deck []card
	highest    int
	bestHand   string
)

func doSort(c []card) { sort.Slice(c, func(i, j int) bool { return c[i].value > c[j].value }) }

func straightFlush(c []card) bool { return straight(c) && flush(c) }

func fourOfAKind(c []card) bool { return c[0].value == c[3].value || c[1].value == c[4].value }

func fullHouse(c []card) bool { return threeOfAKind(c) && getValues(c) == 2 }

func flush(c []card) bool {
	for i := 1; i < cardNumber; i++ {
		if c[i].suit != c[0].suit {
			return false
		}
	}
	return true
}

func testStraight(c []card) bool {
	for i := 1; i < cardNumber; i++ {
		if c[i].value != c[0].value-i {
			return false
		}
	}
	return true
}

func straight(c []card) bool { return testStraight(c) || testStraight(aceAsOne(c)) }

func aceAsOne(c []card) []card {
	newC := make([]card, len(c))
	copy(newC, c)
	for i := range newC {
		if newC[i].value == 14 {
			newC[i].value = 1
		}
	}
	doSort(newC)
	return newC
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

func score(c []card) (int, string) {
	doSort(c)
	// order matters
	switch {
	case straightFlush(c):
		return 9, "straight-flush"
	case fourOfAKind(c):
		return 8, "four-of-a-kind"
	case fullHouse(c):
		return 7, "full-house"
	case flush(c):
		return 6, "flush"
	case straight(c):
		return 5, "straight"
	case threeOfAKind(c):
		return 4, "three-of-a-kind"
	case twoPairs(c):
		return 3, "two-pairs"
	case pair(c):
		return 2, "one-pair"
	default:
		return 1, "highest-card"
	}
}

func buildCards(token []string) []card {
	hand := make([]card, cardNumber)
	for i := range hand {
		hand[i] = buildCard(token[i])
	}
	return hand
}

func dfs(level int, used []bool) {
	if level == cardNumber {
		var newHand []card
		for i := range used {
			if used[i] {
				newHand = append(newHand, hand[i])
			}
		}
		for idx := 0; len(newHand) < cardNumber; idx++ {
			newHand = append(newHand, deck[idx])
		}
		if newScore, newBestHand := score(newHand); newScore > highest {
			highest = newScore
			bestHand = newBestHand
		}
		return
	}
	used[level] = true
	dfs(level+1, used)
	used[level] = false
	dfs(level+1, used)
}

func solve(out io.Writer, line string) {
	token := strings.Fields(line)
	hand = buildCards(token[:cardNumber])
	deck = buildCards(token[cardNumber:])
	highest = 0
	dfs(0, make([]bool, cardNumber))
	fmt.Fprintln(out, "Hand:", strings.Join(token[:cardNumber], " "), "Deck:", strings.Join(token[cardNumber:], " "), "Best hand:", bestHand)
}

func main() {
	in, _ := os.Open("131.in")
	defer in.Close()
	out, _ := os.Create("131.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		solve(out, s.Text())
	}
}
