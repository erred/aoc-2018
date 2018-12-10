package main

import "fmt"

type Point struct {
	x, y, vx, vy int
}

func main() {
	var err error
	ps := []Point{}
	for {
		p := Point{}
		_, err = fmt.Scanf("position=<%d, %d> velocity=<%d,  %d>\n", &p.x, &p.y, &p.vx, &p.vy)
		if err != nil {
			break
		}
		ps = append(ps, p)
	}

	arr := make([][]bool, 20)
	for j := range arr {
		arr[j] = make([]bool, 100)
	}
	for i := 0; i < 10408; i++ {
		for j, p := range ps {
			ps[j].x += p.vx
			ps[j].y += p.vy
			if i >= 10407 {
				if p.x < 280 && p.x > 180 && p.y > 130 && p.y < 150 {
					arr[p.y-130][p.x-180] = true
				}
			}
		}
	}
	for _, row := range arr {
		for _, v := range row {
			if v {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
