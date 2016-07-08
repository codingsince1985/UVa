// UVa 10954 - Add All

package main

import (
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}

func main() {
	in, _ := os.Open("10954.in")
	defer in.Close()
	out, _ := os.Create("10954.out")
	defer out.Close()

	var n, num int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		pq := make(PriorityQueue, n)
		for i := range pq {
			fmt.Fscanf(in, "%d", &num)
			pq[i] = &Item{num, i}
		}
		heap.Init(&pq)
		total := 0
		for pq.Len() > 1 {
			n1 := heap.Pop(&pq).(*Item)
			n2 := heap.Pop(&pq).(*Item)
			total += n1.value + n2.value
			heap.Push(&pq, &Item{n1.value + n2.value, pq.Len()})
		}
		fmt.Fprintln(out, total)
	}
}
