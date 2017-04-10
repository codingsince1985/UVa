// UVa 902 - Password Search

package main

import (
	"fmt"
	"os"
)

func solve(n int, code string) string {
	var max int
	var pass string
	passMap := make(map[string]int)
	for i := n; i <= len(code); i++ {
		if passMap[code[i-n:i]]++; passMap[code[i-n:i]] > max {
			max = passMap[code[i-n:i]]
			pass = code[i-n : i]
		}
	}
	return pass
}

func main() {
	in, _ := os.Open("902.in")
	defer in.Close()
	out, _ := os.Create("902.out")
	defer out.Close()

	var n int
	var code string
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fscanf(in, "%s", &code)
		fmt.Fprintln(out, solve(n, code))
	}
}
