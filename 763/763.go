// UVa 763 - Fibinary Numbers

package main

import (
	"fmt"
	"os"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func add(n1, n2 string) []uint8 {
	l1, l2 := len(n1), len(n2)
	sum := make([]uint8, max(l1, l2)+2)
	var idx int
	for i := range n1 {
		sum[idx] = n1[l1-i-1] - '0'
		idx++
	}
	idx = 0
	for i := range n2 {
		sum[idx] += n1[l2-i-1] - '0'
		idx++
	}
	return sum
}

func reverse(sum []uint8) string {
	var reversed string
	for i := len(sum) - 1; i >= 0; i-- {
		reversed += string('0' + sum[i])
	}
	return reversed
}

func solve(n1, n2 string) string {
	sum := add(n1, n2)
	for ok := true; ok; {
		ok = true
		for i := 0; i < len(sum)-2; i++ {
			if sum[i] != 0 && sum[i+1] != 0 {
				ok = false
				sum[i]--
				sum[i+1]--
				sum[i+2]++
			}
			if sum[i] > 1 {
				ok = false
				sum[i] -= 2
				sum[i+1]++
				switch {
				case i == 1:
					sum[0]++
				case i > 1:
					sum[i-2]++
				}
			}
		}
	}
	return strings.TrimLeft(reverse(sum), "0")
}

func main() {
	in, _ := os.Open("763.in")
	defer in.Close()
	out, _ := os.Create("763.out")
	defer out.Close()

	var n1, n2 string
	first := true
	for {
		if first {
			first = false
		} else {
			fmt.Fscanln(in)
		}
		if _, err := fmt.Fscanf(in, "%s\n%s", &n1, &n2); err != nil {
			break
		}
		fmt.Fprintf(out, "%s\n\n", solve(n1, n2))
	}
}
