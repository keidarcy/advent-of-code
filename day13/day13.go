package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func getPart1(input string) int {
	lines := strings.Split(input, "\n")
	pair := 1
	count := 0
	for i := 0; i < len(lines); i += 3 {
		var left, right any
		json.Unmarshal([]byte(lines[i]), &left)
		json.Unmarshal([]byte(lines[i+1]), &right)
		if compare(left, right) <= 0 {
			count += pair
		}
		pair++
	}
	return count
}

func getPart2(input string) int {
	lines := strings.Split(input, "\n")
	// packets := []any{[]any{[]any{2.}}, []any{[]any{6.}}}
	packets := []any{}
	for i := 0; i < len(lines); i += 3 {
		var left, right any
		json.Unmarshal([]byte(lines[i]), &left)
		json.Unmarshal([]byte(lines[i+1]), &right)
		packets = append(packets, left, right)
	}

	count := 1
	packets = append(packets, []any{[]any{2.}}, []any{[]any{6.}})
	// sort.Slice(packets, func(i, j int) bool { return compare(packets[i], packets[j]) < 0 })
	packets = bubbleSort(packets)
	for i, p := range packets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			count *= i + 1
		}
	}
	return count
}

func bubbleSort(arr []any) []any {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if compare(arr[j-1], arr[j]) > 0 {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr

}

func compare(left any, right any) int {

	leftValue, leftOk := left.([]any)
	rightValue, rightOk := right.([]any)

	if !leftOk && !rightOk {
		return int(left.(float64) - right.(float64))
	}

	if !leftOk {
		leftValue = []any{left}
	}

	if !rightOk {
		rightValue = []any{right}
	}

	for i := 0; i < len(leftValue) && i < len(rightValue); i++ {
		if c := compare(leftValue[i], rightValue[i]); c != 0 {
			return c
		}
	}
	return len(leftValue) - len(rightValue)
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
