// UVa 10201 - Adventures in Moving - Part IV

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	max  = 200
	half = 100
)

type station struct{ distance, price int }

func solve(distance int, stations []station) int {
	if stations[0].distance != 0 {
		stations = append([]station{{0, 0}}, stations...)
	}
	if stations[len(stations)-1].distance != distance {
		stations = append(stations, station{distance, math.MaxInt32})
	}
	l := len(stations)
	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, max+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][half] = 0
	for i := 1; i < l; i++ {
		dist := stations[i].distance - stations[i-1].distance
		for j := dist; j <= max; j++ {
			dp[i][j-dist] = min(dp[i][j-dist], dp[i-1][j])
		}
		for j := 1; j <= max; j++ {
			dp[i][j] = min(dp[i][j], dp[i][j-1]+stations[i].price)
		}
	}
	lowest := math.MaxInt32
	for i := half; i <= max; i++ {
		lowest = min(lowest, dp[l-1][i])
	}
	return lowest
}

func main() {
	in, _ := os.Open("10201.in")
	defer in.Close()
	out, _ := os.Create("10201.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	var kase, d, p, distance int
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &kase)
	for s.Scan(); kase > 0 && s.Scan(); kase-- {
		fmt.Sscanf(s.Text(), "%d", &distance)
		var stations []station
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(line, "%d%d", &d, &p)
			stations = append(stations, station{d, p})
		}
		fmt.Fprintln(out, solve(distance, stations))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
