// UVa 10928 - My Dear Neighbours

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(places [][]int) []string {
	min := math.MaxInt32
	for _, place := range places {
		if l := len(place); l < min {
			min = l
		}
	}
	var placeToBe []string
	for i, place := range places {
		if l := len(place); l == min {
			placeToBe = append(placeToBe, strconv.Itoa(i+1))
		}
	}
	return placeToBe
}

func main() {
	in, _ := os.Open("10928.in")
	defer in.Close()
	out, _ := os.Create("10928.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, place, p int
	s.Scan()
	for fmt.Sscanf(s.Text(), "%d", &n); n > 0 && s.Scan(); n-- {
		fmt.Sscanf(s.Text(), "%d", &p)
		places := make([][]int, p)
		for i := range places {
			s.Scan()
			for r := strings.NewReader(s.Text()); ; {
				if _, err := fmt.Fscanf(r, "%d", &place); err != nil {
					break
				}
				places[i] = append(places[i], place)
			}
		}
		fmt.Fprintln(out, strings.Join(solve(places), " "))
		s.Scan()
	}
}
