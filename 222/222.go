// UVa 222 - Budget Travel

package main

import (
	"fmt"
	"math"
	"os"
)

type station struct{ distance, price float64 }

var (
	num                                 int
	stations                            []station
	destination, capacity, mpg, minCost float64
)

func backtracking(cost, location, remaining float64, curr int) {
	if curr+1 == num {
		if remaining-(destination-location)/mpg >= 0 && cost < minCost {
			minCost = cost
		}
		return
	}
	for i := curr + 1; i < num; i++ {
		newRemaining := remaining - (stations[i].distance-location)/mpg
		if remaining >= 0 {
			backtracking(cost, stations[i].distance, newRemaining, i)
		}
		if remaining < 0 || remaining >= 0 && newRemaining <= capacity/2 {
			backtracking(cost+(capacity-newRemaining)*stations[i].price*0.01+2, stations[i].distance, capacity, i)
		}
	}
}

func main() {
	in, _ := os.Open("222.in")
	defer in.Close()
	out, _ := os.Create("222.out")
	defer out.Close()

	var initCost float64
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%f", &destination); destination < 0 {
			break
		}
		fmt.Fscanf(in, "%f%f%f%d", &capacity, &mpg, &initCost, &num)
		stations = make([]station, num)
		for i := range stations {
			fmt.Fscanf(in, "%f%f", &stations[i].distance, &stations[i].price)
		}
		minCost = math.MaxFloat64
		backtracking(initCost, 0, capacity, -1)
		fmt.Fprintf(out, "Data Set #%d\n", kase)
		fmt.Fprintf(out, "minimum cost = $%.2f\n", minCost)
	}
}
