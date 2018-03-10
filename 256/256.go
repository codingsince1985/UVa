// UVa 256 - Quirksome Squares

package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func squaresOf(n int) []int {
	var ret []int
	for i := 0; i < n; i++ {
		ret = append(ret, i*i)
	}
	return ret
}

func split(num, limit int) [2]int { return [2]int{num / limit, num % limit} }

func quirksome(n int, out io.Writer) {
	limit := int(math.Pow10(n / 2))
	squares := squaresOf(limit)
	for _, v := range squares {
		rt := int(math.Sqrt(float64(v)))
		pairs := split(v, limit)
		if pairs[0]+pairs[1] == rt {
			format := fmt.Sprintf("%%0%dd\n", n)
			fmt.Fprintf(out, format, v)
		}
	}
}

func main() {
	in, _ := os.Open("256.in")
	defer in.Close()
	out, _ := os.Create("256.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		quirksome(n, out)
	}
}
