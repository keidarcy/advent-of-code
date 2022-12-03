package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	lines := strings.Split(getInput(), "\n")
	result := getPriority2(lines)
	fmt.Println("Result: ", result)
}

func getPriority2(lines []string) int {
	count := 0
	items := make(map[rune][]int)
	sum := 0

	// lineLoop:
	for _, l := range lines {
		for _, r := range l {
			value, ok := items[r]
			if ok {
				if items[r][0] == count {
					continue
				}
				if count == 1 && value[1] == 1 {
					items[r][1] = 2
				} else if count == 2 && value[1] == 2 {
					items[r][1] = 3
					sum += getNum(r)
					break
				}
			} else {
				items[r] = append(items[r], count)
				items[r] = append(items[r], 1)
			}
		}
		count++
		if count == 3 {
			count = 0
			items = make(map[rune][]int)
		}
	}
	return sum
}

// func getPriority1(lines []string) int {

// 	sum := 0
// 	for _, l := range lines {
// 		sum += parseLine(l)
// 	}

// 	return sum

// }

// func parseLine(l string) int {
// 	halfLength := len(l) / 2

// 	items := make(map[rune]bool)

// 	for _, leftPart := range l[:halfLength] {
// 		items[leftPart] = true
// 	}

// 	for _, rightPart := range l[halfLength:] {
// 		if items[rightPart] {
// 			return getNum(rightPart)
// 		}
// 	}

// 	return 0
// }

func getNum(r rune) int {
	if unicode.IsUpper(r) {
		return int(unicode.ToLower(r) - 96 + 26)
	}
	return int(r - 96)
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}
