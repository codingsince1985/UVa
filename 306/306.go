// UVa 306 - Cipher

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert(a []int) map[int][]int {
	m := make(map[int][]int)
	for i := range a {
		tmp := a[i]
		m[i+1] = append(m[i+1], tmp)
		for a[tmp-1] != a[i] {
			m[i+1] = append(m[i+1], a[tmp-1])
			tmp = a[tmp-1]
		}
	}
	return m
}

func solve(m map[int][]int, n, k int, line string) string {
	if len(line) < n {
		line += strings.Repeat(" ", n-len(line))
	}
	encoded := make([]byte, n)
	for i := range line {
		encoded[m[i+1][(k-1)%len(m[i+1])]-1] = line[i]
	}
	return string(encoded)
}

func main() {
	in, _ := os.Open("306.in")
	defer in.Close()
	out, _ := os.Create("306.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, k int
	var line string
	for s.Scan() {
		if n, _ = strconv.Atoi(s.Text()); n == 0 {
			break
		}
		s.Scan()
		r := strings.NewReader(s.Text())
		a := make([]int, n)
		for i := range a {
			fmt.Fscanf(r, "%d", &a[i])
		}
		m := convert(a)
		for s.Scan() {
			if line = s.Text(); line == "0" {
				fmt.Fprintln(out)
				break
			}
			tokens := strings.SplitN(line, " ", 2)
			k, _ = strconv.Atoi(tokens[0])
			fmt.Fprintln(out, solve(m, n, k, tokens[1]))
		}
	}
}
