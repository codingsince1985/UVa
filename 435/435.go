// UVa 435 - Block Voting

package main

import (
	"fmt"
	"os"
)

func solve(p, total int, votes []int) []int {
	powerIndex := make([]int, p)
	for i := 0; i < (1 << uint(p)); i++ {
		var sum int
		for j := range votes {
			if i&(1<<uint(j)) != 0 {
				sum += votes[j]
			}
		}
		if sum <= total/2 {
			for j := range votes {
				if i&(1<<uint(j)) == 0 && sum+votes[j] > total/2 {
					powerIndex[j]++
				}
			}
		}
	}
	return powerIndex
}

func main() {
	in, _ := os.Open("435.in")
	defer in.Close()
	out, _ := os.Create("435.out")
	defer out.Close()

	var kase, p int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &p)
		var total int
		votes := make([]int, p)
		for i := range votes {
			fmt.Fscanf(in, "%d", &votes[i])
			total += votes[i]
		}
		powerIndex := solve(p, total, votes)
		for i, power := range powerIndex {
			fmt.Fprintf(out, "party %d has power index %d\n", i+1, power)
		}
		fmt.Fprintln(out)
	}
}
