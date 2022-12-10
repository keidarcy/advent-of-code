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
	getPart2(input)
	fmt.Println("part1 result: ", part1)
}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")

	cycle := 0
	x := 1
	signal := 0
	for _, l := range lines {
		line := parseLine1(l)
		// add cycle by command but add x end of command
		for i := 0; i < line.c; i++ {
			cycle++
			// 20, 60, 100
			if (cycle-20)%40 == 0 {
				signal += cycle * x
				fmt.Printf("cycle: %d, x: %d, signal: %d\n", cycle, x, signal)
			}
		}
		x += line.x

	}

	return signal
}

type Line struct {
	c, x int
}

func parseLine1(l string) Line {
	if rune(l[0]) == 'n' {
		return Line{c: 1, x: 0}
	}
	parts := strings.Split(l, " ")
	x, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatal("Bad integer input")
	}

	return Line{c: 2, x: x}
}

func getPart2(input string) {
	lines := strings.Split(input, "\n")
	images := [][]string{}
	image := []string{}

	// crd = cycle - 1
	cycle := 0
	x := 1
	for _, l := range lines {
		line := parseLine1(l)
		// add cycle by command but add x end of command
		for i := 0; i < line.c; i++ {
			if cycle%40 == 0 {
				images = append(images, image)
				image = []string{}
			}
			cycle++
			image = append(image, ".")
			handleImage(image, x)
		}
		x += line.x

	}
	images = append(images, image)

	// print it
	for _, i := range images[1:] {
		fmt.Println(i)
	}
}

func handleImage(image []string, x int) {
	current := len(image) - 1
	if x-1 == current || x == current || x+1 == current {
		image[len(image)-1] = "#"
	}
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
