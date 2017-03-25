// UVa 10183 - How Many Fibs?

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 101

var p = func() []big.Int {
	p := []big.Int{*big.NewInt(1), *big.NewInt(1)}
	var s string
	for i := 2; ; i++ {
		var tmp big.Int
		tmp.Add(&p[i-2], &p[i-1])
		s = fmt.Sprintf("%v", tmp)
		if len(s) > max {
			break
		}
		p = append(p, tmp)
	}
	return p
}()

func main() {
	in, _ := os.Open("10183.in")
	defer in.Close()
	out, _ := os.Create("10183.out")
	defer out.Close()

	var a, b string
	var n1, n2 big.Int
	for {
		if fmt.Fscanf(in, "%s%s", &a, &b); a == "0" && b == "0" {
			break
		}
		n1.SetString(a, 10)
		n2.SetString(b, 10)

		c := 0
		for i := 1; i < len(p); i++ {
			if p[i].Cmp(&n1) >= 0 && p[i].Cmp(&n2) <= 0 {
				c++
			}
			if p[i].Cmp(&n2) > 0 {
				break
			}
		}
		fmt.Fprintln(out, c)
	}
}
