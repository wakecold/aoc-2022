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

	// idxToLookup := []int{20, 60, 100, 140, 180, 220}
	idxToLookup := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	count := 1
	res := 0
	for idx, val := range q {
		fmt.Println(val)

		if idxToLookup[idx+1] {
			res += count * (idx + 1)
			fmt.Println("count", count)
			fmt.Println("cur res and idx:", idx, res)
		}

		count += val
	}

	fmt.Println("count:", count)
	fmt.Println("res:", res)
}
