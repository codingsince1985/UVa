// UVa 333 - Recognizing Good ISBNs

package main

import (
	"fmt"
	"os"
)

func sum(digits []int) []int {
	sum := make([]int, 10)
	total := 0
	for i, d := range digits {
		total += d
		sum[i] = total
	}
	return sum
}

func isValid(isbn string) bool {
	var digits []int
	for i := range isbn {
		switch {
		case isbn[i] >= '0' && isbn[i] <= '9':
			digits = append(digits, int(isbn[i]-'0'))
		case isbn[i] == 'X':
			if len(digits) != 9 {
				return false
			}
			digits = append(digits, 10)
		case isbn[i] == '-':
			continue
		default:
			return false
		}
	}
	if len(digits) != 10 {
		return false
	}
	return sum(sum(digits))[9]%11 == 0
}

func main() {
	in, _ := os.Open("333.in")
	defer in.Close()
	out, _ := os.Create("333.out")
	defer out.Close()

	var line string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		if isValid(line) {
			fmt.Fprintf(out, "%s is correct.\n", line)
		} else {
			fmt.Fprintf(out, "%s is incorrect.\n", line)
		}
	}
}
