// UVa 160 - Factors and Factorials

package main

import (
	"fmt"
	"os"
)

var cache map[int]map[int]int

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int) []int {
	var ps []int
	ps = append(ps, 2)
	for i := 3; i <= n; i += 2 {
		if isPrime(i) {
			ps = append(ps, i)
		}
	}
	return ps
}

func factors(n int, ps []int) map[int]int {
	if fs, ok := cache[n]; ok {
		return fs
	}
	fs := make(map[int]int)
	for i := n; i != 1; {
		for _, p := range ps {
			if i%p == 0 {
				fs[p] = fs[p] + 1
				i /= p
			}
		}
	}
	cache[n] = fs
	return fs
}

func output(out *os.File, n int, fs map[int]int, ps []int) {
	fmt.Fprintf(out, "%3d! =", n)
	l := len(fs)
	cnt := 0
	for _, v := range ps {
		fmt.Fprintf(out, "%3d", fs[v])
		if l--; l == 0 {
			break
		}
		if cnt++; cnt == 15 {
			cnt = 0
			fmt.Fprint(out, "\n      ")
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("160.in")
	defer in.Close()
	out, _ := os.Create("160.out")
	defer out.Close()

	ps := primes(100)
	var n int
	cache = make(map[int]map[int]int)
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fs := make(map[int]int)
		for i := 2; i <= n; i++ {
			tmp := factors(i, ps)
			for k, v := range tmp {
				fs[k] = fs[k] + v
			}
		}
		output(out, n, fs, ps)
	}
}
