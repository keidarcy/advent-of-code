package main

import (
	"fmt"
	"log"
	"math"
	"os"
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

func getPart1(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, 0)
	for _, line := range lines {
		// line string to []byte
		line := []byte(line)
		grid = append(grid, line)
	}

	var start, end point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				start = point{x: i, y: j}
				grid[i][j] = 'a'
			}
			if grid[i][j] == 'E' {
				end = point{x: i, y: j}
				grid[i][j] = 'z'
			}
		}
	}
	result := bfs1(grid, start, end)
	return result
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, 0)
	for _, line := range lines {
		// line string to []byte
		line := []byte(line)
		grid = append(grid, line)
	}

	var end point
	var starts []point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = 'a'
			}
			if grid[i][j] == 'a' {
				starts = append(starts, point{x: i, y: j})
			}
			if grid[i][j] == 'E' {
				end = point{x: i, y: j}
				grid[i][j] = 'z'
			}
		}
	}
	minResult := math.MaxInt
	for _, start := range starts {
		result := bfs2(grid, start, end)
		if minResult > result {
			minResult = result
		}
	}
	return minResult
}

func bfs1(grid [][]byte, start point, end point) int {
	queue := []point{start}
	visitedPos := make(map[point]bool)
	visitedPos[start] = true
	steps := 0
	dirs := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	width := len(grid)
	length := len(grid[0])

	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == end {
				return steps
			}

			for _, dir := range dirs {
				next := point{x: dir.x + current.x, y: dir.y + current.y}
				// out of range
				if next.x < 0 || next.x >= width || next.y < 0 || next.y >= length {
					continue
				}

				// is visited
				if visitedPos[next] {
					continue
				}

				// is too high
				if grid[next.x][next.y]-grid[current.x][current.y] > 1 && grid[next.x][next.y] > grid[current.x][current.y] {
					continue
				}

				queue = append(queue, next)
				visitedPos[next] = true
			}
		}
		steps++
	}

	return math.MaxInt

}

func bfs2(grid [][]byte, start point, end point) int {
	queue := []point{start}
	visitedPos := make(map[point]bool)
	visitedPos[start] = true
	steps := 0
	dirs := []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	width := len(grid)
	length := len(grid[0])

	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == end {
				return steps
			}

			for _, dir := range dirs {
				next := point{x: dir.x + current.x, y: dir.y + current.y}

				// if next.x < 0 || next.x >= width || next.y < 0 || next.y >= length || grid[next.x][next.y]-grid[current.x][current.y] > 1 || visitedPos[next] {
				// 	continue
				// }

				// out of range
				if next.x < 0 || next.x >= width || next.y < 0 || next.y >= length {
					continue
				}

				// is visited
				if visitedPos[next] {
					continue
				}

				// is too high
				if grid[next.x][next.y]-grid[current.x][current.y] > 1 && grid[next.x][next.y] > grid[current.x][current.y] {
					continue
				}

				queue = append(queue, next)
				visitedPos[next] = true
			}
		}
		steps++
	}

	return math.MaxInt

}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
