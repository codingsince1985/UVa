// UVa 380 - Call Forwarding

package main

import (
	"fmt"
	"os"
)

const (
	deadEnd  = "9999"
	notFound = ""
)

type (
	forward struct {
		from, to int
		ext      string
	}
	extension struct {
		ext      string
		forwards []forward
	}
)

func getForward(extMap map[string]extension, time int, callExt string) string {
	if ext, ok := extMap[callExt]; ok {
		for _, f := range ext.forwards {
			if time >= f.from && time <= f.to {
				if ext.ext == f.ext {
					return deadEnd
				}
				return f.ext
			}
		}
	}
	return notFound
}

func solve(extMap map[string]extension, time int, callExt string) string {
	for visited := map[string]bool{callExt: true}; ; {
		nextExt := getForward(extMap, time, callExt)
		if nextExt == notFound {
			return callExt
		}
		if visited[nextExt] {
			return deadEnd
		}
		visited[nextExt], callExt = true, nextExt
	}
}

func main() {
	in, _ := os.Open("380.in")
	defer in.Close()
	out, _ := os.Create("380.out")
	defer out.Close()

	fmt.Fprintln(out, "CALL FORWARDING OUTPUT")
	var n, from, duration, time int
	var fromExt, toExt, callExt string
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		fmt.Fprintf(out, "SYSTEM %d\n", kase)
		extMap := make(map[string]extension)
		for {
			if fmt.Fscanf(in, "%s", &fromExt); fromExt == "0000" {
				break
			}
			fmt.Fscanf(in, "%d%d%s", &from, &duration, &toExt)
			if _, ok := extMap[fromExt]; !ok {
				extMap[fromExt] = extension{fromExt, nil}
			}
			extMap[fromExt] = extension{fromExt, append(extMap[fromExt].forwards, forward{from, from + duration, toExt})}
		}
		for {
			if fmt.Fscanf(in, "%d", &time); time == 9000 {
				break
			}
			fmt.Fscanf(in, "%s", &callExt)
			fmt.Fprintf(out, "AT %04d CALL TO %s RINGS %s\n", time, callExt, solve(extMap, time, callExt))
		}
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}
