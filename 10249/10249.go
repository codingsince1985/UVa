// UVa 10249 - The Grand Dinner

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type table struct{ no, size int }

func solve(m int, teams []int, tables []table) [][]string {
	seats := make([][]string, m)
	for i, team := range teams {
		sort.Slice(tables, func(i, j int) bool { return tables[i].size > tables[j].size })
		for j := 0; j < team; j++ {
			if tables[j].size--; tables[j].size < 0 {
				return nil
			}
			seats[i] = append(seats[i], strconv.Itoa(tables[j].no))
		}
	}
	return seats
}

func main() {
	in, _ := os.Open("10249.in")
	defer in.Close()
	out, _ := os.Create("10249.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		teams := make([]int, m)
		for i := range teams {
			fmt.Fscanf(in, "%d", &teams[i])
		}
		tables := make([]table, n)
		for i := range tables {
			fmt.Fscanf(in, "%d", &tables[i].size)
			tables[i].no = i + 1
		}
		if seats := solve(m, teams, tables); seats != nil {
			fmt.Fprintln(out, 1)
			for _, seat := range seats {
				fmt.Fprintln(out, strings.Join(seat, " "))
			}
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}
