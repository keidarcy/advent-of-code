package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	contain := countContain(input)
	overlap := countOverlap(input)
	fmt.Println("contain", contain)
	fmt.Println("overlap", overlap)
}

func countContain(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, l := range lines {
		if parseLine(l, true) {
			count++
		}
	}
	return count
}

func countOverlap(input string) int {
	lines := strings.Split(input, "\n")
	count := 0
	for _, l := range lines {
		if parseLine(l, false) {
			count++
		}
	}
	return count
}

func parseLine(line string, isContain bool) bool {
	assignments := strings.Split(line, ",")
	firstRange, secondRange := assignments[0], assignments[1]
	if isContain {

		return hasContain(firstRange, secondRange)
	}
	return hasOverlap(firstRange, secondRange)
}

func hasOverlap(first string, second string) bool {
	firstStart, firstEnd := parsePart(first)
	secondStart, secondEnd := parsePart(second)
	return (firstEnd >= secondStart && firstStart <= secondEnd) || (firstStart <= secondEnd && firstEnd >= secondStart)
}

func hasContain(first string, second string) bool {
	firstStart, firstEnd := parsePart(first)
	secondStart, secondEnd := parsePart(second)
	return (firstStart <= secondStart && firstEnd >= secondEnd) || (firstStart >= secondStart && firstEnd <= secondEnd)
}

func parsePart(part string) (int, int) {
	section := strings.Split(part, "-")
	start, end := section[0], section[1]

	startInt, err := strconv.Atoi(start)

	if err != nil {
		log.Fatal(err)
	}

	endInt, err := strconv.Atoi(end)

	if err != nil {
		log.Fatal(err)
	}

	return startInt, endInt
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
