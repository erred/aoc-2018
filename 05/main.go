package main

import (
	"fmt"
	"strings"
)

var d = 'a' - 'A'

func main() {
	var s string
	fmt.Scan(&s)

	// part 1
	fmt.Printf("units: %v\n", findLen(s))

	// part 2
	minlen := len(s)
	for c := 'a'; c <= 'z'; c++ {
		ss := strings.Replace(s, string(c), "", -1)
		ss = strings.Replace(ss, string(c-d), "", -1)
		l := findLen(ss)
		if minlen > l {
			minlen = l
		}

	}
	fmt.Printf("Min len: %v\n", minlen)
}
func findLen(s string) int {
	i := 0
	for i < len(s)-1 {
		diff := rune(s[i]) - rune(s[i+1])
		if diff == d || diff == -d {
			if i == 0 {
				s = s[2:]
			} else if i == len(s)-2 {
				s = s[:len(s)-2]
				break
			} else {
				s = s[:i] + s[i+2:]
				i--
			}
		} else {
			i++
		}
	}
	return len(s)
}
