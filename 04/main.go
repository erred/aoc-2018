package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type R struct {
	t time.Time
	a string
}

type Rs []R

func (rs Rs) Len() int           { return len(rs) }
func (rs Rs) Less(i, j int) bool { return rs[i].t.Before(rs[j].t) }
func (rs Rs) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	records := []R{}
	for {
		if scanner.Scan() {
			s := scanner.Text()
			if s == "" {
				break
			}
			t, _ := time.Parse("2006-01-02 15:04", s[1:17])
			records = append(records, R{t, s[19:]})
		}
	}

	sort.Sort(Rs(records))

	guards := map[int][]int{}
	gid := -1
	var startsleep time.Time
	for _, r := range records {
		switch r.a {
		case "falls asleep":
			startsleep = r.t
		case "wakes up":
			for i := startsleep.Minute(); i < r.t.Minute(); i++ {
				guards[gid][i]++
			}
		default:
			fmt.Sscanf(r.a, "Guard #%d begins shift", &gid)
			if _, ok := guards[gid]; !ok {
				guards[gid] = make([]int, 60)
			}
		}
	}

	// part 1
	total, idx, maxt := 0, 0, 0
	for g, minutes := range guards {
		sum, id := 0, 0
		for i, m := range minutes {
			sum += m
			if m > maxt {
				maxt, id = m, i
			}
		}
		if sum > total {
			total, idx, gid = sum, id, g
		}
	}
	fmt.Printf("Strategy 1: %v x %v = %v\n", gid, idx, gid*idx)

	// part 2
	maxt = 0
	for g, minutes := range guards {
		for i, m := range minutes {
			if m > maxt {
				maxt, gid, idx = m, g, i
			}
		}
	}
	fmt.Printf("Strategy 2: %vx%v = %v\n", gid, idx, gid*idx)
}
