// UVa 482 - Permutation Arrays

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func index(line string) map[int]int {
	var m int
	r := strings.NewReader(line)
	idx := make(map[int]int)
	for cnt := 0; ; cnt++ {
		if _, err := fmt.Fscanf(r, "%d", &m); err != nil {
			break
		}
		idx[m-1] = cnt
	}
	return idx
}

func number(line string) []string {
	var s string
	r := strings.NewReader(line)
	var num []string
	for {
		if _, err := fmt.Fscanf(r, "%s", &s); err != nil {
			break
		}
		num = append(num, s)
	}
	return num
}

func main() {
	in, _ := os.Open("482.in")
	defer in.Close()
	out, _ := os.Create("482.out")
	defer out.Close()

	var kase int
	fmt.Fscanf(in, "%d", &kase)

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	for kase > 0 {
		s.Scan()
		s.Text()

		s.Scan()
		idx := index(s.Text())
		s.Scan()
		num := number(s.Text())

		for i := 0; i < len(idx); i++ {
			fmt.Fprintln(out, num[idx[i]])
		}
		if kase--; kase != 0 {
			fmt.Fprintln(out)
		}
	}
}
