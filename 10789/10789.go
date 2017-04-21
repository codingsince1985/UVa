// UVa 10789 - Prime Frequency

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primes() map[int]bool {
	p := map[int]bool{2: true}
	for i := 3; i < 2001; i += 2 {
		if isPrime(i) {
			p[i] = true
		}
	}
	return p
}

func chars() []byte {
	var c []byte
	var i byte
	for i = '0'; i <= '9'; i++ {
		c = append(c, i)
	}
	for i = 'A'; i <= 'Z'; i++ {
		c = append(c, i)
	}
	for i = 'a'; i <= 'z'; i++ {
		c = append(c, i)
	}
	return c
}

func main() {
	in, _ := os.Open("10789.in")
	defer in.Close()
	out, _ := os.Create("10789.out")
	defer out.Close()

	p := primes()
	c := chars()
	var t int
	var line string
	fmt.Fscanf(in, "%d", &t)
	for i := 1; i <= t; i++ {
		fmt.Fscanf(in, "%s", &line)
		m := make(map[byte]int)
		for j := range line {
			m[line[j]]++
		}

		empty := true
		fmt.Fprintf(out, "Case %d: ", i)
		for _, v := range c {
			if n, ok := m[v]; ok && p[n] {
				empty = false
				fmt.Fprintf(out, "%c", v)
			}
		}
		if empty {
			fmt.Fprint(out, "empty")
		}
		fmt.Fprintln(out)
	}
}
