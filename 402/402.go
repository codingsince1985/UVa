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
	eliminated := 0
here:
	for idx := 0; ; idx++ {
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

	var n, x int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &x); err != nil {
			break
		}
		for i := range cards {
			fmt.Fscanf(in, "%d", &cards[i])
		}
		fmt.Fprintf(out, "Selection #%d\n", kase)
		fmt.Fprintf(out, "%s\n\n", strings.Join(count(n, x), " "))
	}
}
