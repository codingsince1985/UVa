// UVa 755 - 487--3279

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var charMap = map[byte]int{
	'A': 2, 'B': 2, 'C': 2,
	'D': 3, 'E': 3, 'F': 3,
	'G': 4, 'H': 4, 'I': 4,
	'J': 5, 'K': 5, 'L': 5,
	'M': 6, 'N': 6, 'O': 6,
	'P': 7, 'R': 7, 'S': 7,
	'T': 8, 'U': 8, 'V': 8,
	'W': 9, 'X': 9, 'Y': 9,
}

func solve(telephones []string) map[string]int {
	numbers := make(map[string]int)
	for _, telephone := range telephones {
		var number string
		for j := range telephone {
			switch {
			case telephone[j] >= '0' && telephone[j] <= '9':
				number += string(telephone[j])
			case telephone[j] >= 'A' && telephone[j] <= 'Z':
				number += strconv.Itoa(charMap[telephone[j]])
			}
		}
		numbers[number]++
	}
	return numbers
}

func main() {
	in, _ := os.Open("755.in")
	defer in.Close()
	out, _ := os.Create("755.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		telephones := make([]string, n)
		for i := range telephones {
			fmt.Fscanf(in, "%s", &telephones[i])
		}
		numbers := solve(telephones)
		var dups []string
		for k, v := range numbers {
			if v > 1 {
				dups = append(dups, k)
			}
		}
		if len(dups) == 0 {
			fmt.Fprintln(out, "No duplicates.")
		} else {
			sort.Strings(dups)
			for _, dup := range dups {
				fmt.Fprintf(out, "%s-%s %d\n", dup[:3], dup[3:], numbers[dup])
			}
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
