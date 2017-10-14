// UVa 622 - Grammar Evaluation

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(token string) (bool, int64) {
	if n, err := strconv.Atoi(token); err == nil {
		return true, int64(n)
	}
	return false, -1
}

func isFactor(token string) (bool, int64) {
	if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
		return isExpression(token[1 : len(token)-1])
	}
	return parse(token)
}

func indices(token string, sep byte) []int {
	var indexes []int
	for i := range token {
		if token[i] == sep {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func isComponent(token string) (bool, int64) {
	if indexes := indices(token, '*'); len(indexes) > 0 {
		for _, index := range indexes {
			if ok, v1 := isFactor(token[:index]); ok {
				if ok, v2 := isComponent(token[index+1:]); ok {
					return true, v1 * v2
				}
			}
		}
		return false, -1
	}
	return isFactor(token)
}

func isExpression(token string) (bool, int64) {
	if indexes := indices(token, '+'); len(indexes) > 0 {
		for _, index := range indexes {
			if ok, v1 := isComponent(token[:index]); ok {
				if ok, v2 := isExpression(token[index+1:]); ok {
					return true, v1 + v2
				}
			}
		}
		return false, -1
	}
	return isComponent(token)
}

func main() {
	in, _ := os.Open("622.in")
	defer in.Close()
	out, _ := os.Create("622.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%s", &line)
		if ok, v := isExpression(line); ok {
			fmt.Fprintln(out, v)
		} else {
			fmt.Fprintln(out, "ERROR")
		}
	}
}
