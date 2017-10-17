// UVa 10278 - Fire Station

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func floydWarshall(matrix [][]int) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k][j])
			}
		}
	}
}

func distanceToFireStation(fireStations []int, matrix [][]int) (int, []int) {
	var longest int
	distances := make([]int, len(matrix))
	for i := range distances {
		distances[i] = math.MaxInt32
		for _, fireStation := range fireStations {
			distances[i] = min(distances[i], matrix[i][fireStation-1])
		}
		longest = max(longest, distances[i])
	}
	return longest, distances
}

func solve(fireStations []int, matrix [][]int) int {
	floydWarshall(matrix)
	longest, distances := distanceToFireStation(fireStations, matrix)
	var newFireStation int
	for i := range matrix {
		var maxDistance int
		for j := range matrix {
			maxDistance = max(maxDistance, min(matrix[i][j], distances[j]))
		}
		if maxDistance < longest {
			longest = maxDistance
			newFireStation = i + 1
		}
	}
	return newFireStation
}

func main() {
	in, _ := os.Open("10278.in")
	defer in.Close()
	out, _ := os.Create("10278.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, f, inter, c1, c2, distance int
	var line string
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &kase)
	for s.Scan(); kase > 0 && s.Scan(); kase-- {
		fmt.Sscanf(s.Text(), "%d%d", &f, &inter)
		fireStations := make([]int, f)
		for i := range fireStations {
			s.Scan()
			fmt.Sscanf(s.Text(), "%d", &fireStations[i])
		}
		matrix := make([][]int, inter)
		for i := range matrix {
			matrix[i] = make([]int, inter)
			for j := range matrix[i] {
				if i != j {
					matrix[i][j] = math.MaxInt32
				}
			}
		}
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%d%d%d", &c1, &c2, &distance)
			matrix[c1-1][c2-1], matrix[c2-1][c1-1] = distance, distance
		}
		fmt.Fprintln(out, solve(fireStations, matrix))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
