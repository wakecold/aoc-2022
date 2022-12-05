package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	topThree := make([]int, 3) // [0]small, [1]big, [2]bigger
	buf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if buf > topThree[2] {
				topThree[0] = topThree[1]
				topThree[1] = topThree[2]
				topThree[2] = buf
			} else if buf > topThree[1] {
				topThree[0] = topThree[1]
				topThree[1] = buf
			} else if buf > topThree[0] {
				topThree[0] = buf
			}
			buf = 0
		} else {
			v, _ := strconv.Atoi(line)
			buf += v
		}
	}

	fmt.Println(topThree)
	s := 0
	for _, v := range topThree {
		s += v
	}
	fmt.Println("sum:", s)
}
