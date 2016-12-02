// UVa 263 - Number Chains

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type digits []byte

func (a digits) Len() int { return len(a) }

func (a digits) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a digits) Less(i, j int) bool { return a[i] < a[j] }

func sortInt(num string, descending bool) int {
	d := digits(num)
	if descending {
		sort.Sort(sort.Reverse(d))
	} else {
		sort.Sort(d)
	}
	n, _ := strconv.Atoi(string(d))
	return n
}

func solve(out *os.File, num string) {
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
