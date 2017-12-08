// UVa 418 - Molecules

package main

import (
	"fmt"
	"os"
)

var (
	chains  = make([]string, 4)
	index   []int
	visited []bool
	area    int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func try() {
	for hul := 1; hul < 11; hul++ {
		for vlu := 1; vlu < 11; vlu++ {
			if chains[index[0]][hul] == chains[index[1]][vlu] {
				for hur := hul + 1; hur < 11; hur++ {
					w := hur - hul
					for vru := 1; vru < 11; vru++ {
						if chains[index[0]][hur] == chains[index[2]][vru] {
							for vld := vlu + 1; vld < 11; vld++ {
								h := vld - vlu
								for hdl := 1; hdl < 11; hdl++ {
									if chains[index[3]][hdl] == chains[index[1]][vld] {
										if hdl+w < 11 && vru+h < 11 {
											if chains[index[3]][hdl+w] == chains[index[2]][vru+h] {
												area = max(area, (w-1)*(h-1))
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func dfs(level int) {
	if level == 4 {
		try()
		return
	}
	for i := 0; i < 4; i++ {
		if !visited[i] {
			visited[i] = true
			index[level] = i
			dfs(level + 1)
			visited[i] = false
		}
	}
}

func solve() int {
	visited = make([]bool, 4)
	index = make([]int, 4)
	area = 0
	dfs(0)
	return area
}

func main() {
	in, _ := os.Open("418.in")
	defer in.Close()
	out, _ := os.Create("418.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%s\n%s\n%s\n%s", &chains[0], &chains[1], &chains[2], &chains[3]); chains[0] == "Q" {
			break
		}
		fmt.Fprintln(out, solve())
	}
}
