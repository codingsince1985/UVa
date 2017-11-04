// UVa 756 - Biorhythms

package main

import (
	"fmt"
	"os"
)

const (
	cycle  = 23 * 28 * 33
	pCycle = 28 * 33 * 6
	eCycle = 23 * 33 * 19
	iCycle = 23 * 28 * 2
)

func solve(p, e, i, d int) int { return (p*pCycle+e*eCycle+i*iCycle-d+(cycle-1))%cycle + 1 }

func main() {
	in, _ := os.Open("756.in")
	defer in.Close()
	out, _ := os.Create("756.out")
	defer out.Close()

	var p, e, i, d int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d%d%d", &p, &e, &i, &d); p == -1 && e == -1 && i == -1 && d == -1 {
			break
		}
		fmt.Fprintf(out, "Case %d: the next triple peak occurs in %d days.\n", kase, solve(p, e, i, d))
	}
}
