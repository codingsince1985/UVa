// UVa 300 - Maya Calendar

package main

import (
	"fmt"
	"os"
)

var (
	haabMonths = map[string]int{
		"pop":    0 * 20,
		"no":     1 * 20,
		"zip":    2 * 20,
		"zotz":   3 * 20,
		"tzec":   4 * 20,
		"xul":    5 * 20,
		"yoxkin": 6 * 20,
		"mol":    7 * 20,
		"chen":   8 * 20,
		"yax":    9 * 20,
		"zac":    10 * 20,
		"ceh":    11 * 20,
		"mac":    12 * 20,
		"kankin": 13 * 20,
		"muan":   14 * 20,
		"pax":    15 * 20,
		"koyab":  16 * 20,
		"cumhu":  17 * 20,
		"uayet":  18 * 20,
	}
	tzolkinMonths = []string{
		"imix",
		"ik",
		"akbal",
		"kan",
		"chicchan",
		"cimi",
		"manik",
		"lamat",
		"muluk",
		"ok",
		"chuen",
		"eb",
		"ben",
		"ix",
		"mem",
		"cib",
		"caban",
		"eznab",
		"canac",
		"ahau",
	}
)

func main() {
	in, _ := os.Open("300.in")
	defer in.Close()
	out, _ := os.Create("300.out")
	defer out.Close()

	var n, d, y int
	var m string
	fmt.Fscanf(in, "%d", &n)
	fmt.Fprintln(out, n)
	for ; n > 0; n-- {
		fmt.Fscanf(in, "%d.%s%d", &d, &m, &y)
		days := 365*y + haabMonths[m] + d + 1
		year := days / (13 * 20)
		dayInYear := days % (13 * 20)
		number := (dayInYear-1)%13 + 1
		name := tzolkinMonths[(dayInYear-1)%20]
		fmt.Fprintln(out, number, name, year)
	}
}
