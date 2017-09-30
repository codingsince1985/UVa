// UVa 10268 - 498-bis

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(x int, a []int) int {
	var sum int
	n := len(a) - 1
	for i := n; i > 0; i-- {
		sum += a[n-i] * i * int(math.Pow(float64(x), float64(i-1)))
	}
	return sum
}

func main() {
	in, _ := os.Open("10268.in")
	defer in.Close()
	out, _ := os.Create("10268.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	var tmp int
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		x, _ := strconv.Atoi(line)
		var a []int
		s.Scan()
		for r := strings.NewReader(s.Text()); ; {
			if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
				break
			}
			a = append(a, tmp)
		}
		fmt.Fprintln(out, solve(x, a))
	}
}
