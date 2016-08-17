// UVa 516 - Prime Land

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX = 32767

func sieve() []bool {
	p := make([]bool, MAX+1)
	for i := 2; i*i <= MAX; i++ {
		if !p[i] {
			for j := 2 * i; j <= MAX; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func pow(b, e int) int {
	num := 1
	for e > 0 {
		num *= b
		e--
	}
	return num
}

func factorize(primes []bool, n int) [][2]int {
	var factors [][2]int
	f := n
	for n > 1 {
		if !primes[f] && n%f == 0 {
			e := 0
			for n > 1 && n%f == 0 {
				n /= f
				e++
			}
			factors = append(factors, [2]int{f, e})
			f = n
		} else {
			f--
		}
	}
	return factors
}

func main() {
	in, _ := os.Open("516.in")
	defer in.Close()
	out, _ := os.Create("516.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	primes := sieve()
	var b, e int
	for s.Scan() {
		if line := s.Text(); line != "0" {
			r := strings.NewReader(line)
			num := 1
			for {
				if _, err := fmt.Fscanf(r, "%d%d", &b, &e); err != nil {
					break
				}
				num *= pow(b, e)
			}
			factors := factorize(primes, num-1)
			for i, v := range factors {
				if i != 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, v[0], v[1])
			}
			fmt.Fprintln(out)
		}
	}
}
