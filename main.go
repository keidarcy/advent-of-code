package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getInput() string {
	// https://adventofcode.com/2022/day/1/input
	testLine := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	return testLine
}

func main() {
	lines := strings.Split(getInput(), "\n")
	elf := make([]int, 0)
	elf = append(elf, 0)
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

	sort.Sort(sort.Reverse(sort.IntSlice(elf)))
	sum := 0
	for _, num := range elf[0:3] {
		sum += num
	}

	fmt.Println(sum)
}
