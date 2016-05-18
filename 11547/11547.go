// UVa 11547 - Automatic Answer

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11547.in")
	defer in.Close()
	out, _ := os.Create("11547.out")
	defer out.Close()

	var t, n int
	fmt.Fscanf(in, "%d", &t)
	for t > 0 {
		fmt.Fscanf(in, "%d", &n)
		n = (n*63+7492)*235/47 - 498
		if n < 0 {
			n *= -1
		}
		n = n / 10 % 10
		fmt.Fprintln(out, n)
		t--
	}
}
