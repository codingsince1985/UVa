// UVa 144 - Student Grants

package main

import (
	"fmt"
	"os"
)

type student struct{ id, amount int }

func atm(out *os.File, queue []student, k int) {
	for len(queue) > 0 {
		for i := 1; len(queue) > 0 && i <= k; i++ {
			for amount := i; amount != 0 && len(queue) > 0; queue = queue[1:] {
				current := queue[0]
				if current.amount > amount {
					current.amount -= amount
					amount = 0
					queue = append(queue, current)
				} else {
					fmt.Fprintf(out, "%3d", current.id)
					amount -= current.amount
				}
			}
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("144.in")
	defer in.Close()
	out, _ := os.Create("144.out")
	defer out.Close()

	var n, k int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &k); n == 0 && k == 0 {
			break
		}
		queue := make([]student, n)
		for i := range queue {
			queue[i] = student{i + 1, 40}
		}
		atm(out, queue, k)
	}
}
