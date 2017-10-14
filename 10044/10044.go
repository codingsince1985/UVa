// UVa 10044 - Erdos Numbers

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct{ idx, length int }

func parseAuthor(line string) []string {
	var authors []string
	for i, tokens := 0, strings.Split(strings.Split(line, ": ")[0], ", "); i < len(tokens)/2; i++ {
		authors = append(authors, tokens[i*2]+", "+tokens[i*2+1])
	}
	return authors
}

func buildAuthorMap(authorsList [][]string) map[string]int {
	authorMap := make(map[string]int)
	idx := 0
	for _, authors := range authorsList {
		for _, author := range authors {
			if _, ok := authorMap[author]; !ok {
				authorMap[author] = idx
				idx++
			}
		}
	}
	return authorMap
}

func buildMatrix(authorMap map[string]int, authorsList [][]string) [][]bool {
	l := len(authorMap)
	matrix := make([][]bool, l)
	for i := range matrix {
		matrix[i] = make([]bool, l)
	}
	for _, authors := range authorsList {
		for _, a1 := range authors {
			for _, a2 := range authors {
				if a1 != a2 {
					matrix[authorMap[a1]][authorMap[a2]] = true
					matrix[authorMap[a2]][authorMap[a1]] = true
				}
			}
		}
	}
	return matrix
}

func bfs(matrix [][]bool, fm, to int) int {
	if fm == to {
		return 0
	}
	visited := map[int]bool{fm: true}
	for queue := []node{{fm, 0}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		for i, isCoauthor := range matrix[curr.idx] {
			if !visited[i] && isCoauthor {
				if i == to {
					return curr.length + 1
				}
				visited[i] = true
				queue = append(queue, node{i, curr.length + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10044.in")
	defer in.Close()
	out, _ := os.Create("10044.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var k, p, n int
	s.Scan()
	fmt.Sscanf(s.Text(), "%d", &k)
	for kase := 1; s.Scan() && kase <= k; kase++ {
		fmt.Fprintf(out, "Scenario %d\n", kase)
		fmt.Sscanf(s.Text(), "%d%d", &p, &n)
		var authorsList [][]string
		for ; p > 0 && s.Scan(); p-- {
			authorsList = append(authorsList, parseAuthor(s.Text()))
		}
		authorMap := buildAuthorMap(authorsList)
		matrix := buildMatrix(authorMap, authorsList)
		to := authorMap["Erdos, P."]
		for ; n > 0 && s.Scan(); n-- {
			author := s.Text()
			fmt.Fprintf(out, "%s ", author)
			if fm, ok := authorMap[author]; !ok {
				fmt.Fprintln(out, "infinity")
			} else {
				if length := bfs(matrix, fm, to); length != -1 {
					fmt.Fprintln(out, length)
				} else {
					fmt.Fprintln(out, "infinity")
				}
			}
		}
	}
}
