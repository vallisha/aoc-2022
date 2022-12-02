package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scores := map[string][]int{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	total1, total2 := 0, 0
	for _, scr := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		total1 += scores[scr][0]
		total2 += scores[scr][1]
	}
	fmt.Println(total1)
	fmt.Println(total2)
}
