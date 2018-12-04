package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type R struct {
	t time.Time
	a string
}

type Rs []R

func (rs Rs) Len() int {
	return len(rs)
}

func (rs Rs) Less(i, j int) bool {
	return rs[i].t.Before(rs[j].t)
}
func (rs Rs) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func main() {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	records := []R{}
	for {
		r := R{}
		if scanner.Scan() {
			s := scanner.Text()
			if s == "" {
				break
			}
			r.t, err = time.Parse("2006-01-02 15:04", s[1:17])
			if err != nil {
				log.Println("time.parse: ", err)
			}
			r.a = s[19:]
			records = append(records, r)
		}
	}

	sort.Sort(Rs(records))

	// part 1
	guards := map[int][]int{}
	gid := -1
	var sleeptime time.Time
	for _, r := range records {
		if r.a == "wakes up" {
			for i := sleeptime.Minute(); i < r.t.Minute(); i++ {
				guards[gid][i]++
			}
		} else if r.a == "falls asleep" {
			sleeptime = r.t
		} else {
			fmt.Sscanf(r.a, "Guard #%d begins shift", &gid)
			if _, ok := guards[gid]; !ok {
				guards[gid] = make([]int, 60)
			}
		}
	}

	total, idx, maxt := 0, 0, 0
	for g, minutes := range guards {
		sum, id := 0, 0
		for i, m := range minutes {
			sum += m
			if m > maxt {
				maxt = m
				id = i
			}
		}
		if sum > total {
			total = sum
			idx = id
			gid = g
		}
	}
	fmt.Printf("Strategy 1: %v x %v = %v\n", gid, idx, gid*idx)

	// part 2
	maxt = 0
	for g, minutes := range guards {
		for i, m := range minutes {
			if m > maxt {
				maxt = m
				gid = g
				idx = i
			}
		}
	}
	fmt.Printf("Strategy 2: %vx%v = %v\n", gid, idx, gid*idx)
}
