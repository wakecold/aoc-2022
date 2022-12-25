package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	scan := bufio.NewScanner(f)
	data := make([][]int, 0)

	for scan.Scan() {
		raw := scan.Text()
		d := []int{}
		for _, c := range raw {
			val, _ := strconv.Atoi(string(c))
			d = append(d, val)
		}
		data = append(data, d)

	}

	count := len(data)*4 - 4

	for i := 1; i < len(data)-1; i++ {

	}
}
