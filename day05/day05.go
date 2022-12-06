package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/goutil/dump"
)

func main() {
	input := getInput()
	oldTop := getTop(input, false)
	newTop := getTop(input, true)
	// overlap := getOther(input)
	dump.P("old top", oldTop)
	dump.P("new top", newTop)
}

func getTop(input string, isNew bool) string {
	lines := []string{strings.Repeat(" ", 100)}
	lines = append(lines, strings.Split(input, "\n")...)

	stack := make([][]string, 0)

	n := (len(lines[4]) + 1) / 4

	for i, l := range lines {
		if l == "" {
			break
		}
		if len(l) < n*4 {
			lines[i] += strings.Repeat(" ", n*4-len(l))
		}
	}
	// build stack
	for row, l := range lines {

		// hit number line break
		if string(l[1]) == "1" {
			break
		}

		// init row
		stack = append(stack, make([]string, 0))

		// loop line build stack
		for i := 0; i < n; i++ {
			if string(l[i*4]) == "[" {
				stack[row] = append(stack[row], string(l[i*4+1]))
			} else {
				stack[row] = append(stack[row], " ")
			}
		}

	}
	rotateMatrix(n, stack)
	dropEmpty(n, stack)

	// move stack
	for _, l := range lines[n+2:] {
		move, from, to := parseLine(l)
		fromLen := len(stack[from])
		cut := stack[from][fromLen-move:]
		stack[from] = stack[from][:fromLen-move]
		// reverse cut
		if !isNew {

			for i, j := 0, len(cut)-1; i < j; i, j = i+1, j-1 {
				cut[i], cut[j] = cut[j], cut[i]
			}
		}
		stack[to] = append(stack[to], cut...)
	}

	result := ""

	// dump.P(stack)
	for i := 0; i < n; i++ {
		result += stack[i][len(stack[i])-1]
	}

	return result
}

func dropEmpty(n int, matrix [][]string) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == " " {
				matrix[i] = matrix[i][:j]
				break
			}
		}
	}
}

func parseLine(line string) (int, int, int) {
	l := strings.Split(line, " ")
	move, err := strconv.Atoi(l[1])
	if err != nil {
		log.Fatal("BAD INPUT")
	}
	from, err := strconv.Atoi(l[3])
	if err != nil {
		log.Fatal("BAD INPUT")
	}
	to, err := strconv.Atoi(l[5])
	if err != nil {
		log.Fatal("BAD INPUT")
	}
	return move, from - 1, to - 1
}

func rotateMatrix(n int, matrix [][]string) {
	for i := 0; i < n/2+1; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}

// 0, D, 0
// N, C, 0
// Z, M ,P

// Z, N, 0
// M, C, D
// P, 0, 0

// 1, 2, 3
// 4, 5, 6
// 7, 8, 9

// 7, 4, 1
// 8, 5, 2
// 9, 6, 3

// func parseLine(line string) bool {
// 	return false
// }

// func getOther(input string) string {
// 	return ""
// }

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
