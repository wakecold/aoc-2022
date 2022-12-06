package main

import (
	"bufio"
	"fmt"
	"os"
)

const lowercaseLetterOffset = 96
const capitalLetterOffset = 38

func main() {
	f, _ := os.Open("input.txt")
	scan := bufio.NewScanner(f)
	res := 0
	for scan.Scan() {
		line := scan.Text()
		v := findCommonChar(line)
		if v == 0 {
			continue
		}

		if v >= 97 && v <= 122 {
			//lowercase
			v = v - lowercaseLetterOffset
		} else if v >= 65 && v <= 90 {
			//capital
			v = v - capitalLetterOffset
		}
		res += v
	}
	fmt.Println("Result:", res)
}

// Returns 0 for no char
func findCommonChar(s string) int {
	l := len(s)
	h := l / 2

	m := make(map[int]bool)
	for i := 0; i < h; i++ {
		m[int(s[i])] = true
	}
	for i := h; i < l; i++ {
		if _, ok := m[int(s[i])]; ok {
			return int(s[i])
		}
	}
	return 0
}
