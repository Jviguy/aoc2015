package aoc 

import "fmt"

func DayOne(input string) {
	f := 0
	p := 0
	for i, c := range input {
		if c == '(' {
			f += 1
		} else if c == ')' {
			f -= 1
			if f < 0 && p == 0 {
				p = i+1
			}
		}
	}
	fmt.Printf("Part One Solution: %d\n", f)
	fmt.Printf("Part Two Solution: %d\n", p)
}