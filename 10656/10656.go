// UVa 10656 - Maximum Sum (II)

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("10656.in")
	defer in.Close()
	out, _ := os.Create("10656.out")
	defer out.Close()

	var n, tmp int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var sub []string
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d", &tmp)
			if tmp > 0 {
				sub = append(sub, strconv.Itoa(tmp))
			}
		}
		if len(sub) == 0 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, strings.Join(sub, " "))
		}
	}
}
