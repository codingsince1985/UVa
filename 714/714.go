// UVa 714 - Copying Books

package main

import (
	"fmt"
	"os"
)

var (
	out       *os.File
	sum, m, k int64
)

func binarySearch(high int64, books []int64) int64 {
	var low, mid, sum, groups, i int64
	for low < high {
		sum, groups = 0, 1
		mid = (low + high) / 2
		for i = 0; i < m; i++ {
			if sum+books[i] > mid {
				sum = books[i]
				groups++
			} else {
				sum += books[i]
			}
		}
		if groups <= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func solve(books []int64) [][]int64 {
	max := binarySearch(sum, books)
	var subTotal int64
	scribers := make([][]int64, k)
	for i, j := m-1, k-1; i >= 0; i-- {
		if subTotal+books[i] > max || j > i {
			j--
			subTotal = 0
		}
		subTotal += books[i]
		scribers[j] = append([]int64{books[i]}, scribers[j]...)
	}
	return scribers
}

func output(scribers [][]int64) {
	for i, scriber := range scribers {
		s := fmt.Sprint(scriber)
		fmt.Fprint(out, s[1:len(s)-1])
		if i < int(k-1) {
			fmt.Fprint(out, " / ")
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("714.in")
	defer in.Close()
	out, _ = os.Create("714.out")
	defer out.Close()

	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &m, &k)
		sum = 0
		books := make([]int64, m)
		for i := range books {
			fmt.Fscanf(in, "%d", &books[i])
			sum += books[i]
		}
		output(solve(books))
	}
}
