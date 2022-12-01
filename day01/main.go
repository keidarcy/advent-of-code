package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	// https://adventofcode.com/2022/day/1/input
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}

func main() {
	lines := strings.Split(getInput(), "\n")
	// elf := make([]int, 0)
    // elf = append(elf, 0)

	// var elf []int
	// elf = append(elf, 0)

	elf := []int{0}
	// bigNumber := 0
	// currentNum := 0

	for _, l := range lines {

		if l == "" {
			elf = append(elf, 0)
			// if bigNumber < currentNum {
			// 	bigNumber = currentNum
			// }
			// currentNum = 0
			continue
		}

		i, err := strconv.Atoi(l)

		if err != nil {
			panic(err)
		}
		// currentNum += i

		currentIndex := len(elf) - 1
		currentValue := elf[currentIndex]
		elf[currentIndex] = currentValue + i
	}

	// sort.Sort(sort.Reverse(sort.IntSlice(elf)))
	// max := 0
	// for _, num := range elf[0:3] {
	// 	max += num
	// }

	max := 0

	sort.Ints(elf)

	for _, i := range elf[len(elf)-3:] {
		max += i
	}

	fmt.Println(max)
}
