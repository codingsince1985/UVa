// UVa 202 - Repeating Decimals

package main

import (
	"fmt"
	"os"
	"strconv"
)

func div(n, d int) (string, int) {
	var count, idx int
	var ok bool
	w := n / d
	ans := strconv.Itoa(w) + "."
	r := n % d
	m := make(map[int]int)
	m[r] = count
	for {
		count++
		if r *= 10; r < d {
			ans += "0"
			if idx, ok = m[r]; ok {
				break
			}
		} else {
			ans += strconv.Itoa(r / d)
			r = r % d
			if idx, ok = m[r]; ok {
				break
			}
			m[r] = count
		}
	}
	return ans, count - idx
}

func output(out *os.File, n, d int, ans string, rep int) {
	fmt.Fprintf(out, "%d/%d = ", n, d)
	fmt.Fprint(out, ans[:len(ans)-rep], "(", ans[len(ans)-rep:], ")")
	fmt.Fprintf(out, "\n   %d = number of digits in repeating cycle\n\n", rep)
}

func main() {
	in, _ := os.Open("202.in")
	defer in.Close()
	out, _ := os.Create("202.out")
	defer out.Close()

	var n, d int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &d); err != nil {
			break
		}
		ans, rep := div(n, d)
		output(out, n, d, ans, rep)
	}
}
