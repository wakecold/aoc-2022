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
	i := 0
	groups := make([]string, 3)
	for scan.Scan() {
		line := scan.Text()
		k := i % 3
		groups[k] = line
		i++
		if k != 2 {
			continue
		}

		v := findCommonChar(groups)
		groups = make([]string, 3)
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
func findCommonChar(groups []string) int {
	m1 := make(map[int]bool)
	for _, c := range groups[0] {
		v := int(c)
		m1[v] = true
	}
	m2 := make(map[int]bool)
	for _, c := range groups[1] {
		v := int(c)
		if _, ok := m1[v]; ok {
			m2[v] = true
		}
	}

	for _, c := range groups[2] {
		v := int(c)
		if _, ok := m2[v]; ok {
			return v
		}
	}
	return 0
}
