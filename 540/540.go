// UVa 540 - Team Queue

package main

import (
	"fmt"
	"os"
)

type node struct {
	element int
	next    *node
}

var (
	elementMap map[int]int
	in, out    *os.File
	head, tail *node
)

func enqueue(element int) {
	if head == nil {
		head = &node{element, nil}
		tail = head
		return
	}
	for curr := head; curr != nil; curr = curr.next {
		if elementMap[element] == elementMap[curr.element] {
			for ; curr.next != nil && elementMap[curr.next.element] == elementMap[element]; curr = curr.next {
			}
			if curr.next == nil {
				break
			}
			next := curr.next
			curr.next = &node{element, next}
			return
		}
	}
	tail.next = &node{element, nil}
	tail = tail.next
}

func solve() {
	var command string
	var element int
	for {
		if fmt.Fscanf(in, "%s", &command); command == "STOP" {
			break
		}
		switch command {
		case "ENQUEUE":
			fmt.Fscanf(in, "%d", &element)
			enqueue(element)
		case "DEQUEUE":
			fmt.Fprintln(out, (*head).element)
			head = head.next
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ = os.Open("540.in")
	defer in.Close()
	out, _ = os.Create("540.out")
	defer out.Close()

	var t, n, element int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &t); t == 0 {
			break
		}
		elementMap = make(map[int]int)
		for ; t > 0; t-- {
			for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
				fmt.Fscanf(in, "%d", &element)
				elementMap[element] = t
			}
		}
		fmt.Fprintf(out, "Scenario #%d\n", kase)
		solve()
	}
}
