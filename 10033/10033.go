// UVa 10033 - Interpreter

package main

import (
	"fmt"
	"os"
)

func solve(lines []string) int {
	registers := make([]int, 10)
	ram := make([]int, 1000)
	cnt := 0
	pointer := 0
	for {
		line := lines[pointer]
		jump := false
		switch {
		case line == "100":
			cnt++
			return cnt
		case line[0] == '2':
			registers[line[1]-'0'] = int(line[2] - '0')
		case line[0] == '3':
			registers[line[1]-'0'] += int(line[2] - '0')
		case line[0] == '4':
			registers[line[1]-'0'] *= int(line[2] - '0')
		case line[0] == '5':
			registers[line[1]-'0'] = registers[line[2]-'0']
		case line[0] == '6':
			registers[line[1]-'0'] += registers[line[2]-'0']
		case line[0] == '7':
			registers[line[1]-'0'] *= registers[line[2]-'0']
		case line[0] == '8':
			registers[line[1]-'0'] = ram[registers[line[2]-'0']]
		case line[0] == '9':
			ram[registers[line[2]-'0']] = registers[line[1]-'0']
		case line[0] == '0' && registers[line[2]-'0'] != 0:
			jump = true
			pointer = registers[line[1]-'0']
		}
		registers[line[1]-'0'] %= 1000
		cnt++
		if jump {
			jump = false
		} else {
			pointer++
		}
	}
}

func main() {
	in, _ := os.Open("10033.in")
	defer in.Close()
	out, _ := os.Create("10033.out")
	defer out.Close()

	var kase int
	var line string
	first := true
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		if first {
			fmt.Fscanln(in)
			first = false
		} else {
			fmt.Fprintln(out)
		}
		var lines []string
		for {
			if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
				break
			}
			lines = append(lines, line)
		}
		fmt.Fprintln(out, solve(lines))
		kase--
	}
}
