package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := getInput()
	part1 := getPart1(input)
	// part2 := getPart2(input)
	fmt.Println("part1 result: ", part1)
	// fmt.Println("part2 result: ", part2)
}

func getPart1(input string) int {
	// lines := strings.Split(input, "\n")

	return 1
}

type point struct {
	x, y int
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
