// UVa 114 - Simulation Wizardry

package main

import (
	"fmt"
	"os"
)

type (
	point  struct{ x, y int }
	bumper struct{ value, cost int }
	ball   struct {
		point
		direction, lifetime int
	}
)

var (
	m, n, costWall int
	bumperMap      = make(map[point]bumper)
	directionMap   = map[int][2]int{0: {1, 0}, 1: {0, 1}, 2: {-1, 0}, 3: {0, -1}}
)

func play(b ball) int {
	var points int
	for b.lifetime > 1 {
		b.lifetime--
		direction := directionMap[b.direction]
		p := point{b.point.x + direction[0], b.point.y + direction[1]}
		if p.x == 1 || p.y == 1 || p.x == m || p.y == n {
			b.lifetime -= costWall
			b.direction = (b.direction + 3) % 4
		} else {
			if bp, ok := bumperMap[p]; ok {
				b.lifetime -= bp.cost
				points += bp.value
				b.direction = (b.direction + 3) % 4
			} else {
				b.point = p
			}
		}
	}
	return points
}

func main() {
	in, _ := os.Open("114.in")
	defer in.Close()
	out, _ := os.Create("114.out")
	defer out.Close()

	var p, total int
	var pt point
	var bp bumper
	var b ball
	fmt.Fscanf(in, "%d%d", &m, &n)
	fmt.Fscanf(in, "%d", &costWall)
	for fmt.Fscanf(in, "%d", &p); p > 0; p-- {
		fmt.Fscanf(in, "%d%d%d%d", &pt.x, &pt.y, &bp.value, &bp.cost)
		bumperMap[pt] = bp
	}
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d%d", &b.point.x, &b.point.y, &b.direction, &b.lifetime); err != nil {
			break
		}
		points := play(b)
		fmt.Fprintln(out, points)
		total += points
	}
	fmt.Fprintln(out, total)
}
