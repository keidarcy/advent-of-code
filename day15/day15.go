package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := getInput()
	part1 := getPart1(input, 2000000)
	part2 := getPart2(input, 4000000)
	fmt.Println("part1 result: ", part1)
	fmt.Println("part2 result: ", part2)
}

func getPart1(input string, limit int) int {
	lines := strings.Split(input, "\n")
	targetLine := make(map[int]bool)

	for _, l := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)

		distanceFromLine := sensorY - limit
		if distanceFromLine < 0 {
			distanceFromLine *= -1
		}

		for i := 0; i <= distanceFromBeacon-distanceFromLine; i++ {
			targetLine[sensorX+i] = true
			targetLine[sensorX-i] = true
		}

		if beaconY == limit {
			delete(targetLine, beaconX)
		}
	}

	return len(targetLine)
}

type point struct {
	x, y int
}

func getPart2(input string, limit int) int {
	lines := strings.Split(input, "\n")
	nearestBeacon := make(map[point]int)
	toTry := make(map[point]bool)

	for _, l := range lines {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)

		distanceFromBeacon := manhattan(sensorX, sensorY, beaconX, beaconY)

		nearestBeacon[point{sensorX, sensorY}] = distanceFromBeacon
		distanceFromBeacon++
		for i := 0; i < distanceFromBeacon; i++ {
			if sensorX+i > 0 && sensorX+i < limit {
				if sensorY-distanceFromBeacon+1+i > 0 && sensorY-distanceFromBeacon+1+i < limit {
					toTry[point{sensorX + i, sensorY - distanceFromBeacon + 1 + i}] = true
				}
				if sensorY+distanceFromBeacon-1-i > 0 && sensorY+distanceFromBeacon-1-i < limit {
					toTry[point{sensorX + i, sensorY + distanceFromBeacon - i}] = true
				}
			}

			if sensorX-i > 0 && sensorX-i < limit {
				if sensorY-distanceFromBeacon+1+i > 0 && sensorY-distanceFromBeacon+1+i < limit {
					toTry[point{sensorX - i, sensorY - distanceFromBeacon + 1 + i}] = true
				}
				if sensorY+distanceFromBeacon-1-i > 0 && sensorY+distanceFromBeacon-1-i < limit {
					toTry[point{sensorX - i, sensorY + distanceFromBeacon - 1 - i}] = true
				}
			}
		}
	}

	var result int
	for beacon := range toTry {
		newBeacon := true
		for sensor, nearestBeaconDistance := range nearestBeacon {
			if manhattan(sensor.x, sensor.y, beacon.x, beacon.y) <= nearestBeaconDistance {
				newBeacon = false
				break
			}

		}
		if newBeacon {
			result = beacon.x*4000000 + beacon.y
			break
		}
	}

	return result
}

func manhattan(sensorX, sensorY, beaconX, beaconY int) int {
	diffX := sensorX - beaconX
	diffY := sensorY - beaconY

	if diffX < 0 {
		diffX *= -1
	}
	if diffY < 0 {
		diffY *= -1
	}
	return diffX + diffY

}

func getInput() string {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal("BAD INPUT")
	}

	return string(input)
}
