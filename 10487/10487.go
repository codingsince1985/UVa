// UVa 10487 - Closest Sums

package main

import (
	"fmt"
	"math"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func closestSum(num []int, s int) int {
	var sum int
	diff := math.MaxInt32
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			newDiff := abs(s - num[i] - num[j])
			if newDiff == 0 {
				return num[i] + num[j]
			}
			if newDiff < diff {
				diff = newDiff
				sum = num[i] + num[j]
			}
		}
	}
	return sum
}

func main() {
	in, _ := os.Open("10487.in")
	defer in.Close()
	out, _ := os.Create("10487.out")
	defer out.Close()

	var n, m, s int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		fmt.Fprintf(out, "Case %d:\n", kase)
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%d", &s)
			fmt.Fprintf(out, "Closest sum to %d is %d.\n", s, closestSum(num, s))
		}
	}
}
