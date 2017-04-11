// UVa 10706 - Number Sequence

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const max = 40000

var (
	out        *os.File
	cache, str = func() ([]int, []byte) {
		cache := make([]int, max)
		cache[0] = 0
		str := []byte{'0'}
		for idx, i := 0, 1; i < max; i++ {
			idx += int(math.Log10(float64(i))) + 1
			cache[i] = idx + cache[i-1]
			str = append(str, []byte(strconv.Itoa(i))...)
		}
		return cache, str
	}()
)

func binarySearch(n int) {
	low, high := 1, max
	for low+1 < high {
		mid := (low + high) / 2
		if cache[mid] == n {
			low = mid
			break
		}
		if cache[mid] < n {
			low = mid
		} else {
			high = mid
		}
	}
	if n == cache[low] {
		fmt.Fprintln(out, low)
	} else {
		fmt.Fprintf(out, "%c\n", str[n-cache[low]])
	}
}

func main() {
	in, _ := os.Open("10706.in")
	defer in.Close()
	out, _ = os.Create("10706.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		binarySearch(n)
	}
}
