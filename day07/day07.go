package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	part1 := getPart1(input)
	// part2 := getPart2(input)
	fmt.Println("part1", part1)
	// fmt.Println("part2", part2)
}

type FileNode struct {
}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		fmt.Println(l)
	}
	return 1
}

// func getPart2(input string) int {
// 	lines := strings.Split(input, "\n")
// 	count := 0
// 	for _, l := range lines {
// 		if parseLine(l, false) {
// 			count++
// 		}
// 	}
// 	return count
// }

func parseLine(line string) bool {
	fmt.Println(line)
	return true
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
