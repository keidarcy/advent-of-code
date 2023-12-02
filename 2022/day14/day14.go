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

type point struct {
	x, y int
}

type cave struct {
	content map[point]rune
	maxX    int
	maxY    int
}

func (c *cave) countSandNoFloor() int {
	intoVoid := false
	sand := 0

	for !intoVoid {
		newSand := point{500, 0}

		for {
			c.content[newSand] = '"'
			if newSand.y+1 > c.maxY {
				intoVoid = true
				break
			}
			if c.content[point{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if c.content[point{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.x--
				newSand.y++
			} else if c.content[point{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.x++
				newSand.y++
			} else {
				c.content[newSand] = 'o'
				sand++
				break
			}
		}
	}

	return sand
}

func (c *cave) countSandHasFloor() int {
	sand := 0

	for {
		newSand := point{500, 0}
		if c.content[newSand] == 'o' {
			break
		}

		for {
			c.content[newSand] = '"'
			if newSand.y+1 > c.maxY {
				break
			}
			if c.content[point{newSand.x, newSand.y + 1}] < '#' {
				newSand.y++
			} else if c.content[point{newSand.x - 1, newSand.y + 1}] < '#' {
				newSand.x--
				newSand.y++
			} else if c.content[point{newSand.x + 1, newSand.y + 1}] < '#' {
				newSand.x++
				newSand.y++
			} else {
				c.content[newSand] = 'o'
				sand++
				break
			}
		}
	}

	return sand
}

func (c cave) showCave() {
	for y := 0; y <= c.maxY+1; y++ {
		for x := 460; x <= c.maxX+1; x++ {
			if c.content[point{x, y}] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(c.content[point{x, y}]))
			}
		}
		fmt.Println()
	}
}

func (c *cave) addCaveFloor() {
	c.maxY += 2
	c.maxX += 500
	for x := 0; x <= c.maxX; x++ {
		c.content[point{x: x, y: c.maxY}] = '#'
	}

}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")

	endlessCave := parseCave(lines)
	sand := endlessCave.countSandNoFloor()

	endlessCave.showCave()

	return sand
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")

	newCave := parseCave(lines)
	newCave.addCaveFloor()
	sand := newCave.countSandHasFloor()

	// newCave.showCave()

	return sand
}

func parseCave(lines []string) cave {
	var maxX, maxY int
	content := make(map[point]rune)
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for i := range points[:len(points)-1] {

			from := strings.Split(points[i], ",")
			to := strings.Split(points[i+1], ",")

			fromX, _ := strconv.Atoi(from[0])
			fromY, _ := strconv.Atoi(from[1])
			toX, _ := strconv.Atoi(to[0])
			toY, _ := strconv.Atoi(to[1])

			content[point{fromX, fromY}] = '#'
			content[point{toX, toY}] = '#'

			if toX > maxX {
				maxX = toX
			}

			if toY > maxY {
				maxX = toY
			}

			for fromX != toX || fromY != toY {
				content[point{fromX, fromY}] = '#'
				switch {
				case toX > fromX:
					fromX++
				case toX < fromX:
					fromX--
				case toY > fromY:
					fromY++
				case toY < fromY:
					fromY--
				}
			}
			if fromX > maxX {
				maxX = fromX
			}
			if fromY > maxY {
				maxY = fromY
			}
		}
	}
	return cave{
		content: content,
		maxX:    maxX,
		maxY:    maxY,
	}
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
