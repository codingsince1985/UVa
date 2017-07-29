// UVa 11040 - Add bricks in the wall

package main

import (
	"fmt"
	"os"
)

const row = 9

func solve(wall [][]int) {
	for i := 2; i < row; i += 2 {
		for j := 1; j < i; j += 2 {
			wall[i][j] = (wall[i-2][j-1] - wall[i][j-1] - wall[i][j+1]) / 2
		}
		for j := 0; j < i; j++ {
			wall[i-1][j] = wall[i][j] + wall[i][j+1]
		}
	}
}

func main() {
	in, _ := os.Open("11040.in")
	defer in.Close()
	out, _ := os.Create("11040.out")
	defer out.Close()

	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		wall := make([][]int, row)
		for i := range wall {
			wall[i] = make([]int, i+1)
			for j := range wall[i] {
				if i%2 == 0 && j%2 == 0 {
					fmt.Fscanf(in, "%d", &wall[i][j])
				}
			}
		}
		solve(wall)
		for _, w := range wall {
			line := fmt.Sprint(w)
			fmt.Fprintln(out, line[1:len(line)-1])
		}
	}
}
