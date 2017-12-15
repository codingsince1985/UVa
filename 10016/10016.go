// UVa 10016 - Flip-Flop the Squarelotron

package main

import (
	"fmt"
	"os"
	"strings"
)

func upsideDown(matrix [][]string, n, ring int) {
	row1, row2 := ring, n-1-ring
	for i := ring; i < n-ring; i++ {
		matrix[row1][i], matrix[row2][i] = matrix[row2][i], matrix[row1][i]
	}
	col1, col2 := ring, n-1-ring
	for row := row1 + 1; row < n/2; row++ {
		matrix[row][col1], matrix[n-row-1][col1] = matrix[n-row-1][col1], matrix[row][col1]
		matrix[row][col2], matrix[n-row-1][col2] = matrix[n-row-1][col2], matrix[row][col2]
	}
}

func leftRight(matrix [][]string, n, ring int) {
	col1, col2 := ring, n-1-ring
	for i := ring; i < n-ring; i++ {
		matrix[i][col1], matrix[i][col2] = matrix[i][col2], matrix[i][col1]
	}
	row1, row2 := ring, n-1-ring
	for col := col1 + 1; col < n/2; col++ {
		matrix[row1][col], matrix[row1][n-col-1] = matrix[row1][n-col-1], matrix[row1][col]
		matrix[row2][col], matrix[row2][n-col-1] = matrix[row2][n-col-1], matrix[row2][col]
	}
}

func mainDiagonal(matrix [][]string, n, ring int) {
	for col := ring + 1; col < n-ring; col++ {
		matrix[ring][col], matrix[col][ring] = matrix[col][ring], matrix[ring][col]
	}
	for col := ring + 1; col < n-1-ring; col++ {
		matrix[n-1-ring][col], matrix[col][n-1-ring] = matrix[col][n-1-ring], matrix[n-1-ring][col]
	}
}

func mainInverseDiagonal(matrix [][]string, n, ring int) {
	for col := ring; col < n-1-ring; col++ {
		matrix[ring][col], matrix[n-1-col][n-1-ring] = matrix[n-1-col][n-1-ring], matrix[ring][col]
	}
	for col := ring + 1; col < n-1-ring; col++ {
		matrix[n-1-ring][col], matrix[n-1-col][ring] = matrix[n-1-col][ring], matrix[n-1-ring][col]
	}
}

func main() {
	in, _ := os.Open("10016.in")
	defer in.Close()
	out, _ := os.Create("10016.out")
	defer out.Close()

	var m, n, t, c int
	for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
		fmt.Fscanf(in, "%d", &n)
		matrix := make([][]string, n)
		for i := range matrix {
			matrix[i] = make([]string, n)
			for j := range matrix[i] {
				fmt.Fscanf(in, "%s", &matrix[i][j])
			}
		}
		for ring := 0; ring < (n+1)/2; ring++ {
			for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
				switch fmt.Fscanf(in, "%d", &c); c {
				case 1:
					upsideDown(matrix, n, ring)
				case 2:
					leftRight(matrix, n, ring)
				case 3:
					mainDiagonal(matrix, n, ring)
				case 4:
					mainInverseDiagonal(matrix, n, ring)
				}
			}
		}
		for _, row := range matrix {
			fmt.Fprintln(out, strings.Join(row, " "))
		}
	}
}
