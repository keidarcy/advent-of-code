package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	input := getInput()
	part1 := getPart1(input, 4)
	part2 := getPart1(input, 14)

	fmt.Println("part 1: ", part1)
	fmt.Println("part 2: ", part2)
}

func getPart1(input string, n int) int {
	return findUniqueNWindow(input, n)
}

func findUniqueNWindow(str string, n int) int {
	left := 0
	right := n
	for right < len(str) {

		part := str[left:right]
		if isUnique(part, n) {
			return right
		}
		left++
		right++
	}
	panic("BAD STR INPUT")
}

func isUnique(str string, n int) bool {
	m := make(map[string]int)

	for i, c := range str {
		m[string(c)] = i
	}

	return len(m) == n
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
