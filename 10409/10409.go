// UVa 10409 - Die Game

package main

import (
	"fmt"
	"os"
)

var t, n, w int

func rotate(dir string) {
	switch dir {
	case "north":
		t, n = 7-n, t
	case "south":
		t, n = n, 7-t
	case "east":
		t, w = w, 7-t
	case "west":
		t, w = 7-w, t
	}
}

func main() {
	in, _ := os.Open("10409.in")
	defer in.Close()
	out, _ := os.Create("10409.out")
	defer out.Close()

	var steps int
	var dir string
	for {
		if fmt.Fscanf(in, "%d", &steps); steps == 0 {
			break
		}
		t, n, w = 1, 2, 3
		for ; steps > 0; steps-- {
			fmt.Fscanf(in, "%s", &dir)
			rotate(dir)
		}
		fmt.Fprintln(out, t)
	}
}
