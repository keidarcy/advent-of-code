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

    // sum := 0
	// for _, s := range strings.Fields(string(input)) {
		// n := 0
		// for _, r := range s {
			// n = 5*n + map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}[r]
		// }
		// sum += n
	// }
    // fmt.Printf("sum: %v\n", sum)

    // snafu := ""
	// for sum > 0 {
		// snafu = string("=-012"[(sum+2)%5]) + snafu
    //     fmt.Printf("snafu: %v\n", snafu)
		// sum = (sum + 2) / 5
	// }
	// fmt.Println(snafu)
	part1 := getPart1(input)
	fmt.Println("part1 result: ", part1)
}

func getPart1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	num := 0
	for _, line := range lines {
		num += snafuToDecimal(line)
	}
	return decimalToSnafu(num)
}

func snafuToDecimal(line string) int {
	snafuDecimalMap := map[rune]int{
		'2': 2,
		'1': 1,
		'0': 0,
		'-': -1,
		'=': -2,
	}
	result := 0.0
	for i := len(line) - 1; i >= 0; i-- {
		y := len(line) - 1 - i
		currS := snafuDecimalMap[(rune(line[i]))]
		currN := math.Pow(5.0, float64(y)) * float64(currS)
		result += currN
	}
	return int(result)
}

func decimalToSnafu(snafu int) string {
	decimalSnafuNormalMap := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
	}

	tempSnafu := []string{}
	for snafu >= 1 {
		reminder := snafu % 5
		tempSnafu = append([]string{decimalSnafuNormalMap[reminder]}, tempSnafu...)
		snafu /= 5
	}


    i := len(tempSnafu) - 1
    for i >= 0 {
        temp, val := handleNonInt(tempSnafu, i)
        tempSnafu = temp
        i = i - 1 - val
    }

	snafuRes := []string{}

    for i := len(tempSnafu) - 1; i >= 0; i-- {
		val := string(tempSnafu[i])
		switch val {
		case "0":
			snafuRes = append([]string{"0"}, snafuRes...)
		case "1":
			snafuRes = append([]string{"1"}, snafuRes...)
		case "2":
			snafuRes = append([]string{"2"}, snafuRes...)
		case "3":
			snafuRes = append([]string{"="}, snafuRes...)
		case "4":
			snafuRes = append([]string{"-"}, snafuRes...)
		}
    }

	return strings.Join(snafuRes, "")
}

func handleNonInt(tempSnafu []string, i int) ([]string, int) {
    val := 0
	if tempSnafu[i] == "3" || tempSnafu[i] == "4" {
        if i == 0 {
            tempSnafu = append([]string{"1"}, tempSnafu...)
            return tempSnafu, val
        }
		switch tempSnafu[i-1] {
		case "0":
			tempSnafu[i-1] = "1"
			return tempSnafu, val
		case "1":
			tempSnafu[i-1] = "2"
			return tempSnafu, val
		case "2":
			tempSnafu[i-1] = "3"
		case "3":
			tempSnafu[i-1] = "4"
		case "4":
            val++
			tempSnafu, _ = handleNonInt(tempSnafu, i-1)
			tempSnafu[i-1] = "0"
		}
	}

	return tempSnafu, val
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
