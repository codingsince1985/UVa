// UVa 498 - Polly the Polynomial

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(coefficients []int, x int) int {
	sum, n, tmp := 0, len(coefficients), 1
	for i := 0; i < n; i++ {
		sum += coefficients[len(coefficients)-i-1] * tmp
		tmp *= x
	}
	return sum
}

func main() {
	in, _ := os.Open("498.in")
	defer in.Close()
	out, _ := os.Create("498.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		cs := strings.Fields(s.Text())
		coefficients := make([]int, len(cs))
		for i, v := range cs {
			coefficients[i], _ = strconv.Atoi(v)
		}
		s.Scan()
		xs := strings.Fields(s.Text())
		for i, v := range xs {
			x, _ := strconv.Atoi(v)
			fmt.Fprint(out, calc(coefficients, x))
			if i != len(xs)-1 {
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
	}
}
