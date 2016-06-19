// UVa 10193 - All You Need Is Love

package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	in, _ := os.Open("10193.in")
	defer in.Close()
	out, _ := os.Create("10193.out")
	defer out.Close()

	var n int
	var num string
	fmt.Fscanf(in, "%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%s", &num)
		num1, _ := strconv.ParseInt(num, 2, 64)
		fmt.Fscanf(in, "%s", &num)
		num2, _ := strconv.ParseInt(num, 2, 64)
		if gcd(num1, num2) == 1 {
			fmt.Fprintf(out, "Pair #%d: Love is not all you need!\n", i)
		} else {
			fmt.Fprintf(out, "Pair #%d: All you need is love!\n", i)
		}
	}
}
