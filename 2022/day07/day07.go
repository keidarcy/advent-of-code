package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	part1 := getPart1(input)
	part2 := getPart2(input)
	fmt.Println("part1 result: ", part1)
	fmt.Println("part2 result: ", part2)
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

	dirs := []*node{}
	var currentDir *node

	for _, line := range lines {
		if line == "" {
			continue
		}
		command := strings.Split(line, " ")

		// $ cd a
		if len(command) > 2 {
			if command[2] == ".." {
				currentDir = currentDir.parent
			} else if command[2] == "/" {
				currentDir = &node{"/", 0, false, make(map[string]*node), nil}
			} else {
				currentDir = currentDir.children[command[2]]
			}
		}

		if rune(command[0][0]) == '$' {
			// $ ls
			continue
		} else if rune(command[0][0]) == 'd' {
			// dir a
			currentDir.children[command[1]] = &node{command[1], 0, false, make(map[string]*node), currentDir}
			dirs = append(dirs, currentDir.children[command[1]])
		} else {
			// 2557 g
			size, _ := strconv.Atoi(command[0])
			currentDir.children[command[1]] = &node{command[1], size, true, nil, currentDir}
		}
	}

	result := 0
	for _, dir := range dirs {
		size := countSize(*dir)
		if size <= 100000 {
			result += size
		}
	}
	return result
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")

	dirs := []*node{}
	var currentDir *node

	for _, line := range lines {
		if line == "" {
			continue
		}
		command := strings.Split(line, " ")

		// $ cd a
		if len(command) > 2 {
			if command[2] == ".." {
				currentDir = currentDir.parent
			} else if command[2] == "/" {
				currentDir = &node{"/", 0, false, make(map[string]*node), nil}
				dirs = append(dirs, currentDir)
			} else {
				currentDir = currentDir.children[command[2]]
			}
		}

		if rune(command[0][0]) == '$' {
			// $ ls
			continue
		} else if rune(command[0][0]) == 'd' {
			// dir a
			currentDir.children[command[1]] = &node{command[1], 0, false, make(map[string]*node), currentDir}
			dirs = append(dirs, currentDir.children[command[1]])
		} else {
			// 2557 g
			size, _ := strconv.Atoi(command[0])
			currentDir.children[command[1]] = &node{command[1], size, true, nil, currentDir}
		}
	}

	totalSize := 0
	for _, dir := range dirs {
		size := countSize(*dir)
		dir.size = size
		if size > totalSize {
			totalSize = size
		}
	}
	targetSize := 30000000 - (70000000 - totalSize)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	for _, dir := range dirs {
		if dir.size > targetSize {
			return dir.size
		}

	}
	return -1
}

func countSize(n node) int {

	if n.isFile {
		return n.size
	}
	size := 0
	for _, childNode := range n.children {
		size += countSize(*childNode)
	}
	return size
}
func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
