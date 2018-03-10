// UVa 263 - Number Chains

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func sortInt(num string, descending bool) int {
	d := []byte(num)
	if descending {
		sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
	} else {
		sort.Slice(d, func(i, j int) bool { return d[i] < d[j] })
	}
	var n int
	fmt.Sscanf(string(d), "%d", &n)
	return n
}

func solve(out io.Writer, num string) {
	fmt.Fprintf(out, "Original number was %s\n", num)
	chains := make(map[int]bool)
	for {
		d, a := sortInt(num, true), sortInt(num, false)
		diff := d - a
		fmt.Fprintf(out, "%d - %d = %d\n", d, a, diff)
		if chains[diff] {
			break
		}
		chains[diff] = true
		num = strconv.Itoa(diff)
	}
	fmt.Fprintf(out, "Chain length %d\n\n", len(chains)+1)
}

func main() {
	in, _ := os.Open("263.in")
	defer in.Close()
	out, _ := os.Create("263.out")
	defer out.Close()

	var num string
	for {
		if fmt.Fscanf(in, "%s", &num); num == "0" {
			break
		}
		solve(out, num)
	}
}
