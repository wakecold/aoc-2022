package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")

	s := bufio.NewScanner(f)
	q := []int{}
	for s.Scan() {
		line := s.Text()
		cmd := strings.Split(line, " ")

		switch len(cmd) {
		case 1:
			// do nothing, noop
			q = append(q, 0)
		case 2:
			// addx
			val, _ := strconv.Atoi(cmd[1])
			q = append(q, 0, val)
		default:
			// skip
		}
	}

	count := 1
	fmt.Println(q)
	for idx, val := range q {
		//fmt.Println(val)
		i := idx % 40
		if i == 0 {
			fmt.Printf("\n")
		}
		//fmt.Printf("i: %v, val: %v, res: ", idx, val)
		if i > count-2 && i < count+2 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		count += val

	}

}
