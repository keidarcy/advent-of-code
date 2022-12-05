package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	contain := getTop(input)
	// overlap := getOther(input)
	fmt.Println("top", contain)
	// fmt.Println("other", overlap)
}

func getTop(input string) string {
	lines := strings.Split(input, "\n")
	isStack := true
	for _, l := range lines {
		// if parseLine(l) {
		// }

		if l == "" {
			isStack = false
		}

		if isStack {
			// ele := strings.Split(l, " ")
			fmt.Println(l)
			fmt.Println(len(l))
		} else {

		}

	}
	return ""
}

// func parseLine(line string) bool {
// 	return false
// }

// func getOther(input string) string {
// 	return ""
// }

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
