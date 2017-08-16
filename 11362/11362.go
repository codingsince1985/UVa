// UVa 11362 - Phone List

package main

import (
	"fmt"
	"os"
)

type trie struct {
	last bool
	next [10]*trie
}

func addToTrie(curr *trie, phone string) bool {
	for i := range phone {
		if curr.last {
			return false
		}
		idx := phone[i] - '0'
		if curr.next[idx] == nil {
			curr.next[idx] = &trie{}
		}
		curr = curr.next[idx]
	}
	if curr.last {
		return false
	}
	curr.last = true
	for _, next := range curr.next {
		if next != nil {
			return false
		}
	}
	return true
}

func solve(phones []string) bool {
	root := trie{}
	for _, phone := range phones {
		if !addToTrie(&root, phone) {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("11362.in")
	defer in.Close()
	out, _ := os.Create("11362.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		phones := make([]string, n)
		for i := range phones {
			fmt.Fscanf(in, "%s", &phones[i])
		}
		if solve(phones) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
