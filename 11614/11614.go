// UVa 11614 - Etruscan Warriors Never Play Chess

package main

import (
	"fmt"
	"math"
	"os"
)

func binarySearch(n int64) int64 {
	var low, high int64 = 0, math.MaxUint32
	for low+1 < high {
		mid := (low + high) / 2
		if curr := (mid + 1) * mid / 2; curr <= n {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	in, _ := os.Open("11614.in")
	defer in.Close()
	out, _ := os.Create("11614.out")
	defer out.Close()

	var kase int
	var n int64
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, binarySearch(n))
	}
}
