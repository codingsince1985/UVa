// UVa 10374 - Election

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(max int, candidateMap map[string]string, voteMap map[string]int) string {
	var name string
	for k, v := range voteMap {
		if v == max {
			if name != "" {
				return "tie"
			}
			name = k
		}
	}
	return candidateMap[name]
}

func main() {
	in, _ := os.Open("10374.in")
	defer in.Close()
	out, _ := os.Create("10374.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	first := true
	for kase, _ := strconv.Atoi(s.Text()); kase > 0; kase-- {
		s.Scan()
		s.Text()
		candidateMap := make(map[string]string)
		s.Scan()
		for n, _ := strconv.Atoi(s.Text()); n > 0; n-- {
			s.Scan()
			name := s.Text()
			s.Scan()
			candidateMap[name] = s.Text()
		}
		voteMap := make(map[string]int)
		var max int
		s.Scan()
		for n, _ := strconv.Atoi(s.Text()); n > 0; n-- {
			s.Scan()
			if name := s.Text(); candidateMap[name] != "" {
				voteMap[name]++
				if voteMap[name] > max {
					max = voteMap[name]
				}
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "%s\n", solve(max, candidateMap, voteMap))
	}
}
