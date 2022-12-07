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
	count := 0
	for s.Scan() {
		line := s.Text()
		elfs := strings.Split(line, ",")
		e1, e2 := elfs[0], elfs[1]

		// l1-r1, l2-r2
		e1range := strings.Split(e1, "-")
		l1s, r1s := e1range[0], e1range[1]

		e2range := strings.Split(e2, "-")
		l2s, r2s := e2range[0], e2range[1]

		l1, _ := strconv.Atoi(l1s)
		l2, _ := strconv.Atoi(l2s)
		r1, _ := strconv.Atoi(r1s)
		r2, _ := strconv.Atoi(r2s)

		// 5-17,4-67
		// l1 <= l2 <= r1 && l1 <= r2 <= r1
		e1OverlapL := (l1 <= l2 && l2 <= r1)
		e1OverlapR := (l1 <= r2 && r2 <= r1)
		e1Contains := e1OverlapL && e1OverlapR
		// l2 <= l1 <= r2 && l2 <= r1 <= r2
		e2OverlapL := (l2 <= l1 && l1 <= r2)
		e2OverlapR := (l2 <= r1 && r1 <= r2)
		e2Contains := e2OverlapL && e2OverlapR
		overlap := e1Contains || e2Contains || e1OverlapL || e1OverlapR || e2OverlapL || e2OverlapR
		if overlap {

			fmt.Println(e1Contains, e2Contains, line, l1, r1, l2, r2)
			count++
		}
	}
	fmt.Println("Result:", count)
}
