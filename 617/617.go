// UVa 617 - Nonstop Travel

package main

import (
	"fmt"
	"io"
	"os"
)

type light struct {
	distance float64
	g, y, r  int
}

func solve(lights []light) []int {
	var speeds []int
here:
	for speed := 30; speed <= 60; speed++ {
		for _, l := range lights {
			if int(l.distance/float64(speed)*3600)%(l.g+l.y+l.r) > l.g+l.y {
				continue here
			}
		}
		speeds = append(speeds, speed)
	}
	return speeds
}

func output(out io.Writer, speeds []int) {
	if len(speeds) == 0 {
		fmt.Fprintln(out, "No acceptable speeds.")
		return
	}
	var rangeStarted bool
	for i, speed := range speeds {
		if i == 0 {
			fmt.Fprint(out, speed)
			continue
		}
		if speed == speeds[i-1]+1 {
			if !rangeStarted {
				rangeStarted = true
				fmt.Fprint(out, "-")
			}
			if i == len(speeds)-1 {
				fmt.Fprint(out, speed)
			}
		} else {
			if rangeStarted {
				rangeStarted = false
				fmt.Fprintf(out, "%d", speeds[i-1])
			}
			fmt.Fprintf(out, ", %d", speed)
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("617.in")
	defer in.Close()
	out, _ := os.Create("617.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == -1 {
			break
		}
		fmt.Fprintf(out, "Case %d: ", kase)
		lights := make([]light, n)
		for i := range lights {
			fmt.Fscanf(in, "%f%d%d%d", &lights[i].distance, &lights[i].g, &lights[i].y, &lights[i].r)
		}
		output(out, solve(lights))
		fmt.Fscanln(in)
	}
}
