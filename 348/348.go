// UVa 348 - Optimal Array Multiplication Sequence

package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var (
	m [][2]int
	c [][][2]int
)

func makeCache(n int) {
	c = make([][][2]int, n)
	for i := range c {
		c[i] = make([][2]int, n)
		for j := range c[i] {
			c[i][j][0] = math.MaxInt32
		}
	}
}

//       |min(c[i,k]+c[k+1,j]+m[i].rows*m[k].cols*m[j].cols, k∈[i,j)), if j>i
//c[i,j]=|
//       |0, if j=i

func dp(i, j int) int {
	if c[i][j][0] != math.MaxInt32 {
		return c[i][j][0]
	}

	if i == j {
		c[i][j][0] = 0
		return 0
	}

	sum := math.MaxInt32
	pos := 0
	for k := i; k < j; k++ {
		left := dp(i, k)
		right := dp(k+1, j)
		cost := m[i][0] * m[k][1] * m[j][1]
		if sum > left+right+cost {
			sum = left + right + cost
			pos = k
		}
	}
	c[i][j][0], c[i][j][1] = sum, pos
	return c[i][j][0]
}

func output(out io.Writer, i, j int) {
	if i == j {
		fmt.Fprintf(out, "A%d", i+1)
	} else {
		fmt.Fprint(out, "(")
		output(out, i, c[i][j][1])
		fmt.Fprint(out, " x ")
		output(out, c[i][j][1]+1, j)
		fmt.Fprint(out, ")")
	}
}

func main() {
	in, _ := os.Open("348.in")
	defer in.Close()
	out, _ := os.Create("348.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		m = make([][2]int, n)
		for i := range m {
			fmt.Fscanf(in, "%d%d", &m[i][0], &m[i][1])
		}
		makeCache(n)
		dp(0, n-1)

		fmt.Fprintf(out, "Case %d: ", kase)
		output(out, 0, n-1)
		fmt.Fprintln(out)
	}
}
