package main

import (
	"fmt"
	"strings"
)

func main() {
	data := [][][]int{}
	samples := strings.Split(input, "\n\n")
	// samples := strings.Split(input0, "\n\n")
	data = make([][][]int, len(samples))
	for i, sample := range samples {
		data[i] = make([][]int, 3)
		var a, b, c, d, e, f, g, h, j, k, l, m int
		fmt.Sscanf(sample, "Before: [%d, %d, %d, %d]\n%d %d %d %d\nAfter: [%d, %d, %d, %d]", &a, &b, &c, &d, &e, &f, &g, &h, &j, &k, &l, &m)
		data[i][0] = []int{a, b, c, d}
		data[i][1] = []int{e, f, g, h}
		data[i][2] = []int{j, k, l, m}
	}

	opcodes := map[int]map[string]struct{}{}
	for i := 0; i < 16; i++ {
		opcodes[i] = map[string]struct{}{}
		for _, op := range []string{
			"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr",
		} {
			opcodes[i][op] = struct{}{}
		}
	}
	totalValid := 0
	for _, sample := range data {
		valid := []string{}
		r1, r2, r3 := sample[1][1], sample[1][2], sample[1][3]
		res := sample[2][r3]
		// all registers
		if r1 < 4 && r2 < 4 && r3 < 4 {
			rega, regb := sample[0][r1], sample[0][r2]
			// addr
			if res == rega+regb {
				valid = append(valid, "addr")
			}

			// mulr
			if res == rega*regb {
				valid = append(valid, "mulr")
			}

			// banr
			if res == rega&regb {
				valid = append(valid, "banr")
			}

			// borr
			if res == rega|regb {
				valid = append(valid, "borr")
			}

			// setr
			if res == rega {
				valid = append(valid, "setr")
			}

			// gtrr
			if rega > regb {
				if res == 1 {
					valid = append(valid, "gtrr")
				}
			} else if res == 0 {
				valid = append(valid, "gtrr")
			}

			// eqrr
			if rega == regb {
				if res == 1 {
					valid = append(valid, "eqrr")
				}
			} else if res == 0 {
				valid = append(valid, "eqrr")
			}

		}
		// addi
		rega, valb := sample[0][r1], r2
		if res == rega+valb {
			valid = append(valid, "addi")
		}

		// muli
		if res == rega*valb {
			valid = append(valid, "muli")
		}

		// bani
		if res == rega&valb {
			valid = append(valid, "bani")
		}

		// bori
		if res == rega|valb {
			valid = append(valid, "bori")
		}

		// seti
		if res == r1 {
			valid = append(valid, "seti")
		}

		// gtir
		if r1 > sample[0][r2] {
			if res == 1 {
				valid = append(valid, "gtir")
			}
		} else if res == 0 {
			valid = append(valid, "gtir")
		}

		// gtri
		if rega > valb {
			if res == 1 {
				valid = append(valid, "gtri")
			}
		} else if res == 0 {
			valid = append(valid, "gtri")
		}

		// eqir
		if r1 == sample[0][r2] {
			if res == 1 {
				valid = append(valid, "eqir")
			}
		} else if res == 0 {
			valid = append(valid, "eqir")
		}

		// eqri
		if rega == valb {
			if res == 1 {
				valid = append(valid, "eqri")
			}
		} else if res == 0 {
			valid = append(valid, "eqri")
		}

		// count
		// fmt.Println(sample)
		// fmt.Println(valid)
		if len(valid) >= 3 {
			totalValid++
		}
		op := sample[1][0]
	check:
		for maybe := range opcodes[op] {
			for _, v := range valid {
				if v == maybe {
					continue check
				}

			}
			delete(opcodes[op], maybe)
		}
	}
	fmt.Println(totalValid)

	opc := map[int]string{}
	for len(opc) < 16 {
		for k, v := range opcodes {
			if len(v) == 1 {
				for kk := range v {
					opc[k] = kk
				}
				for kk := range opcodes {
					delete(opcodes[kk], opc[k])
				}
			}
		}
	}
	// for k, v := range opc {
	// 	fmt.Println(k, v)
	// }
	lines := strings.Split(input2, "\n")
	data2 := make([][]int, len(lines))
	for i, line := range lines {
		var a, b, c, d int
		fmt.Sscanf(line, "%d %d %d %d", &a, &b, &c, &d)
		data2[i] = []int{a, b, c, d}
	}
	var regs = make([]int, 4)
	for _, inst := range data2 {
		a, b, c := inst[1], inst[2], inst[3]
		switch opc[inst[0]] {
		case "addr":
			regs[c] = regs[a] + regs[b]
		case "addi":
			regs[c] = regs[a] + b
		case "mulr":
			regs[c] = regs[a] * regs[b]
		case "muli":
			regs[c] = regs[a] * b
		case "banr":
			regs[c] = regs[a] & regs[b]
		case "bani":
			regs[c] = regs[a] & b
		case "borr":
			regs[c] = regs[a] | regs[b]
		case "bori":
			regs[c] = regs[a] | b
		case "setr":
			regs[c] = regs[a]
		case "seti":
			regs[c] = a
		case "gtir":
			if a > regs[b] {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		case "gtri":
			if regs[a] > b {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		case "gtrr":
			if regs[a] > regs[b] {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		case "eqir":
			if a == regs[b] {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		case "eqri":
			if regs[a] == b {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		case "eqrr":
			if regs[a] == regs[b] {
				regs[c] = 1
			} else {
				regs[c] = 0
			}
		}
	}
	fmt.Println(regs)

}
