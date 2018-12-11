package main

import (
	"fmt"
)

var ser = 7165

func main() {
	marr := make([][][]int, 300)
	for s := range marr {
		marr[s] = make([][]int, 300-s)
		for i := range marr[s] {
			marr[s][i] = make([]int, 300-s)
		}
	}

	var maxv, maxx, maxy, maxs int
	var maxv3, maxx3, maxy3 int
	for s := range marr {
		for y := range marr[s] {
			for x := range marr[s][y] {
				if s == 0 {
					v := ((x+1+10)*(y+1) + ser) * (x + 1 + 10)
					v %= 1000
					v /= 100
					v -= 5
					marr[0][x][y] = v
				} else {
					marr[s][x][y] = marr[s-1][x][y]
					for i := 0; i < s; i++ {
						marr[s][x][y] += marr[0][x+s][y+i]
						marr[s][x][y] += marr[0][x+i][y+s]
					}
					marr[s][x][y] += marr[0][x+s][y+s]
				}
				if marr[s][x][y] > maxv {
					maxv, maxx, maxy, maxs = marr[s][x][y], x+1, y+1, s+1
				}
				if s == 2 && marr[s][x][y] > maxv3 {
					maxv3, maxx3, maxy3 = marr[s][x][y], x+1, y+1
				}
			}
		}
	}

	fmt.Printf("max 3x3: %v,%v\n", maxx3, maxy3)
	fmt.Printf("max any: %v,%v,%v\n", maxx, maxy, maxs)
}
