// UVa 10182 - Bee Maja

package main

import (
	"fmt"
	"os"
)

var (
	directions = [][2]int{{-1, 1}, {-1, 0}, {0, -1}, {1, -1}, {1, 0}, {0, 1}}
	rings      = func() []int {
		rings := []int{0}
		for i, step := 1, 6; i < 100000; i, step = i+step, step+6 {
			rings = append(rings, i)
		}
		return rings
	}()
)

func binarySearch(n int) int {
	l, r := 0, len(rings)-1
	for l < r {
		mid := (l + r) / 2
		switch {
		case n == rings[mid]:
			return mid
		case n > rings[mid]:
			l = mid + 1
		default:
			r = mid
		}
	}
	return l
}

func reMap(ring, offset int) (int, int) {
	if ring == 1 {
		return 0, 0
	}
	x, y := ring-2, 1
	if offset > 0 {
	here:
		for i, direction := range directions {
			side := ring - 1
			if i == 0 {
				side--
			}
			for ; side > 0; side-- {
				x += direction[0]
				y += direction[1]
				if offset--; offset == 0 {
					break here
				}
			}
		}
	}
	return x, y
}

func main() {
	in, _ := os.Open("10182.in")
	defer in.Close()
	out, _ := os.Create("10182.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		ring := binarySearch(n)
		x, y := reMap(ring, n-rings[ring-1]-1)
		fmt.Fprintln(out, x, y)
	}
}
