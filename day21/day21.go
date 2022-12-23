package main

import (
	"fmt"
	"log"
	"math"
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

func getPart1(input string) int {
	lines := strings.Split(input, "\n")
	monkeys := make(map[string]string)
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		monkeys[parts[0]] = parts[1]
	}
	result := count(monkeys, "root")
	return int(result)
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")
	monkeys := make(map[string]string)
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		monkeys[parts[0]] = parts[1]
	}

	parts := strings.Split(monkeys["root"], " ")
	left := count(monkeys, parts[0])
	right := count(monkeys, parts[2])
	var min = 0
	var max = math.MaxInt64
	result := 0

	for left != right {
		mid := (min + max) / 2
		monkeys["humn"] = fmt.Sprintf("%d", mid)
		left = count(monkeys, parts[0])
		if left > right {
			min = mid + 1
		}
		if left < right {
			max = mid
		}
		if left == right {
			result = mid
			break
		}
	}
	return result
}

func count(monkeys map[string]string, name string) float64 {
	parts := strings.Split(monkeys[name], " ")

	if len(parts) == 1 {
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("BAD integer INPUT ")
		}
		return float64(val)
	}

	switch parts[1] {
	case "+":
		return count(monkeys, parts[0]) + count(monkeys, parts[2])
	case "-":
		return count(monkeys, parts[0]) - count(monkeys, parts[2])
	case "*":
		return count(monkeys, parts[0]) * count(monkeys, parts[2])
	case "/":
		return count(monkeys, parts[0]) / count(monkeys, parts[2])
	}
	log.Fatal("SHOULD NOT COME HERE")
	return -1
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
