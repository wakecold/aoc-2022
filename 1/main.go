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

	curMax := 0
	buf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if curMax < buf {
				curMax = buf
			}
			buf = 0
		} else {
			v, _ := strconv.Atoi(line)
			buf += v
		}
	}

	fmt.Println(curMax)
}
