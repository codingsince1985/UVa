// UVa 10282 - Babelfish

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10282.in")
	defer in.Close()
	out, _ := os.Create("10282.out")
	defer out.Close()

	var s1, s2 string
	dict := make(map[string]string)
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &s1, &s2); err != nil {
			break
		}
		dict[s2] = s1
	}

	for {
		if _, err := fmt.Fscanf(in, "%s", &s2); err != nil {
			break
		}
		if v := dict[s2]; v != "" {
			fmt.Fprintln(out, v)
		} else {
			fmt.Fprintln(out, "eh")
		}
	}
}
