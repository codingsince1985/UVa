// UVa 542 - France '98

package main

import (
	"fmt"
	"os"
)

const TEAMS = 16

func solve(matrix [][]float64) []float64 {
	dp := make([][][]float64, TEAMS)
	for i := range dp {
		dp[i] = make([][]float64, TEAMS)
		for j := range dp[i] {
			dp[i][j] = make([]float64, TEAMS)
		}
		dp[i][i][i] = 1
	}
	for teamsPerGroup := 2; teamsPerGroup <= 16; teamsPerGroup = teamsPerGroup * 2 {
		for left := 0; left <= TEAMS-teamsPerGroup; left += teamsPerGroup {
			right := left + teamsPerGroup - 1
			mid := (left + right) / 2
			for i := left; i <= mid; i++ {
				for j := mid + 1; j <= right; j++ {
					dp[i][left][right] += matrix[i][j] * dp[i][left][mid] * dp[j][mid+1][right]
					dp[j][left][right] += matrix[j][i] * dp[i][left][mid] * dp[j][mid+1][right]
				}
			}
		}
	}
	probabilities := make([]float64, TEAMS)
	for i := range probabilities {
		probabilities[i] = dp[i][0][TEAMS-1] * 100
	}
	return probabilities
}

func main() {
	in, _ := os.Open("542.in")
	defer in.Close()
	out, _ := os.Create("542.out")
	defer out.Close()

	countries := make([]string, TEAMS)
	for i := range countries {
		fmt.Fscanf(in, "%s", &countries[i])
	}
	matrix := make([][]float64, TEAMS)
	for i := range matrix {
		matrix[i] = make([]float64, TEAMS)
		for j := range matrix[i] {
			fmt.Fscanf(in, "%f", &matrix[i][j])
			matrix[i][j] /= 100
		}
	}
	for i, probability := range solve(matrix) {
		fmt.Fprintf(out, "%-10s p=%.2f%%\n", countries[i], probability)
	}
}
