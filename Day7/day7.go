package Day7

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 7

func evaluateOperators(value int, operators []int) bool {
	n := len(operators)
	var helper func(idx, total int) bool
	helper = func(idx, total int) bool {
		if idx == n {
			return total == value
		}
		if total > value {
			return false
		}
		return helper(idx+1, total+operators[idx]) || helper(idx+1, total*operators[idx])
	}
	return helper(1, operators[0])
}

func generateOperatorCombinations(n int) [][]uint8 {
	if n <= 0 {
		return [][]uint8{}
	}

	operators := []uint8{'+', '*', '|'}
	var result [][]uint8

	var helper func(current []uint8)
	helper = func(current []uint8) {
		if len(current) == n {
			combination := make([]uint8, n)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for _, op := range operators {
			helper(append(current, op))
		}
	}

	helper([]uint8{})
	return result
}

func evaluateOperatorsPart2(value int, operators []int) bool {
	numDigits := func(num int) int {
		count := 0
		for num > 0 {
			num /= 10
			count++
		}
		return count
	}

	var helper func(opers []uint8) bool
	helper = func(opers []uint8) bool {
		init := operators[0]
		for idx, oper := range opers {
			if init > value {
				return false
			}
			if oper == '+' {
				init = init + operators[idx+1]
			}
			if oper == '*' {
				init = init * operators[idx+1]
			}
			if oper == '|' {
				concatDigits := numDigits(operators[idx+1])
				init = init*pow10(concatDigits) + operators[idx+1]
			}
		}
		return init == value
	}
	totalOperators := generateOperatorCombinations(len(operators) - 1)
	for _, opers := range totalOperators {
		if helper(opers) {
			return true
		}
	}
	return false
}

func pow10(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}

func Part1(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		value, _ := strconv.Atoi(split[0])
		operateStr := strings.Split(split[1], " ")
		operates := make([]int, len(operateStr))
		for idx, operate := range operateStr {
			operates[idx], _ = strconv.Atoi(operate)
		}
		if evaluateOperators(value, operates) {
			ans += value
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := 0
	step := 0
	for _, line := range lines {
		split := strings.Split(line, ": ")
		value, _ := strconv.Atoi(split[0])
		operateStr := strings.Split(split[1], " ")
		operates := make([]int, len(operateStr))
		for idx, operate := range operateStr {
			operates[idx], _ = strconv.Atoi(operate)
		}
		fmt.Printf("Process %v with %v", step, value)
		fmt.Println()
		step++
		if evaluateOperatorsPart2(value, operates) {
			ans += value
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
