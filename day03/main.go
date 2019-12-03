package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	firstInput := scanner.Text()
	scanner.Scan()
	secondInput := scanner.Text()

	firstWire := parseInput(firstInput)
	secondWire := parseInput(secondInput)
	min := getMinSteps(firstWire, secondWire)
	println(min)
}

func parseInput(input string) map[Coordinate]int {
	movements := strings.Split(input, ",")
	wire := make(map[Coordinate]int)
	curX := 0
	curY := 0
	totalLength := 0

	for _, movement := range movements {
		dirX := 0
		dirY := 0
		direction := string(movement[0])
		if direction == "R" {
			dirX = 1
		} else if direction == "L" {
			dirX = -1
		} else if direction == "U" {
			dirY = 1
		} else if direction == "D" {
			dirY = -1
		} else {
			panic(fmt.Sprintf("Wrong direction %s", direction))
		}

		length, err := strconv.Atoi(movement[1:])
		if err != nil {
			panic(fmt.Sprintf("Couldnt parse %s", movement[1:]))
		}

		for i := 1; i <= length; i++ {
			X := curX + i*dirX
			Y := curY + i*dirY
			Coordinate := Coordinate{
				X: X,
				Y: Y,
			}
			if _ , ok := wire[Coordinate]; !ok {
				wire[Coordinate] = totalLength + i
			}
		}

		curX += length * dirX
		curY += length * dirY
		totalLength += length
	}

	return wire
}

func getMinSteps(firstWire, secondWire map[Coordinate]int) int {
	steps := 99999999
	for point, firstSteps := range firstWire {
		if secondSteps, ok := secondWire[point]; ok {
			numberOfSteps := Abs(firstSteps) + Abs(secondSteps)
			if numberOfSteps > 0 {
				steps = Min(steps, numberOfSteps)
			}
		}
	}
	return steps
}

func getMinIntersection(firstWire, secondWire map[Coordinate]int) int {
	minDist := 99999999
	for point := range firstWire {
		if _, ok := secondWire[point]; ok {
			dist := Abs(point.X) + Abs(point.Y)
			if dist > 0 {
				minDist = Min(minDist, dist)
			}
		}
	}
	return minDist
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
