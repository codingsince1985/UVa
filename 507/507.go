// UVa 507 - Jill Rides Again

package main

import (
	"fmt"
	"os"
)

func maxsub(s []int) (int, int, int) {
	start, end := -1, -1
	longestStart, longestEnd := -1, -1
	var meh, msf int
	for i, v := range s {
		meh += v
		if meh < 0 {
			meh = 0
			start = -1
		} else {
			if start == -1 {
				start = i
			}
			if msf <= meh {
				end = i + 1
				msf = meh
				if end-start > longestEnd-longestStart {
					longestStart = start
					longestEnd = end
				}
			}
		}
	}
	return longestStart, longestEnd, msf
}

func main() {
	in, _ := os.Open("507.in")
	defer in.Close()
	out, _ := os.Create(("507.out"))
	defer out.Close()

	var b, r int
	var s []int
	fmt.Fscanf(in, "%d", &b)
	for i := 0; i < b; i++ {
		fmt.Fscanf(in, "%d", &r)
		s = make([]int, r-1)
		for j := range s {
			fmt.Fscanf(in, "%d", &s[j])
		}
		if start, end, _ := maxsub(s); start == -1 {
			fmt.Fprintf(out, "Route %d has no nice parts\n", i+1)
		} else {
			fmt.Fprintf(out, "The nicest part of route %d is between stops %d and %d\n", i+1, start+1, end+1)
		}
	}
}
