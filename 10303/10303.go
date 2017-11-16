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

const max = 1001

var t = func() [max]big.Int {
	var t [max]big.Int
	t[1].SetInt64(1)
	for i := int64(2); i < max; i++ {
		t[i].SetInt64(4*i-2).Mul(&t[i], &t[i-1]).Div(&t[i], big.NewInt(i+1))
	}
	return t
}()

func main() {
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
