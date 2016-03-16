// UVa 10127 - Ones

package main

import (
	"fmt"
	"os"
)

func divide(n int) int {
	num, cnt := 1, 1
	for num != 0 {
		if num < n {
			num = num*10 + 1
			cnt++
		} else {
			num %= n
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("10127.in")
	defer in.Close()
	out, _ := os.Create("10127.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, divide(n))
	}
}
