// UVa 103 - Stacking Boxes

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func fit(a, b []int) bool {
	for i := range a {
		if a[i] >= b[i] {
			return false
		}
	}
	return true
}

func lis(boxes [][]int) []int {
	sort.Slice(boxes, func(i, j int) bool {
		for idx := range boxes[i] {
			if boxes[i][idx] != boxes[j][idx] {
				return boxes[i][idx] < boxes[j][idx]
			}
		}
		return false
	})
	size := len(boxes)
	length := make([]int, size)
	for i := range length {
		length[i] = 1
	}
	pre := make([]int, size)
	for i := range pre {
		pre[i] = -1
	}
	max, maxIdx := 0, 0
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if fit(boxes[i], boxes[j]) {
				if length[i]+1 > length[j] {
					length[j] = length[i] + 1
					if length[j] > max {
						max = length[j]
						maxIdx = j
					}
					pre[j] = i
				}
			}
		}
	}

	ret := []int{maxIdx}
	for pre[maxIdx] != -1 {
		maxIdx = pre[maxIdx]
		ret = append([]int{maxIdx}, ret...)
	}
	return ret
}

func isSame(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func output(out *os.File, order []int, box, original [][]int) {
	var ret []string
	for _, v := range order {
		for j := range original {
			if isSame(box[v], original[j]) {
				ret = append(ret, strconv.Itoa(j+1))
				break
			}
		}
	}
	fmt.Fprintln(out, len(ret))
	fmt.Fprintln(out, strings.Join(ret, " "))
}

func main() {
	in, _ := os.Open("103.in")
	defer in.Close()
	out, _ := os.Create("103.out")
	defer out.Close()

	var n, d int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &d); err != nil {
			break
		}
		boxes := make([][]int, n)
		for i := 0; i < n; i++ {
			boxes[i] = make([]int, d)
			for j := 0; j < d; j++ {
				fmt.Fscanf(in, "%d", &boxes[i][j])
			}
			sort.Ints(boxes[i])
		}
		original := make([][]int, n)
		copy(original, boxes)
		output(out, lis(boxes), boxes, original)
	}
}
