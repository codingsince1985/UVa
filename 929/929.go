// UVa 929 - Number Maze

package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
)

type (
	node          struct{ y, x, cost int }
	priorityQueue []*node
)

var directions = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*node)) }

func (pq *priorityQueue) Pop() interface{} {
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return item
}

func dijkstra(n, m int, maze [][]int) int {
	visited := make([][]bool, n)
	cost := make([][]int, n)
	for i := range cost {
		visited[i] = make([]bool, m)
		cost[i] = make([]int, m)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}
	cost[0][0], visited[0][0] = maze[0][0], true
	var pq priorityQueue
	for pq.Push(&node{0, 0, cost[0][0]}); pq.Len() > 0; {
		curr := heap.Pop(&pq).(*node)
		for _, direction := range directions {
			if newY, newX := curr.y+direction[0], curr.x+direction[1]; newY >= 0 && newY < n && newX >= 0 && newX < m && !visited[newY][newX] {
				if newCost := curr.cost + maze[newY][newX]; newCost < cost[newY][newX] {
					cost[newY][newX], visited[newY][newX] = newCost, true
					heap.Push(&pq, &node{newY, newX, newCost})
				}
			}
		}
	}
	return cost[n-1][m-1]
}

func main() {
	in, _ := os.Open("929.in")
	defer in.Close()
	out, _ := os.Create("929.out")
	defer out.Close()

	var kase, n, m int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d\n%d", &n, &m)
		maze := make([][]int, n)
		for i := range maze {
			maze[i] = make([]int, m)
			for j := range maze[i] {
				fmt.Fscanf(in, "%d", &maze[i][j])
			}
		}
		fmt.Fprintln(out, dijkstra(n, m, maze))
	}
}
