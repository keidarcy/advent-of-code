package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type point struct {
	x, y, z int
}

func main() {
	input := getInput()
	part1 := getPart1(input)
	part2 := getPart2(input)
	fmt.Println("part1 result: ", part1)
	fmt.Println("part2 result: ", part2)
}

func getPart1(input string) int {
	lines := strings.Split(input, "\n")
	points := make(map[point]bool)

	for _, l := range lines {
		var p point
		fmt.Sscanf(l, "%d,%d,%d", &p.x, &p.y, &p.z)
		points[p] = true
	}

	count := 0

	delta := []point{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

	for p := range points {
		for _, d := range delta {
			if !points[p.add(d)] {
				count++
			}
		}
	}

	return count
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")

	points := make(map[point]bool)
	min := point{math.MaxInt, math.MaxInt, math.MaxInt}
	max := point{math.MinInt, math.MinInt, math.MinInt}

	for _, l := range lines {
		var p point
		fmt.Sscanf(l, "%d,%d,%d", &p.x, &p.y, &p.z)

		min = point{Min(min.x, p.x), Min(min.y, p.y), Min(min.z, p.z)}
		max = point{Max(max.x, p.x), Max(max.y, p.y), Max(max.z, p.z)}
		points[p] = true
	}

	queue := []point{min}
	visited := map[point]struct{}{min: {}}
	delta := []point{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

	for x := min.x - 1; x <= max.x+1; x++ {
		for y := min.y - 1; y <= max.y+1; y++ {
			for z := min.z - 1; z <= max.z+1; z++ {
				points[point{x, y, z}] = points[point{x, y, z}]
			}
		}
	}

	count := 0

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range delta {
			next := cur.add(d)

			if cube, valid := points[next]; cube {
				count++
			} else if _, seen := visited[next]; valid && !seen {
				visited[next] = struct{}{}
				queue = append(queue, next)
			}
		}
	}

	return count
}

func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y, p.z + q.z}
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
