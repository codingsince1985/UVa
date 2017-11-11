// UVa 410 - Station Balance

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func solve(c int, average float64, masses []int) (float64, [][]string) {
	newMasses := make([]int, 2*c)
	copy(newMasses, masses)
	sort.Ints(newMasses)
	var imbalance float64
	chambers := make([][]string, c)
	for i := range chambers {
		if newMasses[i] != 0 {
			chambers[i] = append(chambers[i], strconv.Itoa(newMasses[i]))
		}
		chambers[i] = append(chambers[i], strconv.Itoa(newMasses[2*c-i-1]))
		imbalance += abs(float64(newMasses[i]+newMasses[2*c-i-1]) - average)
	}
	return imbalance, chambers
}

func main() {
	in, _ := os.Open("410.in")
	defer in.Close()
	out, _ := os.Create("410.out")
	defer out.Close()

	var c, s, sum int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d%d", &c, &s); err != nil {
			break
		}
		sum = 0
		masses := make([]int, s)
		for i := range masses {
			fmt.Fscanf(in, "%d", &masses[i])
			sum += masses[i]
		}
		fmt.Fprintf(out, "Set #%d\n", kase)
		imbalance, chambers := solve(c, float64(sum)/float64(c), masses)
		for i, chamber := range chambers {
			fmt.Fprintf(out, " %d: %s\n", i, strings.Join(chamber, " "))
		}
		fmt.Fprintf(out, "IMBALANCE = %.5f\n\n", imbalance)
	}
}
