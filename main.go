package main 

import (
	"bufio"
	"fmt"
	"github.com/jviguy/aoc2015/aoc"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the day number to run: ")
	in.Scan()
	day, _ := strconv.Atoi(string(in.Bytes()))
	aoc.Run(day)
}