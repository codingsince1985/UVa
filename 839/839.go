// UVa 839 - Not so Mobile

package main

import (
	"fmt"
	"os"
)

type mobile struct {
	wl, dl, wr, dr int
	left, right    *mobile
}

func readLine(in *os.File) *mobile {
	var m mobile
	fmt.Fscanf(in, "%d%d%d%d", &m.wl, &m.dl, &m.wr, &m.dr)
	if m.wl == 0 {
		m.left = readLine(in)
	}
	if m.wr == 0 {
		m.right = readLine(in)
	}
	return &m
}

func solve(root *mobile) (bool, int) {
	var ok bool
	if root.left != nil {
		if ok, root.wl = solve(root.left); !ok {
			return false, -1
		}
	}
	if root.right != nil {
		if ok, root.wr = solve(root.right); !ok {
			return false, -1
		}
	}
	return root.wl*root.dl == root.wr*root.dr, root.wl + root.wr
}

func main() {
	in, _ := os.Open("839.in")
	defer in.Close()
	out, _ := os.Create("839.out")
	defer out.Close()

	var kase int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		root := readLine(in)
		if ok, _ := solve(root); ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
