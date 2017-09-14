// UVa 10801 - Lift Hopping

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const max = 100

func spfa(n, s, t int, matrix [][]int) int {
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[s] = 0
	for queue := []int{s}; len(queue) > 0; {
		curr := queue[0]
		queue = queue[1:]
		for to, time := range matrix[curr] {
			if newDistance := distance[curr] + time + 60; distance[to] > newDistance {
				distance[to] = newDistance
				queue = append(queue, to)
			}
		}
	}
	return distance[t]
}

func solve(k int, speeds []int, elevators [][]int) int {
	matrix := make([][]int, max)
	for i := range matrix {
		matrix[i] = make([]int, max)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}
	for i, elevator := range elevators {
		for j, floor := range elevator {
			for k := j + 1; k < len(elevator); k++ {
				if distance := (elevator[k] - floor) * speeds[i]; distance < matrix[floor][elevator[k]] {
					matrix[floor][elevator[k]], matrix[elevator[k]][floor] = distance, distance
				}
			}
		}
	}
	return spfa(max, 0, k, matrix)
}

func main() {
	in, _ := os.Open("10801.in")
	defer in.Close()
	out, _ := os.Create("10801.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, k, tmp int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d%d", &n, &k)
		speeds := make([]int, n)
		s.Scan()
		r := strings.NewReader(s.Text())
		for i := range speeds {
			fmt.Fscanf(r, "%d", &speeds[i])
		}
		elevators := make([][]int, n)
		for i := range elevators {
			s.Scan()
			for r := strings.NewReader(s.Text()); ; {
				if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
					break
				}
				elevators[i] = append(elevators[i], tmp)
			}
		}
		distance := solve(k, speeds, elevators)
		switch distance {
		case math.MaxInt32:
			fmt.Fprintln(out, "IMPOSSIBLE")
		case 0:
			fmt.Fprintln(out, 0)
		default:
			fmt.Fprintln(out, distance-60)
		}
	}
}
