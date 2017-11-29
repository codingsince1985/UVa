// UVa 10427 - Naughty Sleepy Boys

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var d = func() []int {
	d := make([]int, 8)
	nines, tens := 9, 1
	for i := range d {
		d[i] = (nines - tens + 1) * (i + 1)
		nines = nines*10 + 9
		tens *= 10
	}
	return d
}()

func find(n, total, idx int) byte {
	from := 0
	if idx > 0 {
		from = int(math.Pow(10, float64(idx)))
		from--
	}
	from += (n - total) / (idx + 1)
	if (n-total)%(idx+1) > 0 {
		from++
	}
	return strconv.Itoa(from)[(n-total-1)%(idx+1)] - '0'
}

func solve(n int) byte {
	var total, idx int
	for idx = range d {
		if total+d[idx] >= n {
			break
		}
		total += d[idx]
	}
	return find(n, total, idx)
}

func main() {
	in, _ := os.Open("10427.in")
	defer in.Close()
	out, _ := os.Create("10427.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}
