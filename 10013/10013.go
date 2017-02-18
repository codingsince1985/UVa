// UVa 10013 - Super long sums

package main

import (
	"fmt"
	"os"
)

func solve(n1, n2 []byte, n int) []byte {
	result := make([]byte, n)
	var carry bool
	for i := n - 1; i >= 0; i-- {
		sum := n1[i] + n2[i]
		if carry {
			sum++
		}
		carry = sum >= 10
		result[i] = sum % 10
	}
	return result
}

func main() {
	in, _ := os.Open("10013.in")
	defer in.Close()
	out, _ := os.Create("10013.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		n1 := make([]byte, n)
		n2 := make([]byte, n)
		for i := range n1 {
			fmt.Fscanf(in, "%d%d", &n1[i], &n2[i])
		}
		for _, v := range solve(n1, n2, n) {
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
