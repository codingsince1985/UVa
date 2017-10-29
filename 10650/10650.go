// UVa 10650 - Determinate Prime

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const max = 32000

var primes = func() []int {
	var pl []int
	for i, p := range sieve() {
		if !p {
			pl = append(pl, i)
		}
	}
	return pl
}()

func sieve() []bool {
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func findIndex(x, y int) (int, int) {
	if x > y {
		x, y = y, x
	}
	idx1, idx2 := -1, -1
	for i, p := range primes {
		if p < x {
			continue
		}
		if idx1 == -1 {
			idx1 = i
		}
		if p <= y {
			idx2 = i
		}
	}
	return idx1, idx2
}

func solve(x, y int) [][]string {
	var lists [][]string
	idx1, idx2 := findIndex(x, y)
	for i := idx1; i <= idx2-2; i++ {
		if primes[i+2]-primes[i+1] == primes[i+1]-primes[i] {
			diff := primes[i+1] - primes[i]
			if i == idx1 && idx1 > 0 && primes[idx1]-primes[idx1-1] == diff {
				continue
			}
			list := []string{strconv.Itoa(primes[i]), strconv.Itoa(primes[i+1]), strconv.Itoa(primes[i+2])}
			idx := i + 3
			for ; idx <= idx2 && primes[idx]-primes[idx-1] == diff; idx++ {
				list = append(list, strconv.Itoa(primes[idx]))
			}
			if idx == idx2+1 && idx < len(primes) && primes[idx2+1]-primes[idx2] == diff {
				continue
			}
			lists = append(lists, list)
		}
	}
	return lists
}

func main() {
	in, _ := os.Open("10650.in")
	defer in.Close()
	out, _ := os.Create("10650.out")
	defer out.Close()

	var x, y int
	for {
		if fmt.Fscanf(in, "%d%d", &x, &y); x == 0 && y == 0 {
			break
		}
		for _, list := range solve(x, y) {
			fmt.Fprintln(out, strings.Join(list, " "))
		}
	}
}
