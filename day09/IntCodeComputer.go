package main

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type IntCodeComputer struct {
	offset      *big.Int
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

	computer := NewComputer(input)

	computer.ProcessProgram()

	// for _, value := range computer.output {
	// 	println(value.String())
	// }
	println(computer.output[len(computer.output)-1].String())

}

func NewComputer(rawProgram []string) *IntCodeComputer {
	computer := &IntCodeComputer{
		offset:      big.NewInt(0),
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
	for ic.read(index).Cmp(big.NewInt(99)) != 0 {
		opCode, firstMode, secondMode, thirdMode := parseInstruction(ic.read(index))
		var firstParam, secondParam, thirdParam *big.Int
		// 0: POSITION, 1: immediate, 2: relative
		if firstMode == 0 {
			firstParam = ic.read(new(big.Int).Add(index, big.NewInt(1)))
		} else if firstMode == 1 {
			firstParam = new(big.Int).Add(index, big.NewInt(1))
		} else if firstMode == 2 {
			firstParam = ic.read(new(big.Int).Add(index, big.NewInt(1)))
			firstParam = firstParam.Add(firstParam, ic.offset)
		} else {
			panic(firstMode)
		}

		if secondMode == 0 {
			secondParam = ic.read(new(big.Int).Add(index, big.NewInt(2)))
		} else if secondMode == 1 {
			secondParam = new(big.Int).Add(index, big.NewInt(2))
		} else if secondMode == 2 {
			secondParam = ic.read(new(big.Int).Add(index, big.NewInt(2)))
			secondParam = secondParam.Add(secondParam, ic.offset)
		} else {
			panic(secondMode)
		}

		if thirdMode == 0 {
			thirdParam = ic.read(new(big.Int).Add(index, big.NewInt(3)))
		} else if thirdMode == 1 {
			thirdParam = new(big.Int).Add(index, big.NewInt(3))
		} else if thirdMode == 2 {
			thirdParam = ic.read(new(big.Int).Add(index, big.NewInt(3)))
			thirdParam = thirdParam.Add(thirdParam, ic.offset)
		} else {
			panic(thirdMode)
		}
		switch opCode {
		case 1: // ADD
			result := new(big.Int).Add(ic.read(firstParam), ic.read(secondParam))
			ic.write(result, thirdParam)
			index = index.Add(index, big.NewInt(4))
		case 2: // MULT
			result := new(big.Int).Mul(ic.read(firstParam), ic.read(secondParam))
			ic.write(result, thirdParam)
			index = index.Add(index, big.NewInt(4))
		case 3: // INPUT
			ic.write(big.NewInt(2), firstParam) // TODO: CHANGE INPUT VALUE
			index = index.Add(index, big.NewInt(2))
		case 4: // OUTPUT
			ic.output = append(ic.output, ic.read(firstParam))
			index = index.Add(index, big.NewInt(2))
		case 5: // jump if true
			if ic.read(firstParam).Cmp(big.NewInt(0)) != 0 {
				index = ic.read(secondParam)
			} else {
				index = index.Add(index, big.NewInt(3))
			}
		case 6: // jump if false
			if ic.read(firstParam).Cmp(big.NewInt(0)) == 0 {
				index = ic.read(secondParam)
			} else {
				index = index.Add(index, big.NewInt(3))
			}
		case 7: // LESS THAN
			if ic.read(firstParam).Cmp(ic.read(secondParam)) < 0 {
				ic.write(big.NewInt(1), thirdParam)
			} else {
				ic.write(big.NewInt(0), thirdParam)
			}
			index = index.Add(index, big.NewInt(4))
		case 8: // EQUALS
			if ic.read(firstParam).Cmp(ic.read(secondParam)) == 0 {
				ic.write(big.NewInt(1), thirdParam)
			} else {
				ic.write(big.NewInt(0), thirdParam)
			}
			index = index.Add(index, big.NewInt(4))
		case 9:
			ic.offset = new(big.Int).Add(ic.offset, ic.read(firstParam))
			index = index.Add(index, big.NewInt(2))
		default:
			panic(opCode)
		}

	}
}

func (ic *IntCodeComputer) read(index *big.Int) *big.Int {
	answer := new(big.Int)
	if i := int(index.Int64()); index.IsInt64() && i < len(ic.program) {
		return answer.Set(ic.program[i])
	}
	if value, ok := ic.extraMemory[index.String()]; ok {
		return answer.Set(value)
	} else {
		ic.extraMemory[index.String()] = &big.Int{}
		return answer.Set(ic.extraMemory[index.String()])
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
