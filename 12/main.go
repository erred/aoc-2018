package main

import (
	"fmt"
	"strings"
)

func main() {
	s0 := "#........#.#.#...###..###..###.#..#....###.###.#.#...####..##..##.#####..##...#.#.....#...###.#.####"
	tr := map[string]string{"#..##": ".",
		"##..#": "#",
		"..##.": ".",
		".##.#": "#",
		".....": ".",
		"..###": "#",
		"###.#": "#",
		"#....": ".",
		"#.##.": "#",
		".#.##": "#",
		"#...#": ".",
		"...##": ".",
		"###..": "#",
		".#..#": ".",
		"####.": ".",
		"....#": ".",
		"#####": "#",
		".###.": ".",
		"#..#.": ".",
		"##...": "#",
		".#...": "#",
		"#.#.#": ".",
		"..#..": "#",
		"...#.": "#",
		"##.#.": ".",
		".##..": "#",
		".#.#.": ".",
		"#.#..": ".",
		"..#.#": "#",
		"#.###": ".",
		"##.##": ".",
		".####": "#"}
	p0, p1, p2 := 0, 0, 0
	for i, v := range s0 {
		if v == '#' {
			p0 += i
		}
	}
	p0, p1, p2 = 0, p0, p1
	g := 0
	for ; g < 200; g++ {
		s0 = "....." + s0 + "....."
		s1 := strings.Builder{}
		for i := range s0 {
			if i > len(s0)-5 {
				continue
			}
			s1.WriteString(tr[s0[i:i+5]])
		}
		s0 = s1.String()

		for i, v := range s0 {
			if v == '#' {
				p0 += i - (g+1)*3
			}
		}
		if g == 19 {
			fmt.Println("sum @ 20: ", p0)
		}
		if p0-p1 == p1-p2 {
			break
		}
		p0, p1, p2 = 0, p0, p1
	}
	fmt.Println("sum @ 5e10: ", p1+(p1-p2)*(50000000000-g))
}