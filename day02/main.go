package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	rawInput := scanner.Text()
	input := strings.Split(rawInput, ",")
	originalProgram := make([]int, 0, len(input))
	for _, value := range input {
		parsedValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		originalProgram = append(originalProgram, parsedValue)
	}

	program := make([]int, len(originalProgram))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(program, originalProgram)
			program[1] = noun
			program[2] = verb
			if processProgram(program) == 19690720 {
				println(100*noun + verb)
			}
		}
	}
}

func processProgram(program []int) int {
	for index := 0; program[index] != 99; index += 4 {
		opCode := program[index]
		firstInput := program[index+1]
		secondInput := program[index+2]
		resultPosition := program[index+3]
		if opCode == 1 {
			program[resultPosition] = program[firstInput] + program[secondInput]
		} else if opCode == 2 {
			program[resultPosition] = program[firstInput] * program[secondInput]
		} else {
			panic(opCode)
		}
	}
	return program[0]
}
