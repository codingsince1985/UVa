// UVa 402 - M*A*S*H

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cards [20]int

func count(n, x int) []string {
	members := make([]int, n)
	for i := range members {
		members[i] = i + 1
	}
	idx, eliminated := 0, 0
here:
	for {
		memberCount := 0
		card := cards[idx]
		for i, member := range members {
			if member != 0 {
				memberCount++
				if memberCount%card == 0 {
					members[i] = 0
					eliminated++
					if eliminated == n-x {
						break here
					}
				}
			}
		}
		idx++
	}
	var remaining []string
	for _, member := range members {
		if member != 0 {
			remaining = append(remaining, strconv.Itoa(member))
		}
	}
	return remaining
}

func main() {
	in, _ := os.Open("402.in")
	defer in.Close()
	out, _ := os.Create("402.out")
	defer out.Close()

	var kase, n, x int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &x); err != nil {
			break
		}
		for i := range cards {
			fmt.Fscanf(in, "%d", &cards[i])
		}
		kase++
		fmt.Fprintf(out, "Selection #%d\n", kase)
		fmt.Fprintf(out, "%s\n\n", strings.Join(count(n, x), " "))
	}
}
