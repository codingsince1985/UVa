// UVa 10683 - The decadary watch

package main

import (
	"fmt"
	"os"
	"strconv"
)

const traditionalC = 100 * 60 * 60 * 24

var weights = [...]int{1, 100, 60 * 100, 60 * 60 * 100}

func solve(traditional string) int {
	time, _ := strconv.Atoi(traditional)
	var t int
	for i := 0; time > 0; i++ {
		t += time % 100 * weights[i]
		time /= 100
	}
	return int(float64(t) * 10000000 / traditionalC)
}

func main() {
	in, _ := os.Open("10683.in")
	defer in.Close()
	out, _ := os.Create("10683.out")
	defer out.Close()

	var traditional string
	for {
		if _, err := fmt.Fscanf(in, "%s", &traditional); err != nil {
			break
		}
		fmt.Fprintf(out, "%07d\n", solve(traditional))
	}
}
