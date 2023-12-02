package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	result := calculateScore(input)
	fmt.Println("Result: ", result)
}

func getInput() string {
	// https://adventofcode.com/2022/day/2/input
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func calculateScore(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		item := strings.Split(line, " ")
		elf := item[0]
		// me := item[1]
		// total += countScoreByGame(elf, me)
		// total += countScoreByShape(me)

		need := item[1]
		score, me := countScoreByNeeds(elf, need)
		total += score
		total += countScoreByShape(me)
	}

	return total
}

// A for Rock, B for Paper, and C for Scissors.
// X for Rock, Y for Paper, and Z for Scissors.

// 1 for Rock
// 2 for Paper
// 3 for Scissors

// 0 if you lost
// 3 if the round was a draw
// 6 if you won

// X lose
// Y draw
// Z win
func countScoreByNeeds(opponent string, needShape string) (int, string) {
	opponentMap := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	needMap := map[string]string{"X": "lose", "Y": "draw", "Z": "win"}
	scoreMap := map[string]int{"X": 0, "Y": 3, "Z": 6}
	meShapeMap := map[string]string{"Rock": "X", "Paper": "Y", "Scissors": "Z"}

	need := needMap[needShape]
	opponentShape := opponentMap[opponent]

	meShape := ""

	if need == "draw" {
		meShape = opponentShape
	}

	if need == "lose" {
		if opponentShape == "Rock" {
			meShape = "Scissors"
		} else if opponentShape == "Paper" {
			meShape = "Rock"
		} else {
			meShape = "Paper"
		}
	}

	if need == "win" {
		if opponentShape == "Rock" {
			meShape = "Paper"
		} else if opponentShape == "Paper" {
			meShape = "Scissors"
		} else {
			meShape = "Rock"
		}
	}

	return scoreMap[needShape], meShapeMap[meShape]
}

// A Y
// B X
// C Z

func countScoreByGame(opponent string, me string) int {
	opponentMap := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors"}
	meMap := map[string]string{"X": "Rock", "Y": "Paper", "Z": "Scissors"}
	resultMap := map[string]int{"won": 6, "draw": 3, "lost": 0}

	opponentShape := opponentMap[opponent]
	meShape := meMap[me]

	result := "draw"
	if opponentShape == meShape {
		result = "draw"
		return resultMap[result]
	}

	if meShape == "Rock" {
		if opponentShape == "Paper" {
			result = "lost"
		} else {
			result = "won"
		}
	}

	if meShape == "Paper" {
		if opponentShape == "Scissors" {
			result = "lost"
		} else {
			result = "won"
		}
	}

	if meShape == "Scissors" {
		if opponentShape == "Rock" {
			result = "lost"
		} else {
			result = "won"
		}
	}

	return resultMap[result]
}

// 1 for Rock
// 2 for Paper
// 3 for Scissors
func countScoreByShape(shape string) int {
	m := map[string]int{"X": 1, "Y": 2, "Z": 3}
	return m[shape]
}
