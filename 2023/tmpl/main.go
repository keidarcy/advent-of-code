package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input1 := ``
	lines1 := strings.Split(input1, "\n")
	puzzleOne(lines1)
	puzzleTwo(lines1)

	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("fail to read file")
	}
	lines := strings.Split(string(content), "\n")
	puzzleOne(lines)
	puzzleTwo(lines)
}

func puzzleOne(lines []string) {
	count := 0
	for _, line := range lines {
		fmt.Printf("line %s\n", line)
	}
	fmt.Printf("PuzzleOne %d\n", count)
}

func puzzleTwo(lines []string) {
	count := 0
	for _, line := range lines {
		fmt.Printf("line %s\n", line)
	}
	fmt.Printf("PuzzleTwo %d\n", count)
}
