package main

import "fmt"

type Claim struct {
	id, left, right, top, down int
}

func main() {
	var err error
	data := []Claim{}
	sheet := make([][]int, 1000)
	for i := range sheet {
		sheet[i] = make([]int, 1000)
	}

	for err == nil {
		c := Claim{}
		_, err = fmt.Scanf("#%d @ %d,%d: %dx%d", &c.id, &c.left, &c.top, &c.right, &c.down)
		c.right += c.left
		c.down += c.top
		data = append(data, c)
	}

	// part 1
	for _, c := range data {
		for w := c.left; w < c.right; w++ {
			for h := c.top; h < c.down; h++ {
				sheet[w][h]++
			}
		}
	}
	sum := 0
	for _, row := range sheet {
		for _, cell := range row {
			if cell > 1 {
				sum++
			}
		}
	}
	fmt.Println("overlapping cells: ", sum)

	// part 2
checkclaim:
	for _, c := range data {
		good := true
		for w := c.left; w < c.right; w++ {
			for h := c.top; h < c.down; h++ {
				if sheet[w][h] != 1 {
					good = false
					continue checkclaim
				}
			}
		}
		if good {
			fmt.Println("no overlap claim: ", c.id)
			break checkclaim
		}
	}
}
