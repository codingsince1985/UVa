// UVa 11827 - Maximum GCD

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func solve(num []int) int {
	var max int
	for i, l := 0, len(num); i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if g := gcd(num[i], num[j]); g > max {
				max = g
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("11827.in")
	defer in.Close()
	out, _ := os.Create("11827.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var tmp int
	s.Scan()
	for n, _ := strconv.Atoi(s.Text()); n > 0 && s.Scan(); n-- {
		var num []int
		for r := strings.NewReader(s.Text()); ; {
			if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
				break
			}
			num = append(num, tmp)
		}
		fmt.Fprintln(out, solve(num))
	}
}
