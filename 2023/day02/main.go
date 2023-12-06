package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input1 := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	lines1 := strings.Split(input1, "\n")
	puzzleOne(lines1)
	puzzleTwo(lines1)

	content, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("fail to read file input.txt")
	}
	lines := strings.Split(string(content), "\n")
	puzzleOne(lines)
	puzzleTwo(lines)

}

func puzzleTwo(lines []string) {
	count := 0
	red, green, blue := 0, 0, 0
	for _, line := range lines {
		redSlice, greenSlice, blueSlice := []int{}, []int{}, []int{}
		parts := strings.Split(line, ":")
		gameStr := parts[1]
		games := strings.Split(gameStr, ";")
		for _, game := range games {
			cards := strings.Split(game, ",")
			for _, card := range cards {
				colorStartsFrom := 0
				for i, s := range card[1:] {
					if s == ' ' {
						colorStartsFrom = i + 2
					}
				}
				colorStr := card[1 : colorStartsFrom-1]
				colorCount, err := strconv.Atoi(colorStr)
				if err != nil {
					panic(fmt.Sprintf("line: %s, colorStartsFrom: %d, bad input", line, colorStartsFrom))
				}
				switch card[colorStartsFrom:] {
				case "red":
					redSlice = append(redSlice, colorCount)
				case "green":
					greenSlice = append(greenSlice, colorCount)
				case "blue":
					blueSlice = append(blueSlice, colorCount)
				}
			}
		}
		red = findMax(redSlice)
		green = findMax(greenSlice)
		blue = findMax(blueSlice)
		// fmt.Printf("redSlice %v, red %d\n", redSlice, red)
		// fmt.Printf("greenSlice %v, green %d\n", greenSlice, green)
		// fmt.Printf("blueSlice %v, blue %d\n", blueSlice, blue)
		// fmt.Printf("multiplier %d\n", red*green*blue)
		count += red * green * blue
	}
	fmt.Printf("Puzzle 2: %d\n", count)
}

func findMax(numbers []int) int {
	maxNum := numbers[0]
	for _, num := range numbers {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func puzzleOne(lines []string) {
	count := 0
	redLimit := 12
	greenLimit := 13
	blueLimit := 14
lineLoop:
	for lineIndex, line := range lines {
		goodGame := true
		parts := strings.Split(line, ":")
		gameStr := parts[1]
		games := strings.Split(gameStr, ";")
		for _, game := range games {
			cards := strings.Split(game, ",")
			for _, card := range cards {
				colorStartsFrom := 0
				for i, s := range card[1:] {
					if s == ' ' {
						colorStartsFrom = i + 2
					}

				}
				colorStr := card[1 : colorStartsFrom-1]
				colorCount, err := strconv.Atoi(colorStr)
				if err != nil {
					panic(fmt.Sprintf("line: %s, colorStartsFrom: %d, bad input", line, colorStartsFrom))
				}
				switch card[colorStartsFrom:] {
				case "red":
					if colorCount > redLimit {
						goodGame = false
						continue lineLoop
					}
				case "green":
					if colorCount > greenLimit {
						goodGame = false
						continue lineLoop
					}
				case "blue":
					if colorCount > blueLimit {
						goodGame = false
						continue lineLoop
					}
				}
			}
		}
		if goodGame {
			count += lineIndex + 1
		}
	}
	fmt.Printf("Puzzle 1: %d\n", count)
}
