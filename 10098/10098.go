// UVa 10098 - Generating Fast

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	res map[string]bool
	out *os.File
)

func dfs(strs []string, visited []bool, ans []string) {
	if len(ans) == len(strs) {
		s := strings.Join(ans, "")
		if _, ok := res[s]; !ok {
			res[s] = true
			fmt.Fprintln(out, s)
		}
		return
	}
	for i := range strs {
		if !visited[i] {
			visited[i] = true
			dfs(strs, visited, append(ans, strs[i]))
			visited[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("10098.in")
	defer in.Close()
	out, _ = os.Create("10098.out")
	defer out.Close()

	var n int
	var str string
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		fmt.Fscanf(in, "%s", &str)
		strs := make([]string, len(str))
		for i := range str {
			strs[i] = string(str[i])
		}
		sort.Strings(strs)
		visited := make([]bool, len(str))
		var ans []string
		res = make(map[string]bool)
		dfs(strs, visited, ans)
		fmt.Fprintln(out)
		n--
	}
}
