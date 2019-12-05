package main

import (
	"bufio"
	"fmt"
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

	processProgram(originalProgram)

}

// OPcode, first, sec and third params
func parseInstruction(n int) (int, int, int, int) {
	instruction := strconv.Itoa(n)
	for len(instruction) < 5 {
		instruction = "0" + instruction
	}
	opcode, err := strconv.Atoi(instruction[3:])
	if err != nil {
		panic(err)
	}
	firstMode, err := strconv.Atoi(instruction[2:3])
	if err != nil {
		panic(err)
	}
	secondMode, err := strconv.Atoi(instruction[1:2])
	if err != nil {
		panic(err)
	}
	thirdMode, err := strconv.Atoi(instruction[0:1])
	if err != nil {
		panic(err)
	}
	return opcode, firstMode, secondMode, thirdMode
}

func processProgram(program []int) {
	for index := 0; program[index] != 99; {
		opCode, firstMode, secondMode, thirdMode := parseInstruction(program[index])
		var firstParam, secondParam, thirdParam int
		if firstMode == 0 {
			firstParam = program[index+1]
		} else {
			firstParam = index + 1
		}

		if opCode != 3 && opCode != 4 {
			if secondMode == 0 {
				secondParam = program[index+2]
			} else {
				secondParam = index + 2
			}
		}
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			if thirdMode == 0 {
				thirdParam = program[index+3]
			} else {
				thirdParam = index + 3
			}
		}

		if opCode == 1 {
			program[thirdParam] = program[firstParam] + program[secondParam]
			index += 4
		} else if opCode == 2 {
			program[thirdParam] = program[firstParam] * program[secondParam]
			index += 4
		} else if opCode == 3 {
			program[firstParam] = 5
			index += 2
		} else if opCode == 4 {
			println(fmt.Sprintf("Index of output %d", index))
			println(program[firstParam])
			index += 2
		} else if opCode == 5 {
			if program[firstParam] != 0 {
				index = program[secondParam]
			} else {
				index += 3
			}
		} else if opCode == 6 {
			if program[firstParam] == 0 {
				index = program[secondParam]
			} else {
				index += 3
			}
		} else if opCode == 7 {
			if program[firstParam] < program[secondParam] {
				program[thirdParam] = 1
			} else {
				program[thirdParam] = 0
			}
			index += 4
		} else if opCode == 8 {
			if program[firstParam] == program[secondParam] {
				program[thirdParam] = 1
			} else {
				program[thirdParam] = 0
			}
			index += 4
		} else {
			panic(opCode)
		}
	}
}
