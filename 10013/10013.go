// UVa 10013 - Super long sums

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10013.in")
	defer in.Close()
	out, _ := os.Create("10013.out")
	defer out.Close()

	var kase, n int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "\n%d", &n)
		n1 := make([]byte, n)
		n2 := make([]byte, n)
		for i := range n1 {
			fmt.Fscanf(in, "%d%d", &n1[i], &n2[i])
		}
		res := make([]byte, n)
		var carry bool
		for i := n - 1; i >= 0; i-- {
			sum := n1[i] + n2[i]
			if carry {
				sum++
			}
			carry = sum >= 10
			res[i] = sum % 10
		}
		for _, v := range res {
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
		kase--
	}
}
