// UVa 139 - Telephone Tangles

package main

import (
	"fmt"
	"os"
	"strings"
)

type cost struct {
	code, name string
	price      int
}

func solve(number string, duration int, costs []cost) string {
	if !strings.HasPrefix(number, "0") {
		return fmt.Sprintf("%-16s%-5s%30s%5d%6.2f%7.2f", number, "Local", number, duration, 0.0, 0.0)
	}
	for _, v := range costs {
		if strings.HasPrefix(number, v.code) {
			unit := float64(v.price) / 100.0
			return fmt.Sprintf("%-16s%-25s%10s%5d%6.2f%7.2f", number, v.name, number[len(v.code):], duration, unit, float64(duration)*unit)
		}
	}
	return fmt.Sprintf("%-16s%-25s%10s%5d%6s%7.2f", number, "Unknown", "", duration, "", -1.0)
}

func main() {
	in, _ := os.Open("139.in")
	defer in.Close()
	out, _ := os.Create("139.out")
	defer out.Close()

	var code, token, number string
	var costs []cost
	var duration, price int
	for {
		if fmt.Fscanf(in, "%s", &code); code == "000000" {
			break
		}
		fmt.Fscanf(in, "%s", &token)
		tokens := strings.Split(token, "$")
		fmt.Sscanf(tokens[1], "%d", &price)
		costs = append(costs, cost{code, tokens[0], price})
	}
	for {
		if fmt.Fscanf(in, "%s", &number); number == "#" {
			break
		}
		fmt.Fscanf(in, "%d", &duration)
		fmt.Fprintln(out, solve(number, duration, costs))
	}
}
