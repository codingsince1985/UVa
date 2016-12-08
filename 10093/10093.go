// UVa 10093 - An Easy Problem!

package main

import (
	"fmt"
	"os"
)

func number(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'A' && c <= 'Z':
		return int(10 + c - 'A')
	case c >= 'a' && c <= 'z':
		return int(36 + c - 'a')
	default:
		return 0
	}
}

func solve(line string) int {
	var sum, max int
	for i := range line {
		num := number(line[i])
		sum += num
		if num > max {
			max = num
		}
	}
	for i := max; i <= 62; i++ {
		if sum%i == 0 {
			return i + 1
		}
	}
	return 0
}

func main() {
	in, _ := os.Open("10093.in")
	defer in.Close()
	out, _ := os.Create("10093.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		if base := solve(line); base == 0 {
			fmt.Fprintln(out, "such number is impossible!")
		} else {
			fmt.Fprintln(out, base)
		}
	}
}
