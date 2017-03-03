// UVa 10229 - Modular Fibonacci

package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	I, O, A [2][2]big.Int
	two     = big.NewInt(2)
	out     *os.File
)

func multiply(a, b [2][2]big.Int, m int64) [2][2]big.Int {
	pow := big.NewInt(m)
	var mod big.Int
	mod.Exp(two, pow, nil)
	var tmp [2][2]big.Int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				var t big.Int
				t.Mul(&a[i][k], &b[k][j])
				tmp[i][j].Add(&tmp[i][j], &t)
				tmp[i][j].Mod(&tmp[i][j], &mod)
			}
		}
	}
	return tmp
}

func calc(n, m int64) {
	x, y := I, A
	for n != 0 {
		if n&1 == 1 {
			x = multiply(x, y, m)
		}
		y = multiply(y, y, m)
		n /= 2
	}
	fmt.Fprintln(out, &x[1][0])
}

func setMatrix(i, j, k, l int64, m *[2][2]big.Int) {
	(*m)[0][0].SetInt64(i)
	(*m)[0][1].SetInt64(j)
	(*m)[1][0].SetInt64(k)
	(*m)[1][1].SetInt64(l)
}

func initialize() {
	setMatrix(1, 0, 0, 1, &I)
	setMatrix(0, 0, 0, 0, &O)
	setMatrix(1, 1, 1, 0, &A)
}

func main() {
	in, _ := os.Open("10229.in")
	defer in.Close()
	out, _ = os.Create("10229.out")
	defer out.Close()

	initialize()
	var n, m int64
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		calc(n, m)
	}
}
