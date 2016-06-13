// UVa 155 - All Squares

package main

import (
	"fmt"
	"os"
)

var x, y, cnt int

func testIn(cenX, cenY, k int) {
	if k == 0 {
		return
	}
	if cenX-k <= x && cenX+k >= x && cenY-k <= y && cenY+k >= y {
		cnt++
	}
	testIn(cenX+k, cenY+k, k/2)
	testIn(cenX+k, cenY-k, k/2)
	testIn(cenX-k, cenY+k, k/2)
	testIn(cenX-k, cenY-k, k/2)
}

func main() {
	in, _ := os.Open("155.in")
	defer in.Close()
	out, _ := os.Create("155.out")
	defer out.Close()

	var k int
	for {
		if fmt.Fscanf(in, "%d%d%d", &k, &x, &y); k == 0 {
			break
		}
		cnt = 0
		testIn(1024, 1024, k)
		fmt.Fprintln(out, cnt)
	}
}
