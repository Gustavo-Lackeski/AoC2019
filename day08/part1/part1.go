package main

import (
	"bufio"
	"os"
	"strings"
)

const width = 25
const height = 6

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	rawInput := scanner.Text()
	rawInput = strings.TrimSpace(rawInput)
	numberOfLayers := len(rawInput) / (width * height)
	layerCounting := make([]map[rune]int, numberOfLayers)

	minZeroes := 0
	for i := 0; i < numberOfLayers; i++ {
		layerCounting[i] = map[rune]int{}
		//0:6, 6:12 ..
		for _, c := range rawInput[i*width*height : (i+1)*width*height] {
			layerCounting[i][c]++
		}
		if layerCounting[i]['0'] < layerCounting[minZeroes]['0']{
			minZeroes = i
		}
	}

	println(layerCounting[minZeroes]['1']*layerCounting[minZeroes]['2'])

}
