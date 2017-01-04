// UVa 10591 - Happy Number

package main

import (
	"fmt"
	"os"
)

func happy(num int) bool {
	visited := make(map[int]bool)
	visited[num] = true
	for {
		sum := 0
		for num > 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}
		switch {
		case sum == 1:
			return true
		case visited[sum]:
			return false
		}
		visited[sum] = true
		num = sum
	}
}

func main() {
	in, _ := os.Open("10591.in")
	defer in.Close()
	out, _ := os.Create("10591.out")
	defer out.Close()

	var n, num int
	fmt.Fscanf(in, "%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d", &num)
		if fmt.Fprintf(out, "Case #%d: %d is ", i, num); happy(num) {
			fmt.Fprintln(out, "a Happy number.")
		} else {
			fmt.Fprintln(out, "an Unhappy number.")
		}
	}
}
