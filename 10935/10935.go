// UVa 10935 - Throwing cards away I

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(n int) ([]string, string) {
	var discarded []string
	card := make([]string, n)
	for i := range card {
		card[i] = strconv.Itoa(i + 1)
	}
	for len(card) > 1 {
		discarded = append(discarded, card[0])
		card = append(card[2:], card[1])
	}
	return discarded, card[0]
}

func main() {
	in, _ := os.Open("10935.in")
	defer in.Close()
	out, _ := os.Create("10935.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		discarded, remaining := solve(n)
		fmt.Fprintf(out, "Discarded cards: %s\n", strings.Join(discarded, ", "))
		fmt.Fprintf(out, "Remaining card: %s\n", remaining)
	}
}
