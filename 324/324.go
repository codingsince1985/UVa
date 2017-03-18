// UVa 324 - Factorial Frequencies

package main

import (
	"fmt"
	"math/big"
	"os"
)

func factorial(n int) string {
	res := big.NewInt(1)
	for ; n > 1; n-- {
		res.Mul(res, big.NewInt(int64(n)))
	}
	return res.String()
}

func output(out *os.File, str string) {
	var res [10]int
	for i := range str {
		digit := str[i] - '0'
		res[digit]++
	}
	for i, v := range res {
		if fmt.Fprintf(out, "   (%d)%5d", i, v); i != 4 && i != 9 {
			fmt.Fprint(out, " ")
		} else {
			fmt.Fprintln(out)
		}
	}
}

func main() {
	in, _ := os.Open("324.in")
	defer in.Close()
	out, _ := os.Create("324.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "%d! --\n", n)
		output(out, factorial(n))
	}
}
