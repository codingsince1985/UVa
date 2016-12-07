// UVa 10922 - 2 the 9s

package main

import (
	"fmt"
	"os"
	"strconv"
)

func digitSum(num string) int {
	var total int
	for i := range num {
		total += int(num[i] - '0')
	}
	return total
}

func nineDegree(num int) int {
	var degree int
	for num%9 == 0 {
		degree++
		newNum := digitSum(strconv.Itoa(num))
		if newNum == num {
			break
		}
		num = newNum
	}
	return degree
}

func main() {
	in, _ := os.Open("10922.in")
	defer in.Close()
	out, _ := os.Create("10922.out")
	defer out.Close()

	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "0" {
			break
		}
		if degree := nineDegree(digitSum(line)); degree == 0 {
			fmt.Fprintf(out, "%s is not a multiple of 9.\n", line)
		} else {
			fmt.Fprintf(out, "%s is a multiple of 9 and has 9-degree %d.\n", line, degree)
		}
	}
}
