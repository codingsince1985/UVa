// UVa 10115 - Automatic Editing

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(text string, rules [][2]string) string {
	for _, rule := range rules {
		for strings.Contains(text, rule[0]) {
			text = strings.Replace(text, rule[0], rule[1], 1)
		}
	}
	return text
}

func main() {
	in, _ := os.Open("10115.in")
	defer in.Close()
	out, _ := os.Create("10115.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var r int
	for s.Scan() {
		if r, _ = strconv.Atoi(s.Text()); r == 0 {
			break
		}
		rules := make([][2]string, r)
		for i := range rules {
			s.Scan()
			s1 := s.Text()
			s.Scan()
			rules[i] = [2]string{s1, s.Text()}
		}
		s.Scan()
		fmt.Fprintln(out, solve(s.Text(), rules))
	}
}
