// UVa 485 - Pascal's Triangle of Death

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 60

func main() {
	out, _ := os.Create("485.out")
	defer out.Close()

	var p []big.Int
here:
	for {
		p = append(p, *big.NewInt(1))
		l := len(p)
		fmt.Fprint(out, &p[l-1])
		for i := l - 2; i > 0; i-- {
			fmt.Fprintf(out, " %v", p[i].Add(&p[i], &p[i-1]))
			if s := fmt.Sprint(&p[i]); len(s) > max {
				fmt.Fprintln(out)
				break here
			}
		}

		if l > 1 {
			fmt.Fprintf(out, " %v", &p[0])
		}
		fmt.Fprintln(out)
	}
}
