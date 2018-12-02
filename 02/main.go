package main

import (
	"fmt"
)

func main() {
	boxes := []string{}

	for {
		s := ""
		if _, err := fmt.Scanf("%s", &s); err != nil {
			break
		}
		boxes = append(boxes, s)
	}

	// part 1
	// count chars with a set
	two, three := 0, 0
	for _, box := range boxes {
		cs := map[rune]int{}
		for _, c := range box {
			cs[c]++
		}
		c2, c3 := false, false
		for _, v := range cs {
			if v == 2 && !c2 {
				two++
				c2 = true
			}
			if v == 3 && !c3 {
				three++
				c3 = true
			}
		}
	}
	fmt.Println("checksum: ", two*three)

	// part 2
	// mask a char and compare with a set
	for i := 0; i < len(boxes[0]); i++ {
		set := map[string]struct{}{}
		for _, word := range boxes {
			mword := word[:i] + word[i+1:]
			if _, ok := set[mword]; ok {
				fmt.Println("common ID: ", mword)
				return
			}
			set[mword] = struct{}{}
		}
	}
}
