// UVa 167 - The Sultan's Successors

package main

import (
	"fmt"
	"os"
)

var (
	max     int
	visited [][]bool
)

func valid(visited [][]bool) bool {
	for i := range visited {
		cnt := 0
		for j := range visited[i] {
			if visited[i][j] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}

	for i := 0; i < len(visited[0]); i++ {
		cnt := 0
		for j := range visited {
			if visited[j][i] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}

	for j := 0; j < 8; j++ {
		i := 0
		cnt := 0
		for d := 0; d < 8; d++ {
			x := i + d
			y := j - d
			if !(x < 0 || x > 7 || y < 0 || y > 7) && visited[x][y] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}
	for i := 1; i < 8; i++ {
		j := 7
		cnt := 0
		for d := 0; d < 8; d++ {
			x := i + d
			y := j - d
			if !(x < 0 || x > 7 || y < 0 || y > 7) && visited[x][y] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}

	for j := 0; j < 8; j++ {
		i := 7
		cnt := 0
		for d := 0; d < 8; d++ {
			x := i - d
			y := j - d
			if !(x < 0 || x > 7 || y < 0 || y > 7) && visited[x][y] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}
	for i := 6; i >= 0; i-- {
		j := 7
		cnt := 0
		for d := 0; d < 8; d++ {
			x := i - d
			y := j - d
			if !(x < 0 || x > 7 || y < 0 || y > 7) && visited[x][y] {
				cnt++
			}
		}
		if cnt > 1 {
			return false
		}
	}

	return true
}

func sum(board [][]int, visited [][]bool) {
	total := 0
	for i, v := range visited {
		for j, vv := range v {
			if vv {
				total += board[i][j]
			}
		}
	}
	if total > max {
		max = total
	}
}

func backtracking(board [][]int, row int) {
	if !valid(visited) {
		return
	}
	if row == 8 {
		sum(board, visited)
		return
	}
	for i := range board[0] {
		visited[row][i] = true
		backtracking(board, row+1)
		visited[row][i] = false
	}
}

func main() {
	in, _ := os.Open("167.in")
	defer in.Close()
	out, _ := os.Create("167.out")
	defer out.Close()

	var k int
	for fmt.Fscanf(in, "%d", &k); k > 0; k-- {
		board := make([][]int, 8)
		for i := range board {
			board[i] = make([]int, 8)
			for j := range board[i] {
				fmt.Fscanf(in, "%d", &board[i][j])
			}
		}
		visited = make([][]bool, 8)
		for i := range visited {
			visited[i] = make([]bool, 8)
		}
		max = 0
		backtracking(board, 0)
		fmt.Fprintln(out, max)
	}
}
