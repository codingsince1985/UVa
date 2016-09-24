// UVa 10784 - Diagonal

package main

import (
	"fmt"
	"math"
	"os"
)

func binarySearch(n uint64) uint64 {
	var low, high uint64
	low, high = 0, math.MaxInt64
	for low+1 < high {
		mid := (low + high) / 2
		if mid*(mid-3)/2 >= n {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low*(low-3)/2 >= n {
		return low
	}
	return high
}

func main() {
	in, _ := os.Open("10784.in")
	defer in.Close()
	out, _ := os.Create("10784.out")
	defer out.Close()

	var n, kase uint64
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		kase++
		fmt.Fprintf(out, "Case %d: %d\n", kase, binarySearch(n))
	}
}
