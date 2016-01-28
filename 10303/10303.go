// UVa 10303 - How Many Trees?

// f[n] = C(2*n,n)/(n+1)
// f[n] = f[0]*f[n-1]+f[1]*f[n-2]+...+f[n-1]*f[0]
// f[n] = (4n-2)*f[n-1]/(n+1)

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 1001

var t [MAX]big.Int

func prepare() {
	t[1].SetInt64(1)
	var tmp big.Int
	for i := 2; i < MAX; i++ {
		tmp.SetInt64(int64(i + 1))
		t[i].SetInt64(int64(4*i - 2))
		t[i].Mul(&t[i], &t[i-1])
		t[i].Div(&t[i], &tmp)
	}
}

func main() {
	prepare()

	in, _ := os.Open("10303.in")
	defer in.Close()
	out, _ := os.Create("10303.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, &t[n])
	}
}
