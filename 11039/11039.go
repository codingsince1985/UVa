// UVa 11039 - Building designing

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(red, blue []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(red)))
	sort.Sort(sort.Reverse(sort.IntSlice(blue)))
	var curr int
	var isRed bool
	if red[0] > blue[0] {
		curr = red[0]
		isRed = true
	} else {
		curr = blue[0]
	}
	count := 1
	for ; ; count, isRed = count+1, !isRed {
		if isRed {
			blueIdx := sort.Search(len(blue), func(i int) bool { return blue[i] < curr })
			if blueIdx == len(blue) {
				break
			}
			curr = blue[blueIdx]
		} else {
			redIdx := sort.Search(len(red), func(i int) bool { return red[i] < curr })
			if redIdx == len(red) {
				break
			}
			curr = red[redIdx]
		}
	}
	return count
}

func main() {
	in, _ := os.Open("11039.in")
	defer in.Close()
	out, _ := os.Create("11039.out")
	defer out.Close()

	var p, n, f int
	for fmt.Fscanf(in, "%d", &p); p > 0; p-- {
		var red, blue []int
		for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
			if fmt.Fscanf(in, "%d", &f); f > 0 {
				blue = append(blue, f)
			} else {
				red = append(red, -f)
			}
		}
		fmt.Fprintln(out, solve(red, blue))
	}
}
