// UVa 548 - Tree

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var min, leaf int

type node struct {
	value       int
	left, right *node
}

func split(line string) []int {
	tokens := strings.Split(line, " ")
	nodes := make([]int, len(tokens))
	for i := range nodes {
		nodes[i], _ = strconv.Atoi(tokens[i])
	}
	return nodes
}

func locate(inOrder []int, root int) int {
	for i, node := range inOrder {
		if node == root {
			return i
		}
	}
	return -1
}

func buildTree(inOrder, postOrder []int) *node {
	l := len(postOrder)
	if l == 0 {
		return nil
	}
	v := postOrder[l-1]
	idx := locate(inOrder, v)
	return &node{v, buildTree(inOrder[:idx], postOrder[:idx]), buildTree(inOrder[idx+1:], postOrder[idx:l-1])}
}

func dfs(root *node, sum int) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		newSum := sum + root.value
		switch {
		case newSum < min:
			min = sum + root.value
			leaf = root.value
		case newSum == min && root.value < leaf:
			leaf = root.value
		}
		return
	}
	dfs(root.left, sum+root.value)
	dfs(root.right, sum+root.value)
}

func solve(inOrder, postOrder string) int {
	min = math.MaxInt32
	dfs(buildTree(split(inOrder), split(postOrder)), 0)
	return leaf
}

func main() {
	in, _ := os.Open("548.in")
	defer in.Close()
	out, _ := os.Create("548.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var inOrder string
	for s.Scan() {
		inOrder = s.Text()
		s.Scan()
		fmt.Fprintln(out, solve(inOrder, s.Text()))
	}
}
