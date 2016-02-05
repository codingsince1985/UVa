// UVa 10055 - Hashmat the Brave Warrior

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10055.in")
	defer in.Close()
	out, _ := os.Create("10055.out")
	defer out.Close()

	var n1, n2 uint64
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n1, &n2); err != nil {
			break
		}
		fmt.Fprintln(out, n2-n1)
	}
}
