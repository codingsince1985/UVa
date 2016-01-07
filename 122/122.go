// UVa 122 - Trees on the level

package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type node struct {
	n    int
	l, r *node
}

func buildTree(nodes [][2]string) *node {
	head := node{-1, nil, nil}
	for _, v := range nodes {
		curr := &head
		for _, p := range v[1] {
			if p == 'L' {
				if curr.l == nil {
					node := node{-1, nil, nil}
					curr.l = &node
				}
				curr = curr.l
			} else if p == 'R' {
				if curr.r == nil {
					node := node{-1, nil, nil}
					curr.r = &node
				}
				curr = curr.r
			}
		}
		if curr.n != -1 {
			return nil
		}
		curr.n, _ = strconv.Atoi(v[0])
	}
	return &head
}

func bfs(head *node) []string {
	ans := make([]string, 0)
	queue := make([]node, 0)
	h, l := 0, 0

	queue = append(queue, *head); l++
	for l != 0 {
		curr := queue[h]; l--; h++
		if curr.n == -1 {
			return nil
		}
		ans = append(ans, strconv.Itoa(curr.n))
		if curr.l != nil {
			queue = append(queue, *curr.l); l++
		}
		if curr.r != nil {
			queue = append(queue, *curr.r); l++
		}
	}
	return ans
}

func main() {
	in, _ := os.Open("122.in")
	defer in.Close()
	out, _ := os.Create("122.out")
	defer out.Close()

	var token string
	nodes := make([][2]string, 0)
	for {
		if _, err := fmt.Fscanf(in, "%s", &token); err != nil {
			break
		}
		if len(token) == 2 {
			head := buildTree(nodes)
			if head == nil {
				fmt.Fprintln(out, "not complete")
			} else {
				ans := bfs(head)
				if ans == nil {
					fmt.Fprintln(out, "not complete")
				} else {
					fmt.Fprintln(out, strings.Join(ans, " "))
				}
			}
			nodes = make([][2]string, 0)
		} else {
			token = token[1:len(token) - 1]
			tokens := strings.Split(token, ",")
			nodes = append(nodes, [2]string{tokens[0], tokens[1]})
		}
	}
}
