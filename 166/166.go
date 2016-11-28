// UVa 166 - Making Change

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	coin       = [6]int{5, 10, 20, 50, 100, 200}
	num        [6]int
	value, min int
)

func change(total int) int {
	var num int
	for i := 5; i >= 0; i-- {
		if total >= coin[i] {
			num += total / coin[i]
			total %= coin[i]
			if total == 0 {
				break
			}
		}
	}
	return num
}

func solve(level, sum, pay int) {
	if level == 6 {
		if sum > value {
			pay += change(sum - value)
			if pay < min {
				min = pay
			}
		}
		return
	}
	for i := 0; i <= num[level]; i++ {
		solve(level+1, sum+i*coin[level], pay+i)
	}
}

func main() {
	in, _ := os.Open("166.in")
	defer in.Close()
	out, _ := os.Create("166.out")
	defer out.Close()

	var v float64
	for {
		if fmt.Fscanf(in, "%d%d%d%d%d%d", &num[0], &num[1], &num[2], &num[3], &num[4], &num[5]); num[0] == 0 && num[1] == 0 && num[2] == 0 && num[3] == 0 && num[4] == 0 && num[5] == 0 {
			break
		}
		fmt.Fscanf(in, "%f", &v)
		value = int(v * 100)
		min = math.MaxInt32
		solve(0, 0, 0)
		fmt.Fprintf(out, "%3d\n", min)
	}
}
