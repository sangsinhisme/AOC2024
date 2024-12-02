package Day2

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 2

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

	n := len(lines)
	input := make([][]int, n)
	for i, line := range lines {
		split := strings.Split(line, " ")
		m := len(split)
		if m == 0 {
			m = len(split)
		}
		input[i] = make([]int, m)
		for j, elem := range split {
			input[i][j], _ = strconv.Atoi(elem)
		}
	}
	ans := 0
	for i := range input {
		flagIncrease := true
		if input[i][0] > input[i][1] {
			flagIncrease = false
		}
		flag := true
		m := len(input[i])
		for j := 1; j < m; j++ {
			diff := input[i][j] - input[i][j-1]
			if flagIncrease && (diff < 1 || diff > 3) {
				flag = false
				break
			}
			if !flagIncrease && (diff > -1 || diff < -3) {
				flag = false
				break
			}
		}
		if flag {
			ans++
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

	n := len(lines)
	input := make([][]int, n)
	for i, line := range lines {
		split := strings.Split(line, " ")
		m := len(split)
		if m == 0 {
			m = len(split)
		}
		input[i] = make([]int, m)
		for j, elem := range split {
			input[i][j], _ = strconv.Atoi(elem)
		}
	}
	ans := 0
	for i := range input {
		flag := helper(input[i])
		if flag == true {
			ans++
		} else {
			for j := range len(input[i]) {
				newSlice := append([]int{}, input[i][:j]...)
				newSlice = append(newSlice, input[i][j+1:]...)
				flag = helper(newSlice)
				if flag {
					ans++
					break
				}
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}

func helper(input []int) bool {
	flagIncrease := true
	if input[0] > input[1] {
		flagIncrease = false
	}
	m := len(input)
	for j := 1; j < m; j++ {
		diff := input[j] - input[j-1]
		if flagIncrease && (diff < 1 || diff > 3) {
			return false
		}
		if !flagIncrease && (diff > -1 || diff < -3) {
			return false
		}
	}
	return true
}
