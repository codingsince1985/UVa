// UVa 424 - Integer Inquiry

package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	in, _ := os.Open("424.in")
	defer in.Close()
	out, _ := os.Create("424.out")
	defer out.Close()

	var total, tmp big.Int
	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "0" {
			break
		}
		tmp.SetString(line, 10)
		total.Add(&total, &tmp)
	}
	fmt.Fprintln(out, &total)
}
