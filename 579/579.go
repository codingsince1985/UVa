// UVa 579 - Clock Hands

package main

import (
	"fmt"
	"os"
)

func hrAngle(h, m int) float64 {
	return 30 * (float64(h) + float64(m)/60)
}

func minAngle(m int) float64 {
	return float64(6 * m)
}

func main() {
	in, _ := os.Open("579.in")
	defer in.Close()
	out, _ := os.Create("579.out")
	defer out.Close()

	var h, m int
	var line string
	for {
		if fmt.Fscanf(in, "%s", &line); line == "0:00" {
			break
		}
		fmt.Sscanf(line[:len(line)-3], "%d", &h)
		fmt.Sscanf(line[len(line)-2:], "%d", &m)
		hour, minute := hrAngle(h, m), minAngle(m)
		diff := max(hour, minute) - min(hour, minute)
		fmt.Fprintf(out, "%.3f\n", min(diff, 360-diff))
	}
}
