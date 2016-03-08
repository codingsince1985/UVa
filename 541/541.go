// UVa 541 - Error Correction

package main

import (
	"fmt"
	"os"
)

func cell(matrix [][]byte, i, j int, forRow bool) byte {
	if forRow {
		return matrix[i][j]
	} else {
		return matrix[j][i]
	}
}

func invalid(matrix [][]byte, forRow bool) []int {
	invalids := make([]int, 0)
	for i := range matrix {
		isValid := true
		for j := range matrix[i] {
			if cell(matrix, i, j, forRow) == 1 {
				isValid = !isValid
			}
		}
		if !isValid {
			invalids = append(invalids, i)
		}
	}
	return invalids
}

func main() {
	in, _ := os.Open("541.in")
	defer in.Close()
	out, _ := os.Create("541.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		matrix := make([][]byte, n)
		for i := range matrix {
			matrix[i] = make([]byte, n)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%d", &matrix[i][j])
			}
		}
		invalidRow, invalidColumn := invalid(matrix, true), invalid(matrix, false)
		switch {
		case len(invalidRow) == 0 && len(invalidColumn) == 0:
			fmt.Fprintln(out, "OK")
		case len(invalidRow) == 1 && len(invalidColumn) == 1:
			fmt.Fprintf(out, "Change bit (%d,%d)\n", invalidRow[0]+1, invalidColumn[0]+1)
		default:
			fmt.Fprintln(out, "Corrupt")
		}
	}
}
