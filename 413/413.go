// UVa 413 - Up and Down Sequences

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(num []int) (float64, float64) {
	var ups, downs, upseqs, downseqs, unknown int
	var direction byte
	for i, n := range num {
		if i == 0 {
			direction = 'N'
			continue
		}
		switch {
		case n == num[i-1]:
			switch direction {
			case 'N':
				unknown++
			case 'U':
				ups++
			default:
				downs++
			}
		case n > num[i-1]:
			switch direction {
			case 'D':
				downseqs++
			case 'N':
				ups += unknown
				unknown = 0
			}
			ups++
			direction = 'U'
		default:
			switch direction {
			case 'U':
				upseqs++
			case 'N':
				downs += unknown
				unknown = 0
			}
			downs++
			direction = 'D'
		}
	}
	switch direction {
	case 'U':
		upseqs++
	case 'D':
		downseqs++
	}
	var upAvg, downAvg float64
	if ups > 0 {
		upAvg = float64(ups) / float64(upseqs)
	}
	if downs > 0 {
		downAvg = float64(downs) / float64(downseqs)
	}
	return upAvg, downAvg
}

func main() {
	in, _ := os.Open("413.in")
	defer in.Close()
	out, _ := os.Create("413.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var tmp int
	var line string
	for s.Scan() {
		if line = s.Text(); line == "0" {
			break
		}
		r := strings.NewReader(line)
		var num []int
		for {
			if fmt.Fscanf(r, "%d", &tmp); tmp == 0 {
				break
			}
			num = append(num, tmp)
		}
		upAvg, downAvg := solve(num)
		fmt.Fprintf(out, "Nr values = %d: %.6f %.6f\n", len(num), upAvg, downAvg)
	}
}
