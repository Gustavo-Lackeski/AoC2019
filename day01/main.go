package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		for ; value > 0; {
			value = value/3 - 2
			if value > 0 {
				total += value
			}
		}
	}
	println(total)

}

func partOne() {
	file, err := os.Open("input01.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		total += value/3 - 2
	}
	println(total)
}