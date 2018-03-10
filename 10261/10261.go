// UVa 10261 - Ferry Loading

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	out io.WriteCloser
	pre [][]int
)

func output(index, length int) {
	if index == 0 {
		return
	}
	output(index-1, pre[index][length])
	if pre[index][length] == length {
		fmt.Fprintln(out, "port")
	} else {
		fmt.Fprintln(out, "starboard")
	}
}

func solve(ferry int, cars []int) {
	num := len(cars)
	dp := make([][]bool, num)
	pre = make([][]int, num)
	for i := range dp {
		dp[i] = make([]bool, ferry+1)
		pre[i] = make([]int, ferry+1)
	}
	dp[0][0] = true
	var sum, last, length int
	for i := 1; i <= num; i++ {
		sum += cars[i-1]
		for j := 0; j <= ferry; j++ {
			switch {
			case sum-j <= ferry && dp[i-1][j]:
				pre[i][j] = j
				dp[i][j], last, length = true, i, j
			case j >= cars[i-1] && dp[i-1][j-cars[i-1]]:
				pre[i][j] = j - cars[i-1]
				dp[i][j], last, length = true, i, j
			}
		}
	}
	fmt.Fprintln(out, last)
	output(last, length)
}

func main() {
	in, _ := os.Open("10261.in")
	defer in.Close()
	out, _ = os.Create("10261.out")
	defer out.Close()

	var kase int
	var ferry, car int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &ferry)
		var cars []int
		for {
			if fmt.Fscanf(in, "%d", &car); car == 0 {
				break
			}
			cars = append(cars, car)
		}
		solve(ferry*100, cars)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
