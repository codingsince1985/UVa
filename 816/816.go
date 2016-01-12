// UVa 816 - Abbott's Revenge

package main

import (
	"fmt"
	"os"
)

type (
	node struct {
		r, c, f int
	}

	qnode struct {
		node
		p []node
	}
)

var maze map[node][]node

func toFace(f byte) int {
	switch f {
	case 'N':
		return 0
	case 'E':
		return 1
	case 'S':
		return 2
	default:
		return 3
	}
}

func realFrom(fm node) node {
	switch fm.f {
	case 0:
		return node{fm.r - 1, fm.c, 0}
	case 1:
		return node{fm.r, fm.c + 1, 1}
	case 2:
		return node{fm.r + 1, fm.c, 2}
	default:
		return node{fm.r, fm.c - 1, 3}
	}
}

func toNode(r, c int, f, d byte) node {
	if f == 'N' && d == 'F' || f == 'E' && d == 'L' || f == 'W' && d == 'R' {
		r--
		return node{r, c, 0}
	}
	if f == 'S' && d == 'F' || f == 'E' && d == 'R' || f == 'W' && d == 'L' {
		r++
		return node{r, c, 2}
	}
	if f == 'E' && d == 'F' || f == 'N' && d == 'R' || f == 'S' && d == 'L' {
		c++
		return node{r, c, 1}
	}
	if f == 'W' && d == 'F' || f == 'N' && d == 'L' || f == 'S' && d == 'R' {
		c--
		return node{r, c, 3}
	}
	return node{}
}

func buildMaze(r, c int, dir []string) {
	var f byte
	for _, v := range dir {
		f = v[0]
		fm := node{r, c, toFace(f)}
		if _, ok := maze[fm]; !ok {
			maze[fm] = make([]node, 0)
		}
		for i := 1; i < len(v); i++ {
			maze[fm] = append(maze[fm], toNode(r, c, f, v[i]))
		}
	}
}

func bfs(fm, to node) []node {
	visited := make(map[node]bool)
	queue := make([]qnode, 0); h, l := 0, 0

	visited[fm] = true
	rfm := realFrom(fm)
	visited[rfm] = true
	from := qnode{rfm, []node{fm, rfm}}
	queue = append(queue, from); l++

	for l != 0 {
		curr := queue[h]; h++; l--
		adjs := maze[curr.node]
		for _, v := range adjs {
			if (v.r == to.r && v.c == to.c) {
				return append(curr.p, to)
			}
			if _, ok := visited[v]; !ok {
				visited[v] = true
				queue = append(queue, qnode{v, append(curr.p, v)}); l++
			}
		}
	}
	return nil
}

func output(out *os.File, n string, p []node) {
	fmt.Fprintln(out, n)
	if p == nil {
		fmt.Fprintln(out, "  No Solution Possible")
	} else {
		for idx, v := range p {
			if idx % 10 == 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprintf(out, " (%d,%d)", v.r, v.c)
			if idx != 0 && idx % 9 == 0 {
				fmt.Fprintln(out)
			}
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in, _ := os.Open("816.in")
	defer in.Close()
	out, _ := os.Create("816.out")
	defer out.Close()

	var n, f, token string
	var dir []string
	var r1, c1, r2, c2 int
	var fm, to node
	for {
		fmt.Fscanf(in, "%s", &n)
		if n == "END" {
			break;
		}
		fmt.Fscanf(in, "%d%d%s%d%d", &r1, &c1, &f, &r2, &c2)
		fm = node{r1, c1, toFace(f[0])}
		to = node{r2, c2, -1}
		maze = make(map[node][]node)
		for {
			if _, err := fmt.Fscanf(in, "%d%d", &r1, &c1); err != nil {
				break
			}
			dir = make([]string, 0)
			for {
				fmt.Fscanf(in, "%s", &token)
				if token == "*" {
					break;
				}
				dir = append(dir, token)
			}
			buildMaze(r1, c1, dir)
		}
		p := bfs(fm, to)
		output(out, n, p)
	}
}
