package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rockValue     = 1
	paperValue    = 2
	scissorsValue = 3
	loseValue     = 0
	drawValue     = 3
	winValue      = 6
)

var outcome = map[string]map[string]int{
	// enemy chose rock
	"A": {
		"X": loseValue + scissorsValue,
		"Y": drawValue + rockValue,
		"Z": winValue + paperValue,
	},
	//enemy chose paper
	"B": {
		"X": loseValue + rockValue,
		"Y": drawValue + paperValue,
		"Z": winValue + scissorsValue,
	},
	//enemy chose scissors
	"C": {
		"X": loseValue + paperValue,
		"Y": drawValue + scissorsValue,
		"Z": winValue + rockValue,
	},
}

func main() {
	fmt.Println("vim-go")
	// A for Rock, B for Paper, and C for Scissors.
	// X for Rock, Y for Paper, and Z for Scissors.

	// The score for a single round is the score for
	// the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)

	// outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	// X means you need to lose,
	// Y means you need to end the round in a draw,
	// and Z means you need to win. Good luck!"
	f, _ := os.Open("./input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	points := 0
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		enemy, own := data[0], data[1]
		points += outcome[enemy][own]
	}

	fmt.Println("Result:", points)
}
