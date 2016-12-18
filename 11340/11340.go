// UVa 11340 - Newspaper

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	kase, _ := strconv.Atoi(s.Text())
	for kase > 0 && s.Scan() {
		k, _ := strconv.Atoi(s.Text())
		paid := make(map[byte]int)
		for k > 0 && s.Scan() {
			r := strings.NewReader(s.Text())
			fmt.Fscanf(r, "%c %d", &c, &t)
			paid[c] = t
			k--
		}
		s.Scan()
		l, _ := strconv.Atoi(s.Text())
		total := 0.0
		for l > 0 && s.Scan() {
			line := s.Text()
			for i := range line {
				total += float64(paid[line[i]])
			}
			l--
		}
		fmt.Fprintf(out, "%.2f$\n", total/100)
		kase--
	}
}
