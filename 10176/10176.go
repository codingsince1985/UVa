// UVa 10176 - Ocean Deep! - Make it shallow!!

package main

import (
	"fmt"
	"os"
	"strings"
)

func mod(num string) int {
	var n int
	for i := range num {
		n *= 2
		if num[i] == '1' {
			n++
		}
		n %= 131071
	}
	return n
}

func main() {
	in, _ := os.Open("10176.in")
	defer in.Close()
	out, _ := os.Create("10176.out")
	defer out.Close()

	var line, num string
	for {
		if _, err := fmt.Fscanf(in, "%s", &line); err != nil {
			break
		}
		num += line
		if strings.HasSuffix(line, "#") {
			if mod(num[:len(num)-1]) == 0 {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}
			num = ""
		}
	}
}
