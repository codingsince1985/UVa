// UVa 416 - LED Test

package main

import (
	"fmt"
	"os"
)

var (
	n        int
	segments []int
	digitMap = map[int]int{0: 126, 1: 48, 2: 109, 3: 121, 4: 51, 5: 91, 6: 95, 7: 112, 8: 127, 9: 123}
)

func toBinary(line string) int {
	var total int
	for i, ch := range line {
		if ch == 'Y' {
			total += 1 << uint(6-i)
		}
	}
	return total
}

func dfs(curr, level, mask int) bool {
	if level == n {
		return true
	}
	for i := 0; i <= mask; i++ {
		if i&mask == i && digitMap[curr]&i == segments[level] && dfs(curr-1, level+1, i) {
			return true
		}
	}
	return false
}

func solve() bool {
	for i := 9; i >= n-1; i-- {
		if dfs(i, 0, 127) {
			return true
		}
	}
	return false
}

func main() {
	in, _ := os.Open("416.in")
	defer in.Close()
	out, _ := os.Create("416.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		segments = make([]int, n)
		for i := range segments {
			fmt.Fscanf(in, "%s", &line)
			segments[i] = toBinary(line)
		}
		if solve() {
			fmt.Fprintln(out, "MATCH")
		} else {
			fmt.Fprintln(out, "MISMATCH")
		}
	}
}
