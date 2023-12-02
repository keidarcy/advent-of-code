package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

type OperateFunc func(oldValue int, divide int) (levelAfterBoring int)

type Monkey struct {
	id        int
	items     []int
	operate   OperateFunc
	nextTrue  int
	nextFalse int
	divide    int
	inspected int
}

func throwToNextMonkey(newLevel int, oldLevel int, index int, monkeys []Monkey) {
	monkey := monkeys[index]
	items := []int{}
	for _, item := range monkey.items {
		if item == oldLevel {
			continue
		}
		items = append(items, item)
	}
	monkeys[index].items = items

	if newLevel%monkey.divide == 0 {
		monkeys[monkey.nextTrue].items = append(monkeys[monkey.nextTrue].items, newLevel)
	} else {
		monkeys[monkey.nextFalse].items = append(monkeys[monkey.nextFalse].items, newLevel)
	}
	// fmt.Printf("newLevel: %d, oldLevel: %d, index: %d\n", newLevel, oldLevel, index)
	monkeys[index].inspected++
}

func getPart1(input string) int {

	lines := strings.Split(input, "\n")
	monkeys := []Monkey{}

	id := 0
	for i, l := range lines {
		if l == "" || rune(l[0]) == ' ' {
			continue
		}

		if rune(l[0]) == 'M' {
			id++
			monkeys = append(monkeys, parseMonkey(lines[i:i+6]))
		} else {
			continue
		}

	}

	for turn := 0; turn < 20; turn++ {
		for i, m := range monkeys {
			items := make([]int, len(m.items))
			copy(items, m.items)
			for _, oldLevel := range items {
				newLevel := m.operate(oldLevel, 3)
				throwToNextMonkey(newLevel, oldLevel, i, monkeys)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})

	return monkeys[0].inspected * monkeys[1].inspected
}

func getPart2(input string) int {

	lines := strings.Split(input, "\n")
	monkeys := []Monkey{}
	bigLimit := 1

	id := 0
	for i, l := range lines {
		if l == "" || rune(l[0]) == ' ' {
			continue
		}

		if rune(l[0]) == 'M' {
			id++
			monkeys = append(monkeys, parseMonkey(lines[i:i+6]))
			bigLimit *= monkeys[len(monkeys)-1].divide
		} else {
			continue
		}

	}

	fmt.Println(bigLimit)
	for turn := 0; turn < 10000; turn++ {
		for i, m := range monkeys {
			items := make([]int, len(m.items))
			copy(items, m.items)
			for _, oldLevel := range items {
				newLevel := m.operate(oldLevel, 1) % bigLimit
				throwToNextMonkey(newLevel, oldLevel, i, monkeys)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	// dump.P(monkeys)

	return monkeys[0].inspected * monkeys[1].inspected
}

func parseMonkey(lines []string) Monkey {
	// 1st line(Monkey 0:)
	id, err := strconv.Atoi(string(lines[0][7]))
	if err != nil {
		log.Fatal("BAD integer INPUT")
	}

	items := []int{}

	// 2nd line(  Starting items: 79, 98)
	for _, i := range strings.Split(strings.Replace(lines[1], "  Starting items: ", "", 1), ", ") {
		item, _ := strconv.Atoi(i)
		items = append(items, item)
	}

	// 3rd line(  Operation: new = old * 11)
	operationValue := strings.Replace(lines[2], "  Operation: new = old ", "", 1)

	// 4th line(  Test: divisible by 19)
	divide, _ := strconv.Atoi(strings.Replace(lines[3], "  Test: divisible by ", "", 1))

	// 5th line(    If true: throw to monkey 3)
	nextTrue, _ := strconv.Atoi(strings.Replace(lines[4], "    If true: throw to monkey ", "", 1))

	// 6th line(    If false: throw to monkey 6)
	nextFalse, _ := strconv.Atoi(strings.Replace(lines[5], "    If false: throw to monkey ", "", 1))

	monkey := Monkey{
		id:    id,
		items: items,
		operate: func(oldValue int, divide int) int {
			result := 0
			newValue, err := strconv.Atoi(operationValue[2:])
			if err != nil {
				newValue = oldValue
			}
			switch rune(operationValue[0]) {
			case '+':
				result = oldValue + newValue
			case '-':
				result = oldValue + newValue
			case '*':
				result = oldValue * newValue
			case '/':
				result = oldValue / newValue
			}
			return result / divide
		},
		divide:    divide,
		nextTrue:  nextTrue,
		nextFalse: nextFalse,
		inspected: 0,
	}

	return monkey
}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
