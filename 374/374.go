// UVa 374 - Big Mod

package main

import (
	"fmt"
	"os"
)

func mod(b, p, m int) int {
	if p == 0 {
		return 1;
	}
	if p % 2 == 0 {
		tmp := mod(b, p / 2, m)
		return tmp * tmp % m
	} else {
		return (b % m) * mod(b, p - 1, m) % m
	}
}


func main() {
	in, _ := os.Open("374.in")
	defer in.Close()
	out, _ := os.Create("374.out")
	defer out.Close()

	var b, p, m int
	for {
		_, err := fmt.Fscanf(in, "%d\n%d\n%d\n\n", &b, &p, &m)
		if (err != nil) {
			break;
		}
		fmt.Fprintln(out, mod(b, p, m))
	}
}