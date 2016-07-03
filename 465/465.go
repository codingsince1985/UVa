// UVa 465 - Overflow

package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	in, _ := os.Open("465.in")
	defer in.Close()
	out, _ := os.Create("465.out")
	defer out.Close()

	var s1, s2 string
	var op byte
	var i1, i2, i3 big.Int
	var maxInt = big.NewInt(math.MaxInt32)
	for {
		if _, err := fmt.Fscanf(in, "%s %c %s", &s1, &op, &s2); err != nil {
			break
		}
		fmt.Fprintf(out, "%s %c %s\n", s1, op, s2)
		i1.SetString(s1, 10)
		if i1.Cmp(maxInt) > 0 {
			fmt.Fprintln(out, "first number too big")
		}
		i2.SetString(s2, 10)
		if i2.Cmp(maxInt) > 0 {
			fmt.Fprintln(out, "second number too big")
		}
		if op == '+' {
			i3.Add(&i1, &i2)
		} else {
			i3.Mul(&i1, &i2)
		}
		if i3.Cmp(maxInt) > 0 {
			fmt.Fprintln(out, "result too big")
		}
	}
}
