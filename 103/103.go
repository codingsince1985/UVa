// UVa 103 - Stacking Boxes

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type boxes [][]int

func (b boxes) Len() int { return len(b) }

func (b boxes) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b boxes) Less(i, j int) bool {
	for idx := range b[i] {
		if b[i][idx] != b[j][idx] {
			return b[i][idx] < b[j][idx]
		}
	}
	return false
}

func fit(a, b []int) bool {
	for i := range a {
		if a[i] >= b[i] {
			return false
		}
	}
	return true
}

func lis(box boxes) []int {
	sort.Sort(box)
	size := len(box)
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
			if fit(box[i], box[j]) {
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

	var ret []int
	ret = append(ret, maxIdx)
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

func output(out *os.File, order []int, box, original boxes) {
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
		box := make(boxes, n)
		for i := 0; i < n; i++ {
			box[i] = make([]int, d)
			for j := 0; j < d; j++ {
				fmt.Fscanf(in, "%d", &box[i][j])
			}
			sort.Ints(box[i])
		}
		original := make(boxes, n)
		copy(original, box)
		output(out, lis(box), box, original)
	}
}
