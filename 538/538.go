// UVa 538 - Balancing Bank Accounts

package main

import (
	"fmt"
	"io"
	"os"
)

type traveller struct {
	name    string
	balance int
}

func solve(out io.Writer, travellers []traveller) {
	for {
		var max, min, idx1, idx2 int
		for i, t := range travellers {
			if t.balance > max {
				max, idx1 = t.balance, i
			}
			if t.balance < min {
				min, idx2 = t.balance, i
			}
		}
		if max == 0 {
			break
		}
		fmt.Fprintf(out, "%s %s ", travellers[idx1].name, travellers[idx2].name)
		if max > -min {
			fmt.Fprintln(out, -min)
			travellers[idx1].balance += min
			travellers[idx2].balance = 0
		} else {
			fmt.Fprintln(out, max)
			travellers[idx1].balance = 0
			travellers[idx2].balance += max
		}
	}
}

func main() {
	in, _ := os.Open("538.in")
	defer in.Close()
	out, _ := os.Create("538.out")
	defer out.Close()

	var n, t, amount int
	var t1, t2 string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &t); n == 0 && t == 0 {
			break
		}
		nameMap := make(map[string]int)
		travellers := make([]traveller, n)
		for i := range travellers {
			fmt.Fscanf(in, "%s", &travellers[i].name)
			nameMap[travellers[i].name] = i
		}
		for ; t > 0; t-- {
			fmt.Fscanf(in, "%s%s%d", &t1, &t2, &amount)
			travellers[nameMap[t1]].balance -= amount
			travellers[nameMap[t2]].balance += amount
		}
		fmt.Fprintf(out, "Case #%d\n", kase)
		solve(out, travellers)
		fmt.Fprintln(out)
	}
}
