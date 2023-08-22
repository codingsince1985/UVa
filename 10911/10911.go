// UVa 10911 - Forming Quiz Teams

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	num  uint32
	dp   []float64
	dist [][]float64
)

func distance(s1, s2 [2]int) float64 {
	return math.Sqrt(float64((s1[0]-s2[0])*(s1[0]-s2[0]) + (s1[1]-s2[1])*(s1[1]-s2[1])))
}

func dfs(level int) float64 {
	if dp[level] == math.MaxFloat64 {
		if level == (1<<num)-1 {
			dp[level] = 0
		} else {
			for i := uint32(0); i < num-1; i++ {
				if level&(1<<i) == 0 {
					for j := i + 1; j < num; j++ {
						if level&(1<<j) == 0 {
							dp[level] = min(dp[level], dist[i][j]+dfs(level|1<<i|1<<j))
						}
					}
				}
			}
		}
	}
	return dp[level]
}

func solve(students [][2]int) float64 {
	dist = make([][]float64, num)
	for i := range dist {
		dist[i] = make([]float64, num)
		for j := range dist[i] {
			dist[i][j] = distance(students[i], students[j])
		}
	}
	dp = make([]float64, 1<<num)
	for i := range dp {
		dp[i] = math.MaxFloat64
	}
	return dfs(0)
}

func main() {
	in, _ := os.Open("10911.in")
	defer in.Close()
	out, _ := os.Create("10911.out")
	defer out.Close()

	var n uint32
	var name string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num = 2 * n
		students := make([][2]int, num)
		for i := range students {
			fmt.Fscanf(in, "%s%d%d", &name, &students[i][0], &students[i][1])
		}
		fmt.Fprintf(out, "Case %d: %.2f\n", kase, solve(students))
	}
}
