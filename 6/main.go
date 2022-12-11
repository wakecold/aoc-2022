package main

import (
	"fmt"
	"io"
	"os"
)

const windowSize = 14

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	data, _ := io.ReadAll(f)
	l, r := 0, windowSize
	q := data[l:r]
	//brute force
	for r < len(data) {
		q = data[l:r]
		m := make(map[byte]bool)
		for _, val := range q {
			if _, ok := m[val]; !ok {
				m[val] = true
			}
		}
		if len(m) == windowSize {
			fmt.Println("Result: ", r)
			break
		}

		q = q[1:]
		l++
		r++
	}
	fmt.Println("Finished")
}
