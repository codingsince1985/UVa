// UVa 10494 - If We Were a Child Again

package main

import (
	"fmt"
	"math/big"
	"os"
)

func calc(n1, n2, op string) string {
	var bn1, bn2, result big.Int
	bn1.SetString(n1, 10)
	bn2.SetString(n2, 10)
	switch op {
	case "/":
		result.Div(&bn1, &bn2)
	case "%":
		result.Mod(&bn1, &bn2)
	}
	return result.String()
}

func main() {
	in, _ := os.Open("10494.in")
	defer in.Close()
	out, _ := os.Create("10494.out")
	defer out.Close()

	var n1, n2, op string
	for {
		if _, err := fmt.Fscanf(in, "%s%s%s", &n1, &op, &n2); err != nil {
			break
		}
		fmt.Fprintln(out, calc(n1, n2, op))
	}
}
