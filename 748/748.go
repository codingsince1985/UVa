// UVa 748 - Exponentiation

package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func pow(r string, n int) string {
	pos := len(r) - 1 - strings.Index(r, ".")
	r1 := strings.TrimLeft(strings.Replace(r, ".", "", 1), "0")
	r2 := strings.TrimRight(r1, "0")
	pos -= len(r1) - len(r2)

	var base big.Int
	base.SetString(r2, 10)
	result := big.NewInt(1)
	for i := 0; i < n; i++ {
		result.Mul(result, &base)
	}
	r3 := result.String()

	if pos == 0 {
		return r3
	}
	pos3 := pos * n
	if len(r3) < pos3 {
		r3 = strings.Repeat("0", pos3-len(r3)) + r3
	}
	return r3[:len(r3)-pos3] + "." + r3[len(r3)-pos3:]
}

func main() {
	in, _ := os.Open("748.in")
	defer in.Close()
	out, _ := os.Create("748.out")
	defer out.Close()

	var r string
	var n int
	for {
		if _, err := fmt.Fscanf(in, "%s%d", &r, &n); err != nil {
			break
		}
		fmt.Fprintln(out, pow(r, n))
	}
}
