package main

import (
	"fmt"
	"sort"
	"strings"
)

type E struct {
	pos  complex64
	t    int
	hp   int
	dead bool
}
type Es []E

func (e Es) Less(i, j int) bool {
	if imag(e[i].pos) == imag(e[j].pos) {
		return real(e[i].pos) < real(e[j].pos)
	}
	return imag(e[i].pos) < imag(e[j].pos)
}
func (e Es) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e Es) Len() int {
	return len(e)
}

type Ps []complex64

func (e Ps) Less(i, j int) bool {
	if imag(e[i]) == imag(e[j]) {
		return real(e[i]) < real(e[j])
	}
	return imag(e[i]) < imag(e[j])
}
func (e Ps) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
func (e Ps) Len() int {
	return len(e)
}

type C struct {
	// 0 space 1 wall 2 elve 3 goblin
	t  int
	hp int
}

func main() {
	lines := strings.Split(input, "\n")
	grid := map[complex64]C{}
	ents := []E{}

	for y, line := range lines {
		for x, c := range line {
			t := 0
			if c == '#' {
				t = 1
			} else if c == 'E' {
				t = 2
			} else if c == 'G' {
				t = 3
			}
			grid[complex(float32(x), float32(y))] = C{
				t, 200,
			}
			if t == 2 || t == 3 {
				ents = append(ents, E{complex(float32(x), float32(y)), t, 200, false})
			}
		}
	}

	// arr := make([][]string, len(lines))
	// for i := range arr {
	// 	arr[i] = make([]string, len(lines[i]))
	// }

	var rounds, hpsum int

	var bkents []E
	for _, v := range ents {
		bkents = append(bkents, v)
	}
	bkgrid := map[complex64]C{}
	for k, v := range grid {
		bkgrid[k] = v
	}
power:
	for atk := 4; ; atk++ {
		ents = []E{}
		for _, v := range bkents {
			ents = append(ents, v)
		}
		for k, v := range bkgrid {
			grid[k] = v
		}
		rounds, hpsum = 0, 0
		// fmt.Println(ents)
		// round:
		for {
			sort.Sort(Es(ents))

			// move
			for i, e := range ents {
				if e.dead {
					continue
				}
				hpsum = 0
				elves, goblins := 0, 0
				for _, ee := range ents {
					if ee.dead {
						continue
					}
					hpsum += ee.hp
					if ee.t == 2 {
						elves++
					} else {
						goblins++
					}
				}
				if goblins == 0 {
					break power
				}

				// not in range, move
				if grid[e.pos+1].t+e.t != 5 && grid[e.pos-1].t+e.t != 5 && grid[e.pos-1i].t+e.t != 5 && grid[e.pos+1i].t+e.t != 5 {
					// find distance to each space
					// pos - dist

					visited := map[complex64]int{}
					qu := []complex64{}
					for _, ee := range ents {
						if ee.dead || ee.pos == e.pos || ee.t == e.t {
							continue
						}
						for _, p := range []complex64{ee.pos - 1i, ee.pos - 1, ee.pos + 1, ee.pos + 1i} {
							if grid[p].t != 0 {
								continue
							}
							if _, ok := visited[p]; !ok {
								qu = append(qu, p)
								visited[p] = 1
							}
						}
					}

					for id := 0; id < len(qu); id++ {
						for _, p := range []complex64{qu[id] - 1i, qu[id] - 1, qu[id] + 1, qu[id] + 1i} {
							if grid[p].t != 0 {
								continue
							}
							if _, ok := visited[p]; !ok {
								qu = append(qu, p)
								visited[p] = visited[qu[id]] + 1
							} else if visited[p] > visited[qu[id]]+1 {
								visited[p] = visited[qu[id]] + 1
							}
						}
					}

					mind := 1000
					var np complex64
					for _, p := range []complex64{e.pos - 1i, e.pos - 1, e.pos + 1, e.pos + 1i} {

						if d, ok := visited[p]; ok && d < mind {
							mind = d
							np = p
						}
					}
					if np != 0 {
						grid[np] = grid[e.pos]
						grid[e.pos] = C{}
						ents[i].pos = np
					}
				}
				// select target
				var pos = ents[i].pos
				minhp := 201
				targets := []complex64{pos - 1i, pos - 1, pos + 1, pos + 1i}
				for _, t := range targets {
					if grid[t].t+e.t == 5 && grid[t].hp < minhp {
						minhp = grid[t].hp
						pos = t
					}
				}

				// attack target
				if minhp < 201 {
					if e.t == 3 {
						if grid[pos].hp <= 3 {
							fmt.Printf("Atk: %v Round %v, elf died\n", atk, rounds)
							continue power
						} else {
							for j, ee := range ents {
								if ee.pos == pos {
									ents[j].hp -= 3
								}
							}
							grid[pos] = C{
								grid[pos].t, grid[pos].hp - 3,
							}
						}

					} else {
						if grid[pos].hp <= atk {
							grid[pos] = C{
								0, 200,
							}
							for j, ee := range ents {
								if ee.pos == pos {
									ents[j].hp -= atk
									ents[j].dead = true
								}
							}
						} else {
							for j, ee := range ents {
								if ee.pos == pos {
									ents[j].hp -= atk
								}
							}
							grid[pos] = C{
								grid[pos].t, grid[pos].hp - atk,
							}
						}

					}
				}
			}
			rounds++

			// cleanup dead
			for i := 0; i < len(ents); {
				if ents[i].dead {
					ents[i] = ents[len(ents)-1]
					ents = ents[:len(ents)-1]
				} else {
					i++
				}
			}
			// count alive / score
			elves, goblins := 0, 0
			hpsum = 0
			for _, e := range ents {
				hpsum += e.hp
				if e.t == 2 {
					elves++
				} else {
					goblins++
				}
			}
			if goblins == 0 {
				break power
			}
		}
	}
	// round:
	// 	for {
	// 		sort.Sort(Es(ents))
	//
	// 		// move
	// 		for i, e := range ents {
	// 			elves, goblins := 0, 0
	// 			for _, ee := range ents {
	// 				if ee.dead {
	// 					continue
	// 				}
	// 				hpsum += ee.hp
	// 				if ee.t == 2 {
	// 					elves++
	// 				} else {
	// 					goblins++
	// 				}
	// 			}
	// 			if elves == 0 || goblins == 0 {
	// 				break
	// 			}
	//
	// 			if e.dead {
	// 				continue
	// 			}
	// 			fmt.Printf("\nRound %v, ent %v\n", rounds, i)
	// 			for p, c := range grid {
	// 				v := ""
	// 				switch c.t {
	// 				case 0:
	// 					v = " "
	// 				case 1:
	// 					v = "#"
	// 				case 2:
	// 					v = "E"
	// 				case 3:
	// 					v = "G"
	// 				}
	// 				arr[int(imag(p))][int(real(p))] = v
	// 			}
	// 			for _, row := range arr {
	// 				for _, v := range row {
	// 					fmt.Print(v)
	// 				}
	// 				fmt.Println()
	// 			}
	// 			fmt.Println(ents)
	//
	// 			// not in range, move
	// 			if grid[e.pos+1].t+e.t != 5 && grid[e.pos-1].t+e.t != 5 && grid[e.pos-1i].t+e.t != 5 && grid[e.pos+1i].t+e.t != 5 {
	// 				// find distance to each space
	// 				// pos - dist
	//
	// 				visited := map[complex64]int{}
	// 				qu := []complex64{}
	// 				for _, ee := range ents {
	// 					if ee.dead || ee.pos == e.pos || ee.t == e.t {
	// 						continue
	// 					}
	// 					for _, p := range []complex64{ee.pos - 1i, ee.pos - 1, ee.pos + 1, ee.pos + 1i} {
	// 						if grid[p].t != 0 {
	// 							continue
	// 						}
	// 						if _, ok := visited[p]; !ok {
	// 							qu = append(qu, p)
	// 							visited[p] = 1
	// 						}
	// 					}
	// 				}
	//
	// 				for id := 0; id < len(qu); id++ {
	// 					for _, p := range []complex64{qu[id] - 1i, qu[id] - 1, qu[id] + 1, qu[id] + 1i} {
	// 						if grid[p].t != 0 {
	// 							continue
	// 						}
	// 						if _, ok := visited[p]; !ok {
	// 							qu = append(qu, p)
	// 							visited[p] = visited[qu[id]] + 1
	// 						} else if visited[p] > visited[qu[id]]+1 {
	// 							visited[p] = visited[qu[id]] + 1
	// 						}
	// 					}
	// 				}
	//
	// 				// for p, d := range visited {
	// 				// 	arr[int(imag(p))][int(real(p))] = strconv.Itoa(d)
	// 				// }
	// 				// for _, row := range arr {
	// 				// 	for _, v := range row {
	// 				// 		fmt.Print(v)
	// 				// 	}
	// 				// 	fmt.Println()
	// 				// }
	// 				// fmt.Println()
	//
	// 				mind := 1000
	// 				var np complex64
	// 				for _, p := range []complex64{e.pos - 1i, e.pos - 1, e.pos + 1, e.pos + 1i} {
	//
	// 					if d, ok := visited[p]; ok && d < mind {
	// 						mind = d
	// 						np = p
	// 					}
	// 				}
	// 				if np != 0 {
	// 					grid[np] = grid[e.pos]
	// 					grid[e.pos] = C{}
	// 					ents[i].pos = np
	// 				}
	// 			}
	// 			// select target
	// 			var pos = ents[i].pos
	// 			minhp := 201
	// 			targets := []complex64{pos - 1i, pos - 1, pos + 1, pos + 1i}
	// 			for _, t := range targets {
	// 				if grid[t].t+e.t == 5 && grid[t].hp < minhp {
	// 					minhp = grid[t].hp
	// 					pos = t
	// 				}
	// 			}
	//
	// 			// attack target
	// 			if minhp < 201 {
	// 				if grid[pos].hp <= 3 {
	// 					grid[pos] = C{
	// 						0, 200,
	// 					}
	// 					for j, ee := range ents {
	// 						if ee.pos == pos {
	// 							ents[j].hp -= 3
	// 							ents[j].dead = true
	// 						}
	// 					}
	// 				} else {
	// 					for j, ee := range ents {
	// 						if ee.pos == pos {
	// 							ents[j].hp -= 3
	// 						}
	// 					}
	// 					grid[pos] = C{
	// 						grid[pos].t, grid[pos].hp - 3,
	// 					}
	// 				}
	// 			}
	// 		}
	// 		rounds++
	//
	// 		// cleanup dead
	// 		for i := 0; i < len(ents); {
	// 			if ents[i].dead {
	// 				ents[i] = ents[len(ents)-1]
	// 				ents = ents[:len(ents)-1]
	// 			} else {
	// 				i++
	// 			}
	// 		}
	// 		// count alive / score
	// 		elves, goblins := 0, 0
	// 		hpsum = 0
	// 		for _, e := range ents {
	// 			hpsum += e.hp
	// 			if e.t == 2 {
	// 				elves++
	// 			} else {
	// 				goblins++
	// 			}
	// 		}
	// 		if elves == 0 || goblins == 0 {
	// 			break round
	// 		}
	// 	}
	fmt.Printf("%v x %v = %v", rounds, hpsum, rounds*hpsum)

}
