// UVa 119 - Greedy Gift Givers

package main

import (
	"fmt"
	"os"
)

type giver struct {
	name   string
	amount int
	toList []string
}

func main() {
	in, _ := os.Open("119.in")
	defer in.Close()
	out, _ := os.Create("119.out")
	defer out.Close()

	var n, num int
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		names := make([]string, n)
		for i := range names {
			fmt.Fscanf(in, "%s", &names[i])
		}

		givers := make([]giver, n)
		table := make(map[string]int)
		for i := range givers {
			fmt.Fscanf(in, "%s%d%d", &givers[i].name, &givers[i].amount, &num)
			if num != 0 {
				table[givers[i].name] = -(givers[i].amount / num) * num
			}
			givers[i].toList = make([]string, num)
			for j := range givers[i].toList {
				fmt.Fscanf(in, "%s", &givers[i].toList[j])
			}
		}

		for _, fm := range givers {
			for _, to := range fm.toList {
				table[to] += fm.amount / len(fm.toList)
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		for _, v := range names {
			fmt.Fprintln(out, v, table[v])
		}
	}
}
