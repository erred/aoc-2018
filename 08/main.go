package main

import "fmt"

type Node struct {
	C []Node
	M []int
	V int
}

var sum int

func main() {
	arr := input

	node, _ := findNode(arr)

	// part 1
	fmt.Println("meta sum: ", sum)

	// part 2
	fmt.Println("root value: ", node.V)
}

func findNode(arr []int) (Node, []int) {
	children, meta := arr[0], arr[1]
	arr = arr[2:]
	node := Node{}

	for i := 0; i < children; i++ {
		n := Node{}
		n, arr = findNode(arr)
		node.C = append(node.C, n)
	}

	node.M = arr[:meta]

	for _, v := range node.M {
		sum += v

		if len(node.C) == 0 {
			node.V += v
		} else {
			if v > 0 && v < len(node.C)+1 {
				node.V += node.C[v-1].V
			}
		}
	}

	return node, arr[meta:]
}
