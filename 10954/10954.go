// UVa 10954 - Add All

package main

import (
	"container/heap"
	"fmt"
	"os"
)

type item struct{ value int }

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].value < pq[j].value }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*item)) }

func (pq *priorityQueue) Pop() interface{} {
	item := (*pq)[(*pq).Len()-1]
	*pq = (*pq)[0 : (*pq).Len()-1]
	return item
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
		pq := make(priorityQueue, n)
		for i := range pq {
			fmt.Fscanf(in, "%d", &num)
			pq[i] = &item{num}
		}
		heap.Init(&pq)
		total := 0
		for pq.Len() > 1 {
			n1 := heap.Pop(&pq).(*item)
			n2 := heap.Pop(&pq).(*item)
			total += n1.value + n2.value
			heap.Push(&pq, &item{n1.value + n2.value})
		}
		fmt.Fprintln(out, total)
	}
}
