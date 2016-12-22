// UVa 10491 - Cows and Cars

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10491.in")
	defer in.Close()
	out, _ := os.Create("10491.out")
	defer out.Close()

	var cows, cars, show float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f", &cows, &cars, &show); err != nil {
			break
		}
		p1 := cows / (cows + cars) * cars / (cows + cars - show - 1)
		p2 := cars / (cows + cars) * (cars - 1) / (cows + cars - show - 1)
		fmt.Fprintf(out, "%.5f\n", p1+p2)
	}
}
