package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input1 := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	input2 := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	lines1 := strings.Split(input1, "\n")
	lines2 := strings.Split(input2, "\n")
	puzzleOne(lines1)
	puzzleTwo(lines2)

	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("Fail to read file input.txt")
	}
	fileContext := string(content)
	lines := strings.Split(fileContext, "\n")
	puzzleOne(lines)
	puzzleTwo(lines)
}

func puzzleOne(lines []string) {
	count := 0
	first, last := "", ""

	for _, line := range lines {

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}
		}
		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(fmt.Sprintf("Bad input, line %s, first %s, last %s\n", line, first, last))
		}
		count += num
	}
	fmt.Printf("PuzzleOne %d\n", count)
}

func puzzleTwo(lines []string) {
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	count := 0

	numWordLengthArr := []int{3, 4, 5}

	for _, line := range lines {
		first, last := "", ""

	lineLoop:
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = string(line[i])
				break
			}

			// if i+3 >= len(line) {
			// 	continue
			// }
			// if numMap[line[i:i+3]] != 0 {
			// 	first = strconv.Itoa(numMap[line[i:i+3]])
			// 	break
			// }

			// if i+4 >= len(line) {
			// 	continue
			// }
			// if numMap[line[i:i+4]] != 0 {
			// 	first = strconv.Itoa(numMap[line[i:i+4]])
			// 	break
			// }

			// if i+5 >= len(line) {
			// 	continue
			// }
			// if numMap[line[i:i+5]] != 0 {
			// 	first = strconv.Itoa(numMap[line[i:i+5]])
			// 	break
			// }
			for _, numWordLength := range numWordLengthArr {
				if i+numWordLength >= len(line) {
					continue lineLoop
				}
				if numMap[line[i:i+numWordLength]] != 0 {
					first = strconv.Itoa(numMap[line[i:i+numWordLength]])
					break lineLoop
				}
			}
		}

	lineLoop2:
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = string(line[i])
				break
			}

			for _, numWordLength := range numWordLengthArr {

				if i-numWordLength+1 < 0 {
					continue lineLoop2
				}
				if numMap[line[i-numWordLength+1:i+1]] != 0 {
					last = strconv.Itoa(numMap[line[i-numWordLength+1:i+1]])
					break lineLoop2
				}
			}

			// if i-2 < 0 {
			// 	continue
			// }
			// if numMap[line[i-2:i+1]] != 0 {
			// 	last = strconv.Itoa(numMap[line[i-2:i+1]])
			// 	break
			// }

			// if i-3 < 0 {
			// 	continue
			// }
			// if numMap[line[i-3:i+1]] != 0 {
			// 	last = strconv.Itoa(numMap[line[i-3:i+1]])
			// 	break
			// }

			// if i-4 < 0 {
			// 	continue
			// }
			// if numMap[line[i-4:i+1]] != 0 {
			// 	last = strconv.Itoa(numMap[line[i-4:i+1]])
			// 	break
			// }
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(fmt.Sprintf("Bad input, line %s, first %s, last %s\n", line, first, last))
		}
		count += num
	}
	fmt.Printf("PuzzleTwo %d\n", count)
}
