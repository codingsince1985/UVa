// UVa 507 - Jill Rides Again

package main

import (
	"fmt"
	"os"
)

func maxsub(s []int) (int, int) {
	start := -1
	longestStart, longestEnd := -1, -1
	var meh, msf int
	for i, v := range s {
		if meh += v; meh < 0 {
			meh = 0
			start = -1
		} else {
			if start == -1 {
				start = i
			}
			if msf <= meh {
				msf = meh
				if end := i + 1; end-start > longestEnd-longestStart {
					longestStart = start
					longestEnd = end
				}
			}
		}
	}
	return longestStart, longestEnd
}

func main() {
	in, _ := os.Open("507.in")
	defer in.Close()
	out, _ := os.Create("507.out")
	defer out.Close()

	var b, r int
	fmt.Fscanf(in, "%d", &b)
	for i := 1; i <= b; i++ {
		fmt.Fscanf(in, "%d", &r)
		s := make([]int, r-1)
		for j := range s {
			fmt.Fscanf(in, "%d", &s[j])
		}
		if start, end := maxsub(s); start == -1 {
			fmt.Fprintf(out, "Route %d has no nice parts\n", i)
		} else {
			fmt.Fprintf(out, "The nicest part of route %d is between stops %d and %d\n", i, start+1, end+1)
		}
	}
}
