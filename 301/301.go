// UVa 301 - Transportation

package main

import (
	"fmt"
	"os"
)

var (
	max, capacity, orderNumber int
	orders                     []order
	capacities                 []int
)

type order struct{ starting, destination, passengers int }

func backtracking(level, earning int) {
	if level == orderNumber {
		if earning > max {
			max = earning
		}
		return
	}
	backtracking(level+1, earning)
	order := orders[level]
	for j := order.starting; j < order.destination; j++ {
		if capacities[j]+order.passengers > capacity {
			return
		}
	}
	for j := order.starting; j < order.destination; j++ {
		capacities[j] += order.passengers
	}
	backtracking(level+1, earning+(order.destination-order.starting)*order.passengers)
	for j := order.starting; j < order.destination; j++ {
		capacities[j] -= order.passengers
	}
}

func main() {
	in, _ := os.Open("301.in")
	defer in.Close()
	out, _ := os.Create("301.out")
	defer out.Close()

	var stations int
	for {
		if fmt.Fscanf(in, "%d%d%d", &capacity, &stations, &orderNumber); capacity == 0 && stations == 0 && orderNumber == 0 {
			break
		}
		orders = make([]order, orderNumber)
		for i := range orders {
			fmt.Fscanf(in, "%d%d%d", &orders[i].starting, &orders[i].destination, &orders[i].passengers)
		}
		max = 0
		capacities = make([]int, stations+1)
		backtracking(0, 0)
		fmt.Fprintln(out, max)
	}
}
