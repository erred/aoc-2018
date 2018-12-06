package main

import (
	"fmt"
)

type xy struct {
	x, y int
}
type di struct {
	d, i, r int
}

func main() {
	var err error
	pos := []xy{xy{0, 0}}
	maxx, maxy := 0, 0
	for {
		p := xy{}
		_, err = fmt.Scanf("%d, %d\n", &p.x, &p.y)
		if err != nil {
			break
		}
		pos = append(pos, p)
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	grid := make([][]di, maxy+1)
	for i := range grid {
		grid[i] = make([]di, maxx+1)
	}

	// part 1
	for i, p := range pos[1:] {
		for y, row := range grid {
			for x, v := range row {
				d := abs(p.x-x) + abs(p.y-y)

				if v.d > d || (v.i == 0 && v.d == 0) {
					grid[y][x].d, grid[y][x].i = d, i
				} else if v.d == d {
					grid[y][x].i = 0
				}

				// part 2
				grid[y][x].r += d
			}
		}
	}

	region := 0
	inf := map[int]struct{}{}
	areas := make([]int, len(pos))
	for y, row := range grid {
		for x, v := range row {
			if v.i != 0 {
				areas[v.i]++
			}
			if x == 0 || x == maxx || y == 0 || y == maxy {
				inf[v.i] = struct{}{}
			}

			// part 2
			if v.r < 10000 {
				region++
			}
		}
	}
	maxv, maxi := 0, 0
	for i, v := range areas {
		if _, ok := inf[i]; !ok && v > maxv {
			maxv, maxi = v, i
		}
	}
	fmt.Printf("max area %v: %v\n", maxi, maxv)

	// part 2
	fmt.Printf("region: %v\n", region)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
