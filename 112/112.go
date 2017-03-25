// UVa 112 - Tree Summing

package main

import (
	"fmt"
	"os"
	"strconv"
)

type node struct {
	n    int
	l, r *node
}

func buildTree(line string) node {
	if len(line) == 0 {
		return node{}
	}
	var l, r, count int
	for idx, v := range line {
		if l == 0 && v == '(' {
			l = idx
		}
		switch v {
		case '(':
			count++
		case ')':
			count--
		}
		if r == 0 && l != 0 && count == 0 {
			r = idx
		}
	}
	left := buildTree(line[l+1 : r])
	right := buildTree(line[r+2 : len(line)-1])
	value, _ := strconv.Atoi(line[:l])
	tree := node{value, &left, &right}
	return tree
}

func dfs(sum int, node *node, ssf int) bool {
	if node == nil {
		return false
	}

	if ssf+node.n > sum {
		return false
	}

	if ssf+node.n == sum && node.l == nil && node.r == nil {
		return true
	}

	if dfs(sum, node.l, ssf+node.n) || dfs(sum, node.r, ssf+node.n) {
		return true
	}
	return false
}

func main() {
	in, _ := os.Open("112.in")
	defer in.Close()
	out, _ := os.Create("112.out")
	defer out.Close()

	var sum, count int
	var c byte
	for {
		if _, err := fmt.Fscanf(in, "%d (", &sum); err != nil {
			break
		}
		count++
		chars := []byte{'('}
		for count > 0 {
			fmt.Fscanf(in, "%c", &c)
			if c >= '0' && c <= '9' || c == '(' || c == ')' {
				chars = append(chars, c)
				switch c {
				case '(':
					count++
				case ')':
					count--
				}
			}
		}
		fmt.Fscanln(in)
		tree := buildTree(string(chars[1 : len(chars)-1]))
		if dfs(sum, &tree, 0) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
