// UVa 10669 - Three powers

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

var (
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
)

func power(n int) string {
	var p big.Int
	for p.Set(one); n > 0; n-- {
		p.Mul(&p, three)
	}
	return fmt.Sprintf(" %v", &p)
}

func solve(line string) []string {
	var list []string
	var n, mod big.Int
	n.SetString(line, 10)
	n.Sub(&n, one)
	for i := 0; n.Cmp(zero) != 0; i++ {
		if n.DivMod(&n, two, &mod); mod.Cmp(one) == 0 {
			list = append(list, power(i))
		}
	}
	return list
}

func main() {
	in, _ := os.Open("10669.in")
	defer in.Close()
	out, _ := os.Create("10669.out")
	defer out.Close()

	var n string
	for {
		if fmt.Fscanf(in, "%s", &n); n == "0" {
			break
		}
		fmt.Fprintln(out, "{"+strings.Join(solve(n), ",")+" }")
	}
}
