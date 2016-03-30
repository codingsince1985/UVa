// UVa 713 - Adding Reversed Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
)

func reverse(s string) string {
	size := len(s)
	var ret string
	for i := range s {
		ret += string(s[size-1-i])
	}
	return ret
}

func main() {
	in, _ := os.Open("713.in")
	defer in.Close()
	out, _ := os.Create("713.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	var s1, s2 string
	var n1, n2 big.Int
	for n > 0 {
		fmt.Fscanf(in, "%s%s", &s1, &s2)
		n1.SetString(reverse(s1), 10)
		n2.SetString(reverse(s2), 10)
		n1.Add(&n1, &n2)
		n1.SetString(reverse(n1.String()), 10)
		fmt.Fprintln(out, &n1)
		n--
	}
}
