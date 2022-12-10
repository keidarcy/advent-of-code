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
	fmt.Println("part2 result: ", part2)
}

type Point struct {
	x, y int
}

func getPart1(lines string) int {
	head := Point{0, 0}
	tail := Point{0, 0}
	tailVisitedMap := make(map[Point]bool)
	tailVisitedMap[tail] = true

	for _, l := range strings.Split(lines, "\n") {
		part := strings.Split(l, " ")
		direction := part[0]
		step, err := strconv.Atoi(part[1])
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}

		for step > 0 {
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}
			tail = computeTail1(head, tail)
			tailVisitedMap[tail] = true
			step--
		}
	}
	return len(tailVisitedMap)
}

func getPart2(lines string) int {
	knots := make([]Point, 10)
	tailVisitedMap := make(map[Point]bool)
	tailVisitedMap[knots[9]] = true

	for _, l := range strings.Split(lines, "\n") {
		part := strings.Split(l, " ")
		direction := part[0]
		step, err := strconv.Atoi(part[1])
		if err != nil {
			log.Fatal("BAD INT INPUT")
		}

		for step > 0 {
			switch direction {
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			}

			for i := 0; i < len(knots)-1; i++ {
				knots[i+1] = computeKnot(knots[i], knots[i+1])
			}
			tailVisitedMap[knots[9]] = true
			step--
		}
	}
	return len(tailVisitedMap)
}

func computeKnot(head, tail Point) Point {
	newTail := tail

	diff := Point{head.x - tail.x, head.y - tail.y}

	switch diff {
	case Point{2, 1}, Point{1, 2}, Point{0, 2}, Point{-1, 2}, Point{-2, 1}, Point{2, 2}, Point{-2, 2}:
		newTail.y++
	}

	switch diff {
	case Point{2, -1}, Point{1, -2}, Point{0, -2}, Point{-1, -2}, Point{-2, -1}, Point{2, -2}, Point{-2, -2}:
		newTail.y--
	}

	switch diff {
	case Point{1, 2}, Point{2, 1}, Point{2, 0}, Point{2, -1}, Point{1, -2}, Point{2, 2}, Point{2, -2}:
		newTail.x++
	}

	switch diff {
	case Point{-1, 2}, Point{-2, 1}, Point{-2, 0}, Point{-2, -1}, Point{-1, -2}, Point{-2, 2}, Point{-2, -2}:
		newTail.x--
	}

	return newTail
}

func computeTail1(head, tail Point) Point {
	newTail := tail

	diff := Point{head.x - tail.x, head.y - tail.y}

	switch diff {
	case Point{2, 1}, Point{1, 2}, Point{0, 2}, Point{-1, 2}, Point{-2, 1}:
		newTail.y++
	}

	switch diff {
	case Point{2, -1}, Point{1, -2}, Point{0, -2}, Point{-1, -2}, Point{-2, -1}:
		newTail.y--
	}

	switch diff {
	case Point{1, 2}, Point{2, 1}, Point{2, 0}, Point{2, -1}, Point{1, -2}:
		newTail.x++
	}

	switch diff {
	case Point{-1, 2}, Point{-2, 1}, Point{-2, 0}, Point{-2, -1}, Point{-1, -2}:
		newTail.x--
	}

	return newTail
}

// func getPart2(lines string) int {
// 	result := 0
// 	return result
// }

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
