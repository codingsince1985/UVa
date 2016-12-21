// UVa 437 - The Tower of Babylon

package main

import (
	"fmt"
	"os"
	"sort"
)

type (
	block  struct{ x, y, z int }
	blocks []block
)

func (b blocks) Len() int { return len(b) }

func (b blocks) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b blocks) Less(i, j int) bool {
	if b[i].x == b[j].x {
		return b[i].y < b[j].y
	}
	return b[i].x < b[j].x
}

func lis(b blocks) int {
	sort.Sort(b)
	h, max := make([]int, len(b)), 0
	for i, block := range b {
		h[i] = block.z
		for j := 0; j < i; j++ {
			if block.x > b[j].x && block.y > b[j].y && h[j]+block.z > h[i] {
				h[i] = h[j] + block.z
				if h[i] > max {
					max = h[i]
				}
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("437.in")
	defer in.Close()
	out, _ := os.Create("437.out")
	defer out.Close()

	var n, x, y, z, kase int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		var b blocks
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d%d%d", &x, &y, &z)
			b = append(b, block{x, y, z})
			b = append(b, block{x, z, y})
			b = append(b, block{y, x, z})
			b = append(b, block{y, z, x})
			b = append(b, block{z, x, y})
			b = append(b, block{z, y, x})
		}
		kase++
		fmt.Fprintf(out, "Case %d: maximum height = %d\n", kase, lis(b))
	}
}
