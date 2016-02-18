// UVa 575 - Skew Binary

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("575.in")
	defer in.Close()
	out, _ := os.Create("575.out")
	defer out.Close()

	var s string
	var t, pow int
	for {
		if fmt.Fscanf(in, "%s", &s); s == "0" {
			break
		}
		t, pow = 0, 2
		for i := range s {
			t += int(s[len(s)-i-1]-'0') * (pow - 1)
			pow *= 2
		}
		fmt.Fprintln(out, t)
	}
}
