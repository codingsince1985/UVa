// UVa 202 - Repeating Decimals

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func div(n, d int) (string, int) {
	var count, idx int
	var ok bool
	w := n / d
	var ans strings.Builder
	ans.WriteString(strconv.Itoa(w) + ".")
	r := n % d
	m := map[int]int{r: count}
	for {
		count++
		if r *= 10; r < d {
			ans.WriteString("0")
			if idx, ok = m[r]; ok {
				break
			}
		} else {
			ans.WriteString(strconv.Itoa(r / d))
			r = r % d
			if idx, ok = m[r]; ok {
				break
			}
			m[r] = count
		}
	}
	return ans.String(), count - idx
}

func output(out io.Writer, n, d int, ans string, rep int) {
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
