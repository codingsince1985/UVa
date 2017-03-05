// UVa 412 - Pi

package main

import (
	"fmt"
	"math"
	"os"
)

func noCommonFactor(a, b int) bool {
	var factors []int
	if a%2 == 0 && b%2 == 0 {
		factors = append(factors, 2)
	}
	for i := 3; i <= a && i <= b; i += 2 {
		if a%i == 0 && b%i == 0 {
			factors = append(factors, i)
		}
	}
	return len(factors) == 0
}

func calculatePi(out *os.File, nums []int) {
	size := len(nums)
	total := size * (size - 1) / 2
	cnt := 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if noCommonFactor(nums[i], nums[j]) {
				cnt++
			}
		}
	}
	if cnt == 0 {
		fmt.Fprintln(out, "No estimate for this data set.")
	} else {
		fmt.Fprintf(out, "%.6f\n", math.Sqrt(float64(total)*6/float64(cnt)))
	}
}

func main() {
	in, _ := os.Open("412.in")
	defer in.Close()
	out, _ := os.Create("412.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		nums := make([]int, n)
		for i := range nums {
			fmt.Fscanf(in, "%d", &nums[i])
		}
		calculatePi(out, nums)
	}
}
