// UVa 332 - Rational Numbers from Repeating Fractions

package main

import (
	"fmt"
	"math"
	"os"
)

func numerator(k, j int, x string) int {
	var v1, v2 float64
	fmt.Sscanf(x, "%f", &v1)
	fmt.Sscanf(x[:len(x)-j], "%f", &v2)
	return int(v1*math.Pow(10, float64(k+j)) - v2*math.Pow(10, float64(k)))
}

func denominator(k, j int) int { return int(math.Pow(10, float64(k+j)) - math.Pow(10, float64(k))) }

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	in, _ := os.Open("332.in")
	defer in.Close()
	out, _ := os.Create("332.out")
	defer out.Close()

	var j int
	var x string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &j); j == -1 {
			break
		}
		fmt.Fscanf(in, "%s", &x)
		k := len(x) - 2 - j
		n, d := numerator(k, j, x), denominator(k, j)
		factor := gcd(n, d)
		fmt.Fprintf(out, "Case %d: %d/%d\n", kase, n/factor, d/factor)
	}
}
