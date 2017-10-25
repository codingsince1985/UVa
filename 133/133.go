// UVa 133 - The Dole Queue

package main

import (
	"fmt"
	"os"
	"strings"
)

func solve(n, k, m int, out *os.File) {
	dole := make([]int, n)
	for i := range dole {
		dole[i] = i + 1
	}
	pk, pm := 0, n-1
	outNum := 0
	var outs []string
	for {
		for i := 1; i < k; i++ {
			for {
				if pk++; pk == n {
					pk = 0
				}
				if dole[pk] != -1 {
					break
				}
			}
		}
		for i := 1; i < m; i++ {
			for {
				if pm--; pm == -1 {
					pm = n - 1
				}
				if dole[pm] != -1 {
					break
				}
			}
		}
		if pk == pm {
			outNum++
			outs = append(outs, fmt.Sprintf("% 3d", dole[pk]))
		} else {
			outNum += 2
			outs = append(outs, fmt.Sprintf("% 3d% 3d", dole[pk], dole[pm]))
		}
		if outNum == len(dole) {
			break
		}
		dole[pk], dole[pm] = -1, -1
		for {
			if pk++; pk == n {
				pk = 0
			}
			if dole[pk] != -1 {
				break
			}
		}
		for {
			if pm--; pm == -1 {
				pm = n - 1
			}
			if dole[pm] != -1 {
				break
			}
		}
	}
	fmt.Fprintln(out, strings.Join(outs, ","))
}

func main() {
	in, _ := os.Open("133.in")
	defer in.Close()
	out, _ := os.Create("133.out")
	defer out.Close()

	var n, k, m int
	for {
		if fmt.Fscanf(in, "%d%d%d", &n, &k, &m); n == 0 {
			break
		}
		solve(n, k, m, out)
	}
}
