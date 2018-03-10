// UVa 186 - Trip Routing

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const max = 101

type (
	path struct {
		route  string
		length int
	}
	step struct{ nameId, pathId int }
)

func spfa(n, start, end int, matrix [][][]path) ([]step, int) {
	pre := make([]step, n)
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[start] = 0
	visited := map[int]bool{start: true}
	for queue := []int{start}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		for to, ps := range matrix[curr] {
			for i, p := range ps {
				if distance[to] > distance[curr]+p.length {
					distance[to] = distance[curr] + p.length
					pre[to] = step{curr, i}
					if !visited[to] {
						queue = append(queue, to)
						visited[to] = true
					}
				}
			}
		}
	}
	var paths []step
	for i := end; i != start; i = pre[i].nameId {
		paths = append([]step{{i, pre[i].pathId}}, paths...)
	}
	return append([]step{{start, -1}}, paths...), distance[end]
}

func solve(tokens []string, cityMap map[string]int, orderMap map[int]string, matrix [][][]path, out io.Writer) {
	fmt.Fprintln(out, "\n\nFrom                 To                   Route      Miles")
	fmt.Fprintln(out, "-------------------- -------------------- ---------- -----")
	paths, distance := spfa(len(cityMap), cityMap[tokens[0]], cityMap[tokens[1]], matrix)
	for i := 0; i < len(paths)-1; i++ {
		p := matrix[paths[i].nameId][paths[i+1].nameId][paths[i+1].pathId]
		fmt.Fprintf(out, "%-20s %-20s %-10s %5d\n", orderMap[paths[i].nameId], orderMap[paths[i+1].nameId], p.route, p.length)
	}
	fmt.Fprintln(out, "                                                     -----")
	fmt.Fprintf(out, "                                          Total       %4d\n", distance)
}

func main() {
	in, _ := os.Open("186.in")
	defer in.Close()
	out, _ := os.Create("186.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	matrix := make([][][]path, max)
	for i := range matrix {
		matrix[i] = make([][]path, max)
	}
	cityMap := make(map[string]int)
	orderMap := make(map[int]string)
	var line string
	var idx, l int
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		tokens := strings.Split(line, ",")
		if _, ok := cityMap[tokens[0]]; !ok {
			cityMap[tokens[0]] = idx
			orderMap[idx] = tokens[0]
			idx++
		}
		if _, ok := cityMap[tokens[1]]; !ok {
			cityMap[tokens[1]] = idx
			orderMap[idx] = tokens[1]
			idx++
		}
		fmt.Sscanf(tokens[3], "%d", &l)
		matrix[cityMap[tokens[0]]][cityMap[tokens[1]]] = append(matrix[cityMap[tokens[0]]][cityMap[tokens[1]]], path{tokens[2], l})
		matrix[cityMap[tokens[1]]][cityMap[tokens[0]]] = append(matrix[cityMap[tokens[1]]][cityMap[tokens[0]]], path{tokens[2], l})
	}
	for s.Scan() {
		if line = s.Text(); line == "" {
			break
		}
		solve(strings.Split(line, ","), cityMap, orderMap, matrix, out)
	}
}
