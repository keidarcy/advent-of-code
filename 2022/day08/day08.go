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
	part1 := getPart1(input)
	part2 := getPart2(input)
	fmt.Println("part1 result: ", part1)
	fmt.Println("part2", part2)
}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	row := len(lines)
	column := len(lines[0])
	for i, l := range lines {
		for j, c := range l {
			// left
			if i == 0 || i == row-1 || j == 0 || j == column-1 {
				result++
				continue
			}
			height, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal("BAD INT INPUT")
			}
			if isVisible(lines, height, i, j) {
				result++
			}
		}
	}
	return result
}

func isVisible(lines []string, height, i, j int) bool {

	// left
	result := true
	for _, c := range lines[i][:j] {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		if height <= num {
			result = false
			continue
		}
	}

	if result {
		return result
	}

	// right
	result = true
	for _, c := range lines[i][j+1:] {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		if height <= num {
			result = false
			continue
		}
	}

	if result {
		return result
	}

	// up
	result = true
	for row, l := range lines {
		if row == i {
			break
		}

		c := l[j]
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		if height <= num {
			result = false
			continue
		}
	}

	if result {
		return result
	}

	// down
	result = true
	for row, l := range lines {
		if row <= i {
			continue
		}
		c := l[j]
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		if height <= num {
			result = false
			continue
		}
	}

	return result
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")
	max := 0
	row := len(lines)
	column := len(lines[0])
	for i, l := range lines {
		for j, c := range l {
			// left
			if i == 0 || i == row-1 || j == 0 || j == column-1 {
				continue
			}
			height, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal("BAD INT INPUT")
			}
			result := countView(lines, height, i, j)
			if max < result {
				max = result
			}
		}
	}
	return max
}

func countView(lines []string, height, i, j int) int {

	// left
	left := 0
	for index := j - 1; index >= 0; index-- {

		num, err := strconv.Atoi(string(lines[i][index]))

		if err != nil {
			log.Fatal("BAD INT INPUT")
		}

		left++
		if height <= num {
			break
		}
	}

	// right
	right := 0
	for _, c := range lines[i][j+1:] {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		right++
		if height <= num {
			break
		}
	}

	// up
	up := 0
	var arr []int
	for row, l := range lines {
		if row == i {
			break
		}
		num, err := strconv.Atoi(string(l[j]))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		arr = append(arr, num)
	}

	for index := len(arr) - 1; index >= 0; index-- {
		num := arr[index]
		up++
		if height <= num {
			break
		}
	}

	// down
	down := 0
	for row, l := range lines {
		if row <= i {
			continue
		}
		c := l[j]
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}
		down++
		if height <= num {
			break
		}
	}

	return left * right * up * down
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
