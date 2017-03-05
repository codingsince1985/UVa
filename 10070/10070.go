// UVa 10070 - Leap Year or Not Leap Year and ...

package main

import (
	"fmt"
	"os"
)

func leapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	return year%100 != 0
}

func huluculuYear(year int) bool {
	return year%15 == 0
}

func bulukuluYear(leapYear bool, year int) bool {
	return leapYear && year%55 == 0
}

func main() {
	in, _ := os.Open("10070.in")
	defer in.Close()
	out, _ := os.Create("10070.out")
	defer out.Close()

	var year int
	first := true
	for {
		if _, err := fmt.Fscanf(in, "%d", &year); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}

		leap := leapYear(year)
		huluculu := huluculuYear(year)
		bulukulu := bulukuluYear(leap, year)
		if !leap && !huluculu && !bulukulu {
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
