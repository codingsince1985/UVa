// UVa 10281 - Average Speed

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in, _ := os.Open("10281.in")
	defer in.Close()
	out, _ := os.Create("10281.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var speed, time, hr, min, sec, newSpeed int
	var distance float64
	for s.Scan() {
		line := s.Text()
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%d:%d:%d", &hr, &min, &sec)
		newTime := hr*3600 + min*60 + sec
		newDistance := float64(speed*(newTime-time)) / 3600
		if strings.Contains(line, " ") {
			fmt.Fscanf(r, "%d", &newSpeed)
			distance, time, speed = distance+newDistance, newTime, newSpeed
		} else {
			fmt.Fprintf(out, "%s %.2f km\n", line, distance+newDistance)
		}
	}
}
