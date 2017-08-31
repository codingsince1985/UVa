// UVa 10371 - Time Zones

package main

import (
	"fmt"
	"math"
	"os"
)

var timeZone = map[string]float64{
	"UTC":  0,
	"GMT":  0,
	"BST":  1,
	"IST":  1,
	"WET":  0,
	"WEST": 1,
	"CET":  1,
	"CEST": 2,
	"EET":  2,
	"EEST": 3,
	"MSK":  3,
	"MSD":  4,
	"AST":  -4,
	"ADT":  -3,
	"NST":  -3.5,
	"NDT":  -2.5,
	"EST":  -5,
	"EDT":  -4,
	"CST":  -6,
	"CDT":  -5,
	"MST":  -7,
	"MDT":  -6,
	"PST":  -8,
	"PDT":  -7,
	"HST":  -10,
	"AKST": -9,
	"AKDT": -8,
	"AEST": 10,
	"AEDT": 11,
	"ACST": 9.5,
	"ACDT": 10.5,
	"AWST": 8,
}

func toString(hour, minute int) string {
	switch {
	case hour == 0 && minute == 0:
		return "midnight"
	case hour == 12 && minute == 0:
		return "noon"
	case hour >= 12:
		if hour > 12 {
			hour -= 12
		}
		return fmt.Sprintf("%d:%02d p.m.", hour, minute)
	default:
		if hour == 0 {
			hour = 12
		}
		return fmt.Sprintf("%d:%02d a.m.", hour, minute)
	}
}

func solve(hour, minute int, tz1, tz2 string) string {
	var h, m float64
	if h = timeZone[tz2] - timeZone[tz1]; h < 0 {
		h += 24
	}
	if m = h - math.Floor(h); m != 0 {
		h = math.Floor(h)
		m *= 60
	}
	hour += int(h)
	if minute += int(m); minute >= 60 {
		minute -= 60
		hour++
	}
	return toString(hour%24, minute)
}

func main() {
	in, _ := os.Open("10371.in")
	defer in.Close()
	out, _ := os.Create("10371.out")
	defer out.Close()

	var n, hour, minute int
	var time, tz1, tz2 string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s", &time)
		switch time {
		case "noon":
			hour, minute = 12, 0
		case "midnight":
			hour, minute = 0, 0
		default:
			fmt.Sscanf(time, "%d:%d", &hour, &minute)
			if fmt.Fscanf(in, "%s", &time); time == "p.m." && hour < 12 {
				hour += 12
			}
		}
		fmt.Fscanf(in, "%s%s", &tz1, &tz2)
		fmt.Fprintln(out, solve(hour, minute, tz1, tz2))
	}
}
