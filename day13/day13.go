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
	// part1 := getPart1(input)
	part2 := getPart2(input)
	// fmt.Println("part1 result: ", part1)
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

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
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

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func main() {
// 	input, _ := os.ReadFile("input.txt")

// 	pkts, part1 := []any{}, 0
// 	for i, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
// 		s := strings.Split(s, "\n")
// 		var a, b any
// 		json.Unmarshal([]byte(s[0]), &a)
// 		json.Unmarshal([]byte(s[1]), &b)
// 		pkts = append(pkts, a, b)

// 		if cmp(a, b) <= 0 {
// 			part1 += i + 1
// 		}
// 	}
// 	fmt.Println(part1)

// 	pkts = append(pkts, []any{[]any{2.}}, []any{[]any{6.}})
// 	sort.Slice(pkts, func(i, j int) bool { return cmp(pkts[i], pkts[j]) < 0 })

// 	part2 := 1
// 	for i, p := range pkts {
// 		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
// 			part2 *= i + 1
// 		}
// 	}
// 	fmt.Println(part2)
// }

// func cmp(a, b any) int {
// 	as, aok := a.([]any)
// 	bs, bok := b.([]any)

// 	switch {
// 	case !aok && !bok:
// 		return int(a.(float64) - b.(float64))
// 	case !aok:
// 		as = []any{a}
// 	case !bok:
// 		bs = []any{b}
// 	}

// 	for i := 0; i < len(as) && i < len(bs); i++ {
// 		if c := cmp(as[i], bs[i]); c != 0 {
// 			return c
// 		}
// 	}
// 	return len(as) - len(bs)
// }
