// UVa 161 - Traffic Lights

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(cycles []int) int {
	first := true
	for i := 1; i <= 18000; i++ {
		allGreen := true
		for _, cycle := range cycles {
			if i%(2*cycle) >= cycle-5 {
				allGreen = false
				break
			}
		}
		if !allGreen {
			if first {
				first = false
			}
		} else {
			if !first {
				return i
			}
		}
	}
	return 0
}

func main() {
	in, _ := os.Open("161.in")
	defer in.Close()
	out, _ := os.Create("161.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	var cycle int
	var cycles []int
	for s.Scan() {
		if line = s.Text(); line == "0 0 0" {
			break
		}
		r := strings.NewReader(line)
		for {
			if _, err := fmt.Fscanf(r, "%d", &cycle); err != nil || cycle == 0 {
				break
			}
			cycles = append(cycles, cycle)
		}
		if cycle == 0 {
			seconds := solve(cycles)
			if seconds == 0 {
				fmt.Fprintln(out, "Signals fail to synchronise in 5 hours")
			} else {
				hours := seconds / 3600
				minutes := seconds % 3600 / 60
				seconds = seconds % 3600 % 60
				fmt.Fprintf(out, "%02d:%02d:%02d\n", hours, minutes, seconds)
			}
			cycles = nil
		}
	}
}
