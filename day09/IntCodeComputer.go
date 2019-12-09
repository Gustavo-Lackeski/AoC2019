package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type IntCodeComputer struct {
	program     []*big.Int
	extraMemory map[string]*big.Int
	output      []*big.Int
}

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

func (ic *IntCodeComputer) NewComputer(rawProgram []string) *IntCodeComputer {
	computer := &IntCodeComputer{
		program:     make([]*big.Int, 0, len(rawProgram)),
		extraMemory: make(map[string]*big.Int),
	}
	for _, value := range rawProgram {
		bigInt, ok := new(big.Int).SetString(value, 10)
		if !ok {
			panic(value)
		}
		computer.program = append(computer.program, bigInt)
	}

	return computer
}

func (ic *IntCodeComputer) ProcessProgram() {
	index := big.NewInt(0)
	for program[index] != 99 {
		opCode, firstMode, secondMode, thirdMode := parseInstruction(ic.read(index))
		var firstParam, secondParam, thirdParam *big.Int
		if firstMode == 0 {
			firstParam = ic.read(new(big.Int).Add(index, big.NewInt(1)))
		} else {
			firstParam = new(big.Int).Add(index, big.NewInt(1))
		}

		if secondMode == 0 {
			secondParam = ic.read(new(big.Int).Add(index, big.NewInt(2)))
		} else {
			secondParam = new(big.Int).Add(index, big.NewInt(2))
		}

		if thirdMode == 0 {
			thirdParam = ic.read(new(big.Int).Add(index, big.NewInt(3)))
		} else {
			thirdParam = new(big.Int).Add(index, big.NewInt(3))
		}

		switch opCode {
		case 1:
			program[thirdParam] = program[firstParam] + program[secondParam]
			index += 4
		case 2:
			program[thirdParam] = program[firstParam] * program[secondParam]
			index += 4
		case 3:
			program[firstParam] = 5
			index += 2
		case 4:
			//TODO CHANGE
			println(fmt.Sprintf("Index of output %d", index))
			println(program[firstParam])
			index += 2
		case 5:
			if program[firstParam] != 0 {
				index = program[secondParam]
			} else {
				index += 3
			}
		case 6:
			if program[firstParam] == 0 {
				index = program[secondParam]
			} else {
				index += 3
			}
		case 7:
			if program[firstParam] < program[secondParam] {
				program[thirdParam] = 1
			} else {
				program[thirdParam] = 0
			}
			index += 4
		case 8:
			if program[firstParam] == program[secondParam] {
				program[thirdParam] = 1
			} else {
				program[thirdParam] = 0
			}
			index += 4
		default:
			panic(opCode)
		}

	}
}

func (ic *IntCodeComputer) read(index *big.Int) *big.Int {
	if i := int(index.Int64()); index.IsInt64() && i < len(ic.program) {
		return ic.program[i]
	}
	if value, ok := ic.extraMemory[index.String()]; ok {
		return value
	} else {
		ic.extraMemory[index.String()] = &big.Int{}
		return ic.extraMemory[index.String()]
	}
}

func (ic *IntCodeComputer) write(value, index *big.Int) {
	if i := int(index.Int64()); index.IsInt64() && i < len(ic.program) {
		ic.program[i] = value
	} else {
		ic.extraMemory[index.String()] = value
	}
}

// OPcode, first, sec and third params
func parseInstruction(n *big.Int) (int, int, int, int) {
	instruction := n.String()
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
