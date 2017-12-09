// UVa 602 - What Day Is It?

package main

import (
	"fmt"
	"os"
)

var (
	months = []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	days       = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	dayInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

func newLeapYear(year int) bool {
	switch {
	case year%4 != 0:
		return false
	case year%400 == 0:
		return true
	default:
		return year%100 != 0
	}
}

func oldLeapYear(year int) bool { return year%4 == 0 }

func leapYear(year int) bool {
	if year > 1752 {
		return newLeapYear(year)
	}
	return oldLeapYear(year)
}

func daysInMonth(y, m int) int {
	if m == 2 {
		if leapYear(y) {
			return 29
		}
		return 28
	}
	return dayInMonth[m-1]
}

func newDate(y, m, d int) bool {
	return y > 1752 || y == 1752 && m > 9 || y == 1752 && m == 9 && d >= 14
}

func solve(y, m, d int) string {
	if m < 1 || m > 12 || d > daysInMonth(y, m) || y == 1752 && m == 9 && 3 <= d && d <= 13 {
		return fmt.Sprintf("%d/%d/%d is an invalid date.", m, d, y)
	}
	count := 5
	for i := 1; i < y; i++ {
		count += 365
		if leapYear(i) {
			count++
		}
	}
	for i := 1; i < m; i++ {
		count += daysInMonth(y, i)
	}
	count += d
	if newDate(y, m, d) {
		count -= 11
	}
	return fmt.Sprintf("%s %d, %d is a %s", months[m-1], d, y, days[count%7])
}

func main() {
	in, _ := os.Open("602.in")
	defer in.Close()
	out, _ := os.Create("602.out")
	defer out.Close()

	var y, m, d int
	for {
		if fmt.Fscanf(in, "%d%d%d", &m, &d, &y); m == 0 && d == 0 && y == 0 {
			break
		}
		fmt.Fprintln(out, solve(y, m, d))
	}
}
