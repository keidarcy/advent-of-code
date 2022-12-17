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
	fmt.Println("part1 result: ", part1)
	// fmt.Println("part2 result: ", part2)
}

type node struct {
	name     string
	size     int
	isFile   bool
	children map[string]*node
	parent   *node
}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")

	// dirs := []*node{}
	var currentDir *node

	for _, line := range lines {
		commands := strings.Split(line, " ")

		// $ cd a
		if len(commands) > 2 {
			if commands[2] == ".." {
				currentDir = currentDir.parent
			} else if commands[2] == "/" {
				currentDir = &node{"/", 0, false, make(map[string]*node), nil}
			} else {
				currentDir = currentDir.children[commands[2]]
			}
		}

	}

	return 1
}
func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
