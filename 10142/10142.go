// UVa 10142 - Australian Voting

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var out *os.File

func findLowest(total []int, min int) map[int]bool {
	lowest := make(map[int]bool)
	for i, v := range total {
		if v == min {
			lowest[i+1] = true
		}
	}
	return lowest
}

func eliminate(candidates []string, votes [][]int, eliminated map[int]bool) {
	for i, vi := range votes {
		var newVote []int
		for _, vj := range vi {
			if !eliminated[vj] {
				newVote = append(newVote, vj)
			}
		}
		votes[i] = newVote
	}
	for i := range candidates {
		if eliminated[i+1] {
			candidates[i] = ""
		}
	}
}

func count(candidates []string, votes [][]int) {
	total := make([]int, len(candidates))
	for _, v := range votes {
		if len(v) > 0 {
			total[v[0]-1]++
		}
	}
	max, min := 0, math.MaxInt32
	for i, v := range total {
		if candidates[i] != "" {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
	}
	switch {
	case max > len(votes)/2:
		for i, v := range total {
			if v == max {
				fmt.Fprintln(out, candidates[i])
			}
		}
	case max == min:
		for _, v := range candidates {
			if v != "" {
				fmt.Fprintln(out, v)
			}
		}
	default:
		lowest := findLowest(total, min)
		eliminate(candidates, votes, lowest)
		count(candidates, votes)
	}
}

func main() {
	in, _ := os.Open("10142.in")
	defer in.Close()
	out, _ = os.Create("10142.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	s.Text()
	for ; s.Scan() && kase > 0; kase-- {
		n, _ := strconv.Atoi(s.Text())
		candidates := make([]string, n)
		for i := range candidates {
			s.Scan()
			candidates[i] = s.Text()
		}
		var votes [][]int
		var line string
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			r := strings.NewReader(line)
			vote := make([]int, n)
			for i := range vote {
				fmt.Fscanf(r, "%d", &vote[i])
			}
			votes = append(votes, vote)
		}
		count(candidates, votes)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
