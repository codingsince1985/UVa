// UVa 11995 - I Can Guess the Data Structure!

package main

import (
	"container/heap"
	"fmt"
	"os"
)

type (
	operation     struct{ op, x int }
	node          struct{ x int }
	priorityQueue []*node
)

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].x > pq[j].x }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*node)) }

func (pq *priorityQueue) Pop() interface{} {
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return item
}

func isQueue(operations []operation) bool {
	var queue []int
	for _, o := range operations {
		if o.op == 1 {
			queue = append(queue, o.x)
		} else {
			if len(queue) == 0 {
				return false
			}
			if queue[0] != o.x {
				return false
			}
			queue = queue[1:]
		}
	}
	return true
}

func isStack(operations []operation) bool {
	var stack []int
	for _, o := range operations {
		if o.op == 1 {
			stack = append([]int{o.x}, stack...)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[0] != o.x {
				return false
			}
			stack = stack[1:]
		}
	}
	return true
}

func isPriorityQueue(operations []operation) bool {
	var pq priorityQueue
	for _, o := range operations {
		if o.op == 1 {
			heap.Push(&pq, &node{o.x})
		} else {
			if pq.Len() == 0 {
				return false
			}
			if curr := heap.Pop(&pq).(*node); curr.x != o.x {
				return false
			}
		}
	}
	return true
}

func solve(operations []operation) string {
	switch q, s, pq := isQueue(operations), isStack(operations), isPriorityQueue(operations); {
	case q && !s && !pq:
		return "queue"
	case !q && s && !pq:
		return "stack"
	case !q && !s && pq:
		return "priority queue"
	case !q && !s && !pq:
		return "impossible"
	default:
		return "not sure"
	}
}

func main() {
	in, _ := os.Open("11995.in")
	defer in.Close()
	out, _ := os.Create("11995.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		operations := make([]operation, n)
		for i := range operations {
			fmt.Fscanf(in, "%d%d", &operations[i].op, &operations[i].x)
		}
		fmt.Fprintln(out, solve(operations))
	}
}
