package main

import "fmt"
import "reflect"

func main() {
	// input := 2018
	input := 170641
	inarr := []int{1, 7, 0, 6, 4, 1}
	recipes := make([]int, 2, input+11)
	recipes[0], recipes[1] = 3, 7
	e1, e2 := 0, 1
	for {
		nr := recipes[e1] + recipes[e2]
		if nr > 9 {
			recipes = append(recipes, 1)
		}
		recipes = append(recipes, nr%10)
		e1 = (e1 + recipes[e1] + 1) % len(recipes)
		e2 = (e2 + recipes[e2] + 1) % len(recipes)
		// fmt.Println(recipes)
		if len(recipes) > 7 {
			if reflect.DeepEqual(inarr, recipes[len(recipes)-6:]) {
				break
			}
			if reflect.DeepEqual(inarr, recipes[len(recipes)-7:len(recipes)-1]) {
				recipes = recipes[:len(recipes)-1]
				break
			}
		}
		if len(recipes) == input+10 || len(recipes) == input+11 {
			for _, v := range recipes[input : input+10] {
				fmt.Print(v)
			}
			fmt.Println()
		}
	}
	fmt.Println(len(recipes) - 6)
}
