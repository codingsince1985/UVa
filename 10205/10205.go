// UVa 10205 - Stack 'em Up

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const total = 52

var (
	suits  = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
)

func do(shuffles [][]int, order []int) []int {
	cards := make([]int, total)
	for i := range cards {
		cards[i] = i
	}
	for _, vi := range order {
		newCards := make([]int, total)
		for j, vj := range shuffles[vi] {
			newCards[j] = cards[vj-1]
		}
		copy(cards, newCards)
	}
	return cards
}

func main() {
	in, _ := os.Open("10205.in")
	defer in.Close()
	out, _ := os.Create("10205.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, tmp, n int
	var line string
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &kase)
	for s.Scan(); kase > 0 && s.Scan(); kase-- {
		var shuffles [][]int
		for fmt.Sscanf(s.Text(), "%d", &n); n > 0; n-- {
			var shuffle []int
			for s.Scan() {
				for r := strings.NewReader(s.Text()); ; {
					if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
						break
					}
					shuffle = append(shuffle, tmp)
				}
				if len(shuffle) == total {
					break
				}
			}
			shuffles = append(shuffles, shuffle)
		}
		var order []int
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%d", &tmp)
			order = append(order, tmp-1)
		}
		for _, vi := range do(shuffles, order) {
			fmt.Fprintf(out, "%s of %s\n", values[vi%13], suits[vi/13])
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
