// UVa 196 - Spreadsheet

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var spreadsheet [][]string

func isAlphabet(b byte) bool { return b >= 'A' && b <= 'Z' }

func convert(cellName string) (int, int) {
	var r, c, chars int
	for i := range cellName {
		if isAlphabet(cellName[i]) {
			c = c*26 + int(cellName[i]-'A')
			chars++
		} else {
			r = r*10 + int(cellName[i]-'0')
		}
	}
	switch chars {
	case 2:
		c += 26
	case 3:
		c += 26*26 + 26
	}
	return r - 1, c
}

func resolve(r, c int) int {
	if strings.HasPrefix(spreadsheet[r][c], "=") {
		var total int
		for _, cellName := range strings.Split(spreadsheet[r][c][1:], "+") {
			total += resolve(convert(cellName))
		}
		spreadsheet[r][c] = strconv.Itoa(total)
	}
	v, _ := strconv.Atoi(spreadsheet[r][c])
	return v
}

func solve() {
	for i, row := range spreadsheet {
		for j := range row {
			resolve(i, j)
		}
	}
}

func main() {
	in, _ := os.Open("196.in")
	defer in.Close()
	out, _ := os.Create("196.out")
	defer out.Close()

	var n, c, r int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &c, &r)
		spreadsheet = make([][]string, r)
		for i := range spreadsheet {
			spreadsheet[i] = make([]string, c)
			for j := range spreadsheet[i] {
				fmt.Fscanf(in, "%s", &spreadsheet[i][j])
			}
		}
		solve()
		for _, row := range spreadsheet {
			fmt.Fprintln(out, strings.Join(row, " "))
		}
	}
}
