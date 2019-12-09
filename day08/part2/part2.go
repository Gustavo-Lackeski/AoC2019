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

	//LAYER I, position (x,y): rawInput(I*width*height + width*y + x)

	//for each point..
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			//from first to last layer
			for i := 0; i < numberOfLayers; i++ {
				c := string(rawInput[i*width*height+width*y+x])
				if c != "2" {
					if c == "0" {
						print(" ")
					} else {
						print("A")
					}
					break
				}
			}
		}
		println()
	}

}
