// UVa 594 - One Little, Two Little, Three Little Endians

package main

import (
	"fmt"
	"math"
	"os"
)

func convert(n int64) int64 {
	if n < 0 {
		n = math.MaxUint32 + n + 1
	}

	bytes := make([]int, 4)
	for i := range bytes {
		bytes[i] = int(n & 255)
		n /= 256
	}

	n, base := 0, 1
	for i := 3; i >= 0; i-- {
		n = n + int64(base*bytes[i])
		base *= 256
	}

	if n > math.MaxInt32 {
		n = n - math.MaxUint32 - 1
	}
	return n
}

func main() {
	in, _ := os.Open("594.in")
	defer in.Close()
	out, _ := os.Create("594.out")
	defer out.Close()

	var n int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%d converts to %d\n", n, convert(n))
	}
}
