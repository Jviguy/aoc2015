package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func DayTwo(input string) {
	t:=0
	for _, line := range strings.Split(input, "\n") {
		x := strings.Split(line, "x")
		s := make([]int,0,len(x))
		for _, n := range x {
			w, _ := strconv.Atoi(n)
			fmt.Println(w)
			s = append(s, w) 
		}
		t += 2*s[0]*s[1] + 2*s[1]*s[2] + 2*s[0]+s[2]
		fmt.Println(t)
		smallest := 100
		for _, side := range []int{s[0]*s[1], s[1]*s[2], s[2]*s[0]} {
			if side < smallest {
				smallest = side
			}
		}
		fmt.Println(smallest)
		t+=smallest
	}
	fmt.Printf("Part One Solution: %d\n", t)
}
