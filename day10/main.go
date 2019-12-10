package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var space [][]rune
	for scanner.Scan() {
		line := []rune(scanner.Text())
		space = append(space, line)
	}

	maxNumberOfAsteroids := 0
	for i, line := range space {
		for j, _ := range line {
			if string(space[i][j]) != "#" {
				continue
			}
			if numberOfVisibleAsteroids(i, j, space) > maxNumberOfAsteroids {
				maxNumberOfAsteroids = numberOfVisibleAsteroids(i, j, space)
			}
		}
	}
	println(maxNumberOfAsteroids)
	// Todo dont forget to convert coordinates at the end
}

func numberOfVisibleAsteroids(x, y int, space [][]rune) int {
	asteroidCounting := make(map[string]int)
	for i, line := range space {
		for j, _ := range line {
			if x == i && y == j || string(space[i][j]) != "#" {
				continue
			}
			dx := i - x
			dy := j - y
			atan := math.Atan2(float64(dy), float64(dx))
			rounded := fmt.Sprintf("%.8f", atan)
			asteroidCounting[rounded] = 1
		}
	}
	return len(asteroidCounting)
}
