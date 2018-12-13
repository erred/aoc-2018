package main

import (
	"fmt"
	"sort"
	"strings"
)

type Cart struct {
	x, y             int
	orig, face, turn string
	crashed          bool
}

type Carts []Cart

func (cs Carts) Less(i, j int) bool {
	if cs[i].y == cs[j].y {
		return cs[i].x < cs[j].x
	}
	return cs[i].y < cs[j].y
}
func (cs Carts) Len() int {
	return len(cs)
}
func (cs Carts) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
	return
}

func main() {
	rows := strings.Split(input, "\n")
	rows = rows[1 : len(rows)-1]

	arr := make([][]string, len(rows))
	for i := range arr {
		arr[i] = strings.Split(rows[i], "")
	}

	cs := []Cart{}

	for y, row := range arr {
		for x, cell := range row {
			switch cell {
			case "v":
				// arr[y][x] = "|"
				cs = append(cs, Cart{x, y, "|", "down", "left", false})
			case "^":
				// arr[y][x] = "|"
				cs = append(cs, Cart{x, y, "|", "up", "left", false})
			case "<":
				// arr[y][x] = "-"
				cs = append(cs, Cart{x, y, "-", "left", "left", false})
			case ">":
				// arr[y][x] = "-"
				cs = append(cs, Cart{x, y, "-", "right", "left", false})
			}
		}
	}

	crashx, crashy := 0, 0
	lefti := 0
	// loop:
	for {
		sort.Sort(Carts(cs))
		for i, c := range cs {
			if c.crashed {
				continue
			}
			y, x := c.y, c.x
			switch c.face {
			case "up":
				switch arr[y-1][x] {
				case "^", "v", "<", ">":
					crashx, crashy = x, y-1
					cs[i].crashed = true
					arr[y][x] = c.orig
					for j, cc := range cs {
						if cc.x == x && cc.y == y-1 {
							cs[j].crashed = true
							arr[y-1][x] = cc.orig
						}
					}
				case "|":
					arr[y][x], arr[y-1][x] = c.orig, "^"
					cs[i].orig = "|"
					cs[i].y--
				case "/":
					arr[y][x], arr[y-1][x] = c.orig, ">"
					cs[i].face = "right"
					cs[i].orig = "/"
					cs[i].y--
				case "\\":
					arr[y][x], arr[y-1][x] = c.orig, "<"
					cs[i].face = "left"
					cs[i].orig = "\\"
					cs[i].y--
				case "+":
					cs[i].orig = "+"
					cs[i].y--
					switch cs[i].turn {
					case "left":
						arr[y][x], arr[y-1][x] = c.orig, "<"
						cs[i].face = "left"
						cs[i].turn = "straight"
					case "straight":
						arr[y][x], arr[y-1][x] = c.orig, "^"
						cs[i].face = "up"
						cs[i].turn = "right"
					case "right":
						arr[y][x], arr[y-1][x] = c.orig, ">"
						cs[i].face = "right"
						cs[i].turn = "left"
					}
				}
			case "down":
				switch arr[y+1][x] {
				case "^", "v", "<", ">":
					crashx, crashy = x, y+1
					cs[i].crashed = true
					arr[y][x] = c.orig
					for j, cc := range cs {
						if cc.x == x && cc.y == y+1 {
							cs[j].crashed = true
							arr[y+1][x] = cc.orig
						}
					}
				case "|":
					arr[y][x], arr[y+1][x] = c.orig, "v"
					cs[i].orig = "|"
					cs[i].y++
				case "/":
					arr[y][x], arr[y+1][x] = c.orig, "<"
					cs[i].orig = "/"
					cs[i].face = "left"
					cs[i].y++
				case "\\":
					arr[y][x], arr[y+1][x] = c.orig, ">"
					cs[i].orig = "\\"
					cs[i].face = "right"
					cs[i].y++
				case "+":
					cs[i].orig = "+"
					cs[i].y++
					switch cs[i].turn {
					case "left":
						arr[y][x], arr[y+1][x] = c.orig, ">"
						cs[i].face = "right"
						cs[i].turn = "straight"
					case "straight":
						arr[y][x], arr[y+1][x] = c.orig, "v"
						cs[i].face = "down"
						cs[i].turn = "right"
					case "right":
						arr[y][x], arr[y+1][x] = c.orig, "<"
						cs[i].face = "left"
						cs[i].turn = "left"
					}
				}
			case "left":
				switch arr[y][x-1] {
				case "^", "v", "<", ">":
					crashx, crashy = x-1, y
					cs[i].crashed = true
					arr[y][x] = c.orig
					for j, cc := range cs {
						if cc.x == x-1 && cc.y == y {
							cs[j].crashed = true
							arr[y][x-1] = cc.orig
						}
					}
				case "-":
					arr[y][x], arr[y][x-1] = c.orig, "<"
					cs[i].orig = "-"
					cs[i].x--
				case "/":
					arr[y][x], arr[y][x-1] = c.orig, "v"
					cs[i].orig = "/"
					cs[i].face = "down"
					cs[i].x--
				case "\\":
					arr[y][x], arr[y][x-1] = c.orig, "^"
					cs[i].orig = "\\"
					cs[i].face = "up"
					cs[i].x--
				case "+":
					cs[i].orig = "+"
					cs[i].x--
					switch cs[i].turn {
					case "left":
						arr[y][x], arr[y][x-1] = c.orig, "v"
						cs[i].face = "down"
						cs[i].turn = "straight"
					case "straight":
						arr[y][x], arr[y][x-1] = c.orig, "<"
						cs[i].face = "left"
						cs[i].turn = "right"
					case "right":
						arr[y][x], arr[y][x-1] = c.orig, "^"
						cs[i].face = "up"
						cs[i].turn = "left"
					}
				}
			case "right":
				switch arr[y][x+1] {
				case "^", "v", "<", ">":
					crashx, crashy = x+1, y
					cs[i].crashed = true
					arr[y][x] = c.orig
					for j, cc := range cs {
						if cc.x == x+1 && cc.y == y {
							cs[j].crashed = true
							arr[y][x+1] = cc.orig
						}
					}
				case "-":
					arr[y][x], arr[y][x+1] = c.orig, ">"
					cs[i].orig = "-"
					cs[i].x++
				case "/":
					arr[y][x], arr[y][x+1] = c.orig, "^"
					cs[i].orig = "/"
					cs[i].face = "up"
					cs[i].x++
				case "\\":
					arr[y][x], arr[y][x+1] = c.orig, "v"
					cs[i].orig = "\\"
					cs[i].face = "down"
					cs[i].x++
				case "+":
					cs[i].orig = "+"
					cs[i].x++
					switch cs[i].turn {
					case "left":
						arr[y][x], arr[y][x+1] = c.orig, "^"
						cs[i].face = "up"
						cs[i].turn = "straight"
					case "straight":
						arr[y][x], arr[y][x+1] = c.orig, ">"
						cs[i].face = "right"
						cs[i].turn = "right"
					case "right":
						arr[y][x], arr[y][x+1] = c.orig, "v"
						cs[i].face = "down"
						cs[i].turn = "left"
					}
				}
			}
		}
		left := 0
		for i, c := range cs {
			if !c.crashed {
				left++
				lefti = i
			}
		}
		if left == 1 {
			break
		}
	}
	fmt.Printf("crashed at: %v,%v", crashx, crashy)
	fmt.Printf("final cart at: %v,%v", cs[lefti].x, cs[lefti].y)
}
