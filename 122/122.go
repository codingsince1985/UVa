// UVa 122 - Trees on the level

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
			switch p {
			case 'L':
				if curr.l == nil {
					curr.l = &node{-1, nil, nil}
				}
				curr = curr.l
			case 'R':
				if curr.r == nil {
					curr.r = &node{-1, nil, nil}
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
	var ans []string
	for queue := []node{*head}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr.n == -1 {
			return nil
		}
		ans = append(ans, strconv.Itoa(curr.n))
		if curr.l != nil {
			queue = append(queue, *curr.l)
		}
		if curr.r != nil {
			queue = append(queue, *curr.r)
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
	var nodes [][2]string
	for {
		if _, err := fmt.Fscanf(in, "%s", &token); err != nil {
			break
		}
		if len(token) == 2 {
			if head := buildTree(nodes); head == nil {
				fmt.Fprintln(out, "not complete")
			} else {
				ans := bfs(head)
				if ans == nil {
					fmt.Fprintln(out, "not complete")
				} else {
					fmt.Fprintln(out, strings.Join(ans, " "))
				}
			}
			nodes = nil
		} else {
			token = token[1 : len(token)-1]
			tokens := strings.Split(token, ",")
			nodes = append(nodes, [2]string{tokens[0], tokens[1]})
		}
	}
}
