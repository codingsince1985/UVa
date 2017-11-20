// UVa 10579 - Fibonacci Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
)

var f = func() []big.Int {
	f := []big.Int{*big.NewInt(0), *big.NewInt(1)}
	for i := 2; ; i++ {
		var tmp big.Int
		if len(tmp.Add(&f[i-2], &f[i-1]).String()) > 1000 {
			break
		}
		f = append(f, tmp)
	}
	return f
}()

func main() {
	in, _ := os.Open("10579.in")
	defer in.Close()
	out, _ := os.Create("10579.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, &f[n])
	}
}
