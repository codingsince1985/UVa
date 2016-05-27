// UVa 11185 - Ternary

package main

import (
	"fmt"
	"os"
	"strconv"
)

func ternary(n int) string {
	if n == 0 {
		return "0"
	}
	var str string
	for n > 0 {
		str = strconv.Itoa(n%3) + str
		n /= 3
	}
	return str
}

func main() {
	in, _ := os.Open("11185.in")
	defer in.Close()
	out, _ := os.Create("11185.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		fmt.Fprintln(out, ternary(n))
	}
}
