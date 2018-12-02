package main

import "fmt"

func main() {
	sum := 0
	drift := []int{}

	for {
		d := 0
		if _, err := fmt.Scanf("%d", &d); err != nil {
			break
		}
		drift = append(drift, d)
		sum += d
	}

	// part 1
	fmt.Println("drift: ", sum)

	// part 2
	curr := 0
	vis := map[int]struct{}{0: struct{}{}}
	for {
		for _, v := range drift {
			curr += v
			if _, ok := vis[curr]; ok {
				fmt.Println("visited: ", curr)
				return
			}
			vis[curr] = struct{}{}
		}
	}
}
