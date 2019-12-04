package main

const firstNumber = 387638
const lastNumber = 919123

func main() {
	count := 0
	for n := firstNumber; n <= lastNumber; n++ {
		if IsPasswordCandidate(n) {
			count++
		}
	}
	println(count)
}

func IsPasswordCandidate(n int) bool {
	// digits in REVERSE ORDER
	digits := make([]int, 0, 6)
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	digitCount := make(map[int]int)
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}

	for i := 0; i < len(digits); i++ {
		if _, ok := digitCount[digits[i]]; !ok {
			digitCount[digits[i]] = 1
		} else {
			digitCount[digits[i]] = digitCount[digits[i]] + 1
		}
	}

	for _, count := range digitCount {
		if count == 2 {
			return true
		}
	}
	return false
}
