// UVa 471 - Magic Numbers

package main

import (
	"fmt"
	"os"
)

func valid(n int64) bool {
	digitMap := make(map[int64]bool)
	for n > 0 {
		if d := n % 10; !digitMap[d] {
			digitMap[d] = true
			n /= 10
			continue
		}
		return false
	}
	return true
}

func solve(out *os.File, n int64) {
	for i := int64(1); ; i++ {
		m := n * i
		if m > 9876543210 {
			break
		}
		if valid(i) && valid(m) {
			fmt.Fprintf(out, "%d / %d = %d\n", m, i, n)
		}
	}
}

func main() {
	in, _ := os.Open("471.in")
	defer in.Close()
	out, _ := os.Create("471.out")
	defer out.Close()

	var kase, n int64
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, n)
	}
}
