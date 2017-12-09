// UVa 10070 - Leap Year or Not Leap Year and ...

package main

import (
	"fmt"
	"os"
)

func leapYear(year int) bool {
	switch {
	case year%4 != 0:
		return false
	case year%400 == 0:
		return true
	default:
		return year%100 != 0
	}
}

func huluculuYear(year int) bool { return year%15 == 0 }

func bulukuluYear(year int) bool { return leapYear(year) && year%55 == 0 }

func main() {
	in, _ := os.Open("10070.in")
	defer in.Close()
	out, _ := os.Create("10070.out")
	defer out.Close()

	var year int
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%d", &year); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		if leap, huluculu, bulukulu := leapYear(year), huluculuYear(year), bulukuluYear(year); !leap && !huluculu && !bulukulu {
			fmt.Fprintln(out, "This is an ordinary year.")
		} else {
			if leap {
				fmt.Fprintln(out, "This is leap year.")
			}
			if huluculu {
				fmt.Fprintln(out, "This is huluculu festival year.")
			}
			if bulukulu {
				fmt.Fprintln(out, "This is bulukulu festival year.")
			}
		}
	}
}
