// UVa 644 - Immediate Decodability

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func decodable(code []string) bool {
	sort.Strings(code)
	for i := 0; i < len(code)-1; i++ {
		if strings.HasPrefix(code[i+1], code[i]) {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("644.in")
	defer in.Close()
	out, _ := os.Create("644.out")
	defer out.Close()

	var line string
here:
	for kase := 1; ; kase++ {
		var code []string
		for {
			if _, err := fmt.Fscanf(in, "%s", &line); err != nil || line == "9" {
				if err != nil {
					break here
				}
				break
			}
			code = append(code, line)
		}
		if decodable(code) {
			fmt.Fprintf(out, "Set %d is immediately decodable\n", kase)
		} else {
			fmt.Fprintf(out, "Set %d is not immediately decodable\n", kase)
		}
	}
}
