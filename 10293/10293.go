// UVa 10293 - Word Length and Frequency

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(thesis string) (map[int]int, int) {
	thesis = strings.Replace(thesis, "-", "", -1)
	lengthMap := make(map[int]int)
	var started bool
	var count, max int
	for i := range thesis {
		if thesis[i] == ' ' || thesis[i] == '?' || thesis[i] == '!' || thesis[i] == ',' || thesis[i] == '.' {
			if started {
				lengthMap[count]++
				if count > max {
					max = count
				}
				started = false
			}
		} else {
			if !started {
				started, count = true, 0
			}
			if thesis[i] != '\'' {
				count++
			}
		}
	}
	return lengthMap, max
}

func output(out *os.File, lengthMap map[int]int, max int) {
	for i := 1; i <= max; i++ {
		if lengthMap[i] > 0 {
			fmt.Fprintf(out, "%d %d\n", i, lengthMap[i])
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10293.in")
	defer in.Close()
	out, _ := os.Create("10293.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		var thesis string
		for ; line != "#"; line = s.Text() {
			thesis += line
			if !strings.HasSuffix(line, "-") {
				thesis += " "
			}
			s.Scan()
		}
		lengthMap, max := solve(thesis)
		output(out, lengthMap, max)
	}
}
