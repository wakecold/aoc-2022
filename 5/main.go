package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	s := bufio.NewScanner(f)

	i := 0
	crates := map[int][]rune{}
	for s.Scan() {
		if i == 8 {
			break
		}
		line := s.Text()
		fmt.Println("line", line)
		// got this idea from aoc subreddit
		// thanks to u/satylogin
		if strings.Contains(row_raw, string('[')) {
			for i := 0; i < len(row); i += 4 {
				if row[i+1] != ' ' {
					crates[i/4+1]=append(crates[i/4+1],
				}
			}
		}
		i++
	}
}
