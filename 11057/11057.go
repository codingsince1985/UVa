// UVa 11057 - Exact Sum

package main

import (
	"fmt"
	"os"
	"sort"
)

func binarySearch(price int, books []int) bool {
	l, r := 0, len(books)-1
	for l+1 < r {
		mid := (l + r) / 2
		if books[mid] == price {
			return true
		}
		if books[mid] > price {
			r = mid
		} else {
			l = mid
		}
	}
	return books[l] == price || books[r] == price
}

func solve(m int, books []int) [2]int {
	var pair [2]int
	sort.Ints(books)
	for i, book := range books {
		if book > m/2 || i >= len(books)/2 {
			break
		}
		if binarySearch(m-book, books[i+1:]) {
			pair[0], pair[1] = book, m-book
		}
	}
	return pair
}

func main() {
	in, _ := os.Open("11057.in")
	defer in.Close()
	out, _ := os.Create("11057.out")
	defer out.Close()

	var n, m int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		books := make([]int, n)
		for i := range books {
			fmt.Fscanf(in, "%d", &books[i])
		}
		fmt.Fscanf(in, "%d", &m)
		fmt.Fscanln(in)
		pair := solve(m, books)
		fmt.Fprintf(out, "Peter should buy books whose prices are %d and %d.\n\n", pair[0], pair[1])
	}
}
