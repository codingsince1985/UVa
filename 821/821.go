// UVa 821 - Page Hopping

package main

import (
	"fmt"
	"math"
	"os"
)

func getPageMap(links [][2]int) map[int]int {
	pageMap := make(map[int]int)
	idx := 0
	for _, link := range links {
		if _, ok := pageMap[link[0]]; !ok {
			pageMap[link[0]] = idx
			idx++
		}
		if _, ok := pageMap[link[1]]; !ok {
			pageMap[link[1]] = idx
			idx++
		}
	}
	return pageMap
}

func sum(distance [][]int) float64 {
	var total float64
	for _, row := range distance {
		for _, cell := range row {
			if cell != math.MaxInt32 {
				total += float64(cell)
			}
		}
	}
	return total
}

func initialize(links [][2]int) [][]int {
	pageMap := getPageMap(links)
	n := len(pageMap)
	distance := make([][]int, n)
	for i := range distance {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	for _, link := range links {
		distance[pageMap[link[0]]][pageMap[link[1]]] = 1
	}
	return distance
}

func floydWarshall(links [][2]int) float64 {
	distance := initialize(links)
	for k := range distance {
		for i := range distance {
			for j := range distance {
				if i != j && distance[i][k] != math.MaxInt32 && distance[k][j] != math.MaxInt32 &&
					distance[i][j] > distance[i][k]+distance[k][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}
	total := sum(distance)
	n := len(distance)
	return total / float64(n*(n-1))
}

func main() {
	in, _ := os.Open("821.in")
	defer in.Close()
	out, _ := os.Create("821.out")
	defer out.Close()

	var p1, p2 int
	for kase := 1; ; kase++ {
		var links [][2]int
		for {
			if fmt.Fscanf(in, "%d%d", &p1, &p2); p1 == 0 && p2 == 0 {
				break
			}
			links = append(links, [2]int{p1, p2})
		}
		if len(links) == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: average length between pages = %.3f clicks\n", kase, floydWarshall(links))
	}
}
