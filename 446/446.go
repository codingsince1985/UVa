// UVa 446 - Kibbles "n" Bits "n" Bits "n" Bits

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("446.in")
	defer in.Close()
	out, _ := os.Create("446.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	var n1, n2, sign string
	var result int64
	for n > 0 {
		fmt.Fscanf(in, "%s%s%s", &n1, &sign, &n2)
		num1, _ := strconv.ParseInt(n1, 16, 0)
		num2, _ := strconv.ParseInt(n2, 16, 0)
		switch sign {
		case "+":
			result = num1 + num2
		default:
			result = num1 - num2
		}
		fmt.Fprintf(out, "%013s %s %013s = %d\n", strconv.FormatInt(num1, 2), sign, strconv.FormatInt(num2, 2), result)
		n--
	}
}
