// UVa 10334 - Ray Through Glasses

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 1001
var f [MAX]big.Int

func prepare() {
	f[0].SetInt64(1)
	f[1].SetInt64(2)
	for i := 2; i < MAX; i ++ {
		f[i].Add(&f[i - 2], &f[i - 1])
	}
}

func main() {
	prepare()

	in, _ := os.Open("10334.in")
	out, _ := os.Create("10334.out")
	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break;
		}
		fmt.Fprintln(out, &f[n])
	}
	in.Close()
	out.Close()
}