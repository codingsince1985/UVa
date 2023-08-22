// UVa 338 - Long Multiplication

package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strings"
)

func singleNonZero(n int64) bool {
	var count int
	for ; n > 0; n /= 10 {
		if n%10 != 0 {
			count++
		}
	}
	return count == 1
}

func solve(n1, n2 int64) ([]string, int) {
	s1, s2 := fmt.Sprint(n1), fmt.Sprint(n2)
	length := max(len(s1), len(s2))
	lines := []string{s1, s2, strings.Repeat("-", length)}
	var mid int
	var product big.Int
	for count, tmp := 0, n2; tmp != 0; tmp /= 10 {
		i := tmp % 10
		if n1 != 0 && !singleNonZero(n2) && n1*i != 0 {
			lines = append(lines, fmt.Sprint(n1*i)+strings.Repeat(" ", count))
			mid++
		}
		product.Add(&product, big.NewInt(n1*i*int64(math.Pow(10, float64(count)))))
		count++
	}
	result := product.Text(10)
	length = max(length, len(result))
	if mid >= 2 {
		lines = append(lines, strings.Repeat("-", length))
	}
	lines = append(lines, result)
	return lines, length
}

func main() {
	in, _ := os.Open("338.in")
	defer in.Close()
	out, _ := os.Create("338.out")
	defer out.Close()

	var n1, n2 int64
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n1, &n2); err != nil {
			break
		}
		lines, length := solve(n1, n2)
		for _, line := range lines {
			fmt.Fprintf(out, "%s\n", strings.Repeat(" ", length-len(line))+strings.TrimRight(line, " "))
		}
		fmt.Fprintln(out)
	}
}
