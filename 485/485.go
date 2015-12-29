// UVa 485 - Pascal's Triangle of Death

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 60

var p []big.Int
var s string
var done = false

func main() {
	out, _ := os.Create("485.out")
	defer out.Close()

	for !done {
		var one big.Int
		one.SetInt64(1)
		p = append(p, one)
		l := len(p)

		fmt.Fprint(out, &p[l - 1])
		for i := l - 2; i > 0; i -- {
			p[i].Add(&p[i], &p[i - 1])
			fmt.Fprintf(out, " %v", &p[i])

			s = fmt.Sprint(&p[i])
			if len(s) > MAX {
				done = true
				break
			}
		}

		if !done && l > 1 {
			fmt.Fprintf(out, " %v", &p[0])
		}
		fmt.Fprintln(out)
	}
}