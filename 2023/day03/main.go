package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input1 := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
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
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			digitLen := 0
			if unicode.IsDigit(rune(lines[y][x])) {
				for {
					digitLen++
					if x+digitLen == len(lines[0]) {
						break
					}
					if !unicode.IsDigit(rune(lines[y][x+digitLen])) {
						break
					}
				}
				if !notTarget(lines, x, digitLen, y) {
					n, err := strconv.Atoi(lines[y][x : x+digitLen])
					if err != nil {
						panic(fmt.Sprintf("bad input %s\n", line))
					}
					count += n
				}
				x += digitLen
			}
		}
	}
	fmt.Printf("PuzzleOne %d\n", count)
}

type point struct {
	x int
	y int
}

func puzzleTwo(lines []string) {
	count := 0
	pairMap := map[point]bool{}
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			digitLen := 0
			if unicode.IsDigit(rune(lines[y][x])) {
				for {
					digitLen++
					if x+digitLen == len(lines[0]) {
						break
					}
					if !unicode.IsDigit(rune(lines[y][x+digitLen])) {
						break
					}
				}
				n, err := strconv.Atoi(lines[y][x : x+digitLen])
				if err != nil {
					panic(fmt.Sprintf("bad input %s\n", line))
				}
				if found, pair := getPair(lines, x, digitLen, y, pairMap); found {
					count += n * pair
				}
				x += digitLen
			}
		}
	}
	fmt.Printf("PuzzleTwo %d\n", count)
}

func getPair(lines []string, x, digitLen, y int, pairMap map[point]bool) (bool, int) {
	// if in map return false
	if _, ok := pairMap[point{x: x + digitLen, y: y}]; ok {
		return false, 0
	}

	// right
	if x+digitLen < len(lines[0])-1 && rune(lines[y][x+digitLen]) == '*' {
		if unicode.IsDigit(rune(lines[y][x+digitLen+1])) {
			pairDigitLen := 0
			for {
				pairDigitLen++
				if x+digitLen+1+pairDigitLen == len(lines[0]) {
					break
				}
				if !unicode.IsDigit(rune(lines[y][x+digitLen+1+pairDigitLen])) {
					break
				}
			}
			n, err := strconv.Atoi(lines[y][x+digitLen+1 : x+digitLen+1+pairDigitLen])
			if err != nil {
				panic(fmt.Sprintf("bad input %s\n", lines[y]))
			}
			return true, n
		}
		return false, 0
	}

	// bottom
	left := 0
	if x == 0 {
		left = x
	} else {
		left = x - 1
	}
	right := 0
	if x+digitLen+1 == len(lines[0]) {
		right = x + digitLen + 1
	} else {
		right = x + digitLen + 2
	}
	for i := left; i < right; i++ {
		if y < len(lines)-1 && rune(lines[y+1][i]) == '*' {
			if unicode.IsDigit(rune(lines[y+1][i-1])) {

				pairDigitLen := 0
				for {
					pairDigitLen++
					if x+digitLen+1+pairDigitLen == len(lines[0]) {
						break
					}
					if !unicode.IsDigit(rune(lines[y][x+digitLen+1+pairDigitLen])) {
						break
					}
				}
				n, err := strconv.Atoi(lines[y][x+digitLen+1 : x+digitLen+1+pairDigitLen])
				if err != nil {
					panic(fmt.Sprintf("bad input %s\n", lines[y]))
				}
				return true, n
			}
			return false, 0
		}
	}
	return false, 0
}

func notTarget(lines []string, x, digitLen, y int) bool {
	for i := x; i < x+digitLen; i++ {
		// top
		if y > 0 && rune(lines[y-1][i]) != '.' {
			return false
		}
		// bottom
		if y < len(lines)-1 && rune(lines[y+1][i]) != '.' {
			return false
		}
	}

	// top left
	if y > 0 && x > 0 && rune(lines[y-1][x-1]) != '.' {
		return false
	}
	// left
	if x > 0 && rune(lines[y][x-1]) != '.' {
		return false
	}
	// bottom left
	if y < len(lines)-1 && x > 0 && rune(lines[y+1][x-1]) != '.' {
		return false
	}

	// top right
	if y > 0 && x+digitLen < len(lines[0])-1 && rune(lines[y-1][x+digitLen]) != '.' {
		return false
	}
	// right
	if x+digitLen < len(lines[0])-1 && rune(lines[y][x+digitLen]) != '.' {
		return false
	}
	// bottom right
	if y < len(lines)-1 && x+digitLen < len(lines[0])-1 && rune(lines[y+1][x+digitLen]) != '.' {
		return false
	}
	return true
}
