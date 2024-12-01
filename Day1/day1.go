package Day1

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 1

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
	left := make([]int, n)
	right := make([]int, n)

	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(split[0])
		right[i], _ = strconv.Atoi(split[1])
	}
	ans := 0
	for i := range left {
		ans += Utils.Abs(left[i], right[i])
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
	left := make([]int, n)
	right := make([]int, n)

	for i, line := range lines {
		split := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(split[0])
		right[i], _ = strconv.Atoi(split[1])
	}
	freqRight := make(map[int]int)
	for _, elem := range right {
		freqRight[elem]++
	}
	ans := 0
	for _, elem := range left {
		ans += elem * freqRight[elem]
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
