// UVa 10194 - Football (aka Soccer)

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type team struct {
	name                              string
	points                            int
	games, wins, ties, losses         int
	goalDiff, goalScored, goalAgainst int
}

var teamMap map[string]team

func nextLine(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func cmp(t1, t2 team) bool {
	if t1.points != t2.points {
		return t1.points > t2.points
	}
	if t1.wins != t2.wins {
		return t1.wins > t2.wins
	}
	if t1.goalDiff != t2.goalDiff {
		return t1.goalDiff > t2.goalDiff
	}
	if t1.goalScored != t2.goalScored {
		return t1.goalScored > t2.goalScored
	}
	if t1.games != t2.games {
		return t1.games < t2.games
	}
	return t1.name < t2.name
}

func values(m map[string]team) []team {
	var v []team
	for _, value := range m {
		v = append(v, value)
	}
	return v
}

func process(a, b string, s1, s2 int) {
	t1, t2 := teamMap[a], teamMap[b]
	t1.games++
	t2.games++
	t1.goalScored += s1
	t1.goalAgainst += s2
	t2.goalScored += s2
	t2.goalAgainst += s1
	t1.goalDiff = t1.goalScored - t1.goalAgainst
	t2.goalDiff = t2.goalScored - t2.goalAgainst

	switch {
	case s1 > s2:
		t1.wins++
		t1.points += 3
		t2.losses++
	case s2 > s1:
		t1.losses++
		t2.wins++
		t2.points += 3
	default:
		t1.ties++
		t1.points++
		t2.ties++
		t2.points++
	}
	teamMap[a], teamMap[b] = t1, t2
}

func output(out io.Writer, t []team) {
	for i, v := range t {
		fmt.Fprintf(out, "%d) %s %dp, %dg (%d-%d-%d), %dgd (%d-%d)\n",
			i+1, v.name, v.points, v.games, v.wins, v.ties, v.losses, v.goalDiff, v.goalScored, v.goalAgainst)
	}
}

func main() {
	in, _ := os.Open("10194.in")
	defer in.Close()
	out, _ := os.Create("10194.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, tm, gm, s1, s2 int
	for fmt.Sscanf(nextLine(s), "%d", &kase); kase > 0; kase-- {
		teamMap = make(map[string]team)
		fmt.Fprintln(out, nextLine(s))

		for fmt.Sscanf(nextLine(s), "%d", &tm); tm > 0; tm-- {
			tmp := nextLine(s)
			teamMap[tmp] = team{name: tmp}
		}

		for fmt.Sscanf(nextLine(s), "%d", &gm); gm > 0; gm-- {
			tokens := strings.Split(nextLine(s), "#")
			scores := strings.Split(tokens[1], "@")
			fmt.Sscanf(scores[0], "%d", &s1)
			fmt.Sscanf(scores[1], "%d", &s2)
			process(tokens[0], tokens[2], s1, s2)
		}

		v := values(teamMap)
		sort.Slice(v, func(i, j int) bool { return cmp(v[i], v[j]) })
		output(out, v)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
