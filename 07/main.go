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
		for todo, blockers := range rawSteps1 {
			if len(blockers) == 0 {
				avail = append(avail, todo)
			}
		}
		sort.Ints(avail)
		steps = append(steps, string(rune(avail[0])))
		for todo, blockers := range rawSteps1 {
			for blocker := range blockers {
				if blocker == avail[0] {
					delete(rawSteps1[todo], blocker)
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
		for todo, blockers := range rawSteps2 {
			if len(blockers) == 0 {
				avail = append(avail, todo)
			}
		}
		sort.Ints(avail)

		for _, todo := range avail {
			if _, ok := workers[todo]; !ok && len(workers) < 5 {
				workers[todo] = t + todo - int('A') + 61
			}
		}

		mint := []int{}
		for _, doneTime := range workers {
			mint = append(mint, doneTime)
		}
		sort.Ints(mint)
		t = mint[0]

		for task, doneTime := range workers {
			if doneTime == t {
				delete(workers, task)
				delete(rawSteps2, task)
				for todo, blockers := range rawSteps2 {
					for blocker := range blockers {
						if blocker == task {
							delete(rawSteps2[todo], blocker)
						}
					}
				}
			}
		}
	}
	fmt.Println("total time: ", t)

}
