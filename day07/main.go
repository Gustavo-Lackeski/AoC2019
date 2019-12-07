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

	maxOutput := 0
	allPermutations := permutations([]int{5, 6, 7, 8, 9})
	for _, phaseSettings := range allPermutations {
		//E-A, A-B, B-C, C-D, D-E
		var chans = []chan int{
			make(chan int, 1),
			make(chan int, 1),
			make(chan int, 1),
			make(chan int, 1),
			make(chan int, 1),
		}
		chans[0] <- 0
		finalOutputChan := make(chan int)
		for index, phase := range phaseSettings {
			go func(index int, phase int) {
				inputChan := chans[index]
				outputChan := chans[(index+1)%5]
				program := make([]int, len(originalProgram))
				copy(program, originalProgram)
				processProgram(program, phase, inputChan, outputChan, finalOutputChan, index)

			}(index, phase)
		}
		maxOutput = Max(maxOutput, <-finalOutputChan)
	}

	println(fmt.Sprintf("Final output: %d", maxOutput))

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

func processProgram(program []int, phase int, input, output, processFinished chan int, ampNumber int) {
	var finalOutput int
	//println("started program")
	for index := 0; program[index] != 99; {
		//println("new loop")
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
			if index == 0 {
				program[firstParam] = phase
			} else {
				program[firstParam] = <-input
			}
			index += 2
		} else if opCode == 4 {
			println(fmt.Sprintf("%d Waiting to send output", ampNumber))
			//println(program[firstParam])
			finalOutput = program[firstParam]
			output <- program[firstParam]
			println(fmt.Sprintf("%d output sent", ampNumber))
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
	if ampNumber == 4 {
		processFinished <- finalOutput
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
