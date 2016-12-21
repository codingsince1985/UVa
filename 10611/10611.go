// UVa 10611 - The Playboy Chimp

package main

import (
	"fmt"
	"os"
)

func binarySearchHigh(ladies []int, height int) int {
	l, r := 0, len(ladies)-1
	for len(ladies) > 2 {
		mid := (l + r) / 2
		if ladies[mid] <= height {
			return binarySearchHigh(ladies[mid+1:], height)
		} else {
			return binarySearchHigh(ladies[:mid+1], height)
		}
	}
	if ladies[l] > height {
		return ladies[l]
	} else if ladies[r] > height {
		return ladies[r]
	}
	return -1
}

func binarySearchLow(ladies []int, height int) int {
	l, r := 0, len(ladies)-1
	for len(ladies) > 2 {
		mid := (l + r) / 2
		if ladies[mid] >= height {
			return binarySearchLow(ladies[:mid], height)
		} else {
			return binarySearchLow(ladies[mid:], height)
		}
	}
	if ladies[r] < height {
		return ladies[r]
	} else if ladies[l] < height {
		return ladies[l]
	}
	return -1
}

func main() {
	in, _ := os.Open("10611.in")
	defer in.Close()
	out, _ := os.Create("10611.out")
	defer out.Close()

	var n, q, height int
	fmt.Fscanf(in, "%d", &n)
	ladies := make([]int, n)
	for i := range ladies {
		fmt.Fscanf(in, "%d", &ladies[i])
	}
	for fmt.Fscanf(in, "%d", &q); q > 0; q-- {
		fmt.Fscanf(in, "%d", &height)
		low, high := binarySearchLow(ladies, height), binarySearchHigh(ladies, height)
		if low != -1 {
			fmt.Fprint(out, low)
		} else {
			fmt.Fprint(out, "X")
		}
		if high != -1 {
			fmt.Fprintln(out, "", high)
		} else {
			fmt.Fprintln(out, " X")
		}
	}
}
