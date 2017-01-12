// UVa 10220 - I Love Big Numbers !

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 1001

var f = func() [MAX]big.Int {
	var f [MAX]big.Int
	f[0].SetInt64(1)
	for i := 1; i < MAX; i++ {
		tmp := big.NewInt(int64(i))
		f[i].Mul(&f[i-1], tmp)
	}
	return f
}()

func sum(n big.Int) int {
	var t int
	s := fmt.Sprint(&n)
	for _, v := range s {
		t += int(v - '0')
	}
	return t
}

func main() {
	in, _ := os.Open("10220.in")
	defer in.Close()
	out, _ := os.Create("10220.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, sum(f[n]))
	}
}
