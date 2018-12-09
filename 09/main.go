package main

import "fmt"

type Node struct {
	v    int
	p, n *Node
}

func main() {
	players := make([]int, 413)
	m := 0
	p := &Node{}
	p.p, p.n = p, p
	pv, pn := p, p

	for m < 7108200 {
		m++
		if m%23 == 0 {
			p = p.p.p.p.p.p.p.p
			pv, pn = p.p, p.n
			players[m%len(players)] += m + p.v
			pv.n, pn.p = pn, pv
			p = pn
		} else {
			p = p.n
			pv, pn = p, p.n
			p = &Node{v: m}
			pv.n, p.n = p, pn
			pn.p, p.p = p, pv
		}
	}

	hs := 0
	for _, p := range players {
		if p > hs {
			hs = p
		}
	}
	fmt.Println("highscore: ", hs)
}
