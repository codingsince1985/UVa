// UVa 10220 - I Love Big Numbers !

package main

import (
	"fmt"
	"os"
	"math/big"
)

const MAX = 1001

var f [MAX]big.Int

func prepare() {
	f[0].SetInt64(1)
	var tmp big.Int
	for i := 1; i < MAX; i ++ {
		tmp.SetInt64(int64(i))
		f[i].Mul(&f[i - 1], &tmp)
	}
}

func sum(n big.Int) int {
	var t int
	s := fmt.Sprint(&n)
	for i, _ := range s {
		t += int(s[i] - '0')
	}
	return t
}

func main() {
	prepare()
	in, _ := os.Open("10220.in")
	out, _ := os.Create("10220.out")
	var n int
	for {
		_, err := fmt.Fscanf(in, "%d", &n)
		if err != nil {
			break
		}
		fmt.Fprintln(out, sum(f[n]))
	}
	in.Close()
	out.Close()
}