// UVa 386 - Perfect Cubes

package main

import (
	"fmt"
	"os"
)

func cubes() map[int]int {
	dict := make(map[int]int)
	for i := 2; i <= 200; i++ {
		dict[i*i*i] = i
	}
	return dict
}

func main() {
	out, _ := os.Create("386.out")
	defer out.Close()

	dict := cubes()
	for a := 2; a <= 200; a++ {
		aCube := a * a * a
		for b := 2; b < a; b++ {
			bCube := b * b * b
			for c := b + 1; c < a; c++ {
				cCube := c * c * c
				dCube := aCube - bCube - cCube
				if dCube < cCube {
					break
				}
				if d, ok := dict[dCube]; ok {
					fmt.Fprintf(out, "Cube = %d, Triple = (%d,%d,%d)\n", a, b, c, d)
				}
			}
		}
	}
}
