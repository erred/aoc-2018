package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var err error
	// todo, blockers, ...
	rawSteps1 := map[int]map[int]struct{}{}
	rawSteps2 := map[int]map[int]struct{}{}
	for {
		var one, two string
		_, err = fmt.Scanf("Step %s must be finished before step %s can begin.\n", &one, &two)
		if err != nil || one == "" {
			break
		}
		if _, ok := rawSteps1[int(two[0])]; !ok {
			rawSteps1[int(two[0])] = map[int]struct{}{}
			rawSteps2[int(two[0])] = map[int]struct{}{}
		}
		if _, ok := rawSteps1[int(one[0])]; !ok {
			rawSteps1[int(one[0])] = map[int]struct{}{}
			rawSteps2[int(one[0])] = map[int]struct{}{}
		}
		rawSteps1[int(two[0])][int(one[0])] = struct{}{}
		rawSteps2[int(two[0])][int(one[0])] = struct{}{}
	}

	// part 1
	steps := []string{}
	for len(rawSteps1) > 0 {
		avail := []int{}
		for k, v := range rawSteps1 {
			if len(v) == 0 {
				avail = append(avail, k)
			}
		}
		sort.Ints(avail)
		steps = append(steps, string(rune(avail[0])))
		for k, v := range rawSteps1 {
			for kk, _ := range v {
				if kk == avail[0] {
					delete(rawSteps1[k], kk)
				}
			}
		}
		delete(rawSteps1, avail[0])
	}
	fmt.Printf("steps: %v\n", strings.Join(steps, ""))

	// part 2
	// step, time done
	workers := map[int]int{}
	t := 0
	for len(rawSteps2) > 0 {
		avail := []int{}
		for k, v := range rawSteps2 {
			if len(v) == 0 {
				avail = append(avail, k)
			}
		}
		sort.Ints(avail)

		for _, v := range avail {
			if _, ok := workers[v]; !ok && len(workers) < 5 {
				workers[v] = t + v - int('A') + 61
			}
		}

		mint := []int{}
		for _, v := range workers {
			mint = append(mint, v)
		}
		sort.Ints(mint)
		t = mint[0]

		for k, v := range workers {
			if v == t {
				delete(workers, k)
				delete(rawSteps2, k)
				for kk, vv := range rawSteps2 {
					for kkk, _ := range vv {
						if kkk == k {
							delete(rawSteps2[kk], kkk)
						}
					}
				}
			}
		}
	}
	fmt.Println("total time: ", t)

}
