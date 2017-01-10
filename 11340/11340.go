// UVa 11340 - Newspaper

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("11340.in")
	defer in.Close()
	out, _ := os.Create("11340.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var t int
	var c byte
	s.Scan()
	for kase, _ := strconv.Atoi(s.Text()); kase > 0 && s.Scan(); kase-- {
		paid := make(map[byte]int)
		for k, _ := strconv.Atoi(s.Text()); k > 0 && s.Scan(); k-- {
			fmt.Sscanf(s.Text(), "%c%d", &c, &t)
			paid[c] = t
		}
		total := 0.0
		s.Scan()
		for l, _ := strconv.Atoi(s.Text()); l > 0 && s.Scan(); l-- {
			line := s.Text()
			for i := range line {
				total += float64(paid[line[i]])
			}
		}
		fmt.Fprintf(out, "%.2f$\n", total/100)
	}
}
