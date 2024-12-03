package Day3

import (
	"AOD2024/Utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const day = 3

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
		re := regexp.MustCompile(`mul\((?P<I1>\d+),(?P<I2>\d+)\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			ans += num1 * num2
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
	lastStatus := true
	for _, line := range lines {
		re := regexp.MustCompile(`mul\((?P<I1>\d+),(?P<I2>\d+)\)`)
		matches := re.FindAllStringSubmatchIndex(line, -1)

		mapDoDont := make(map[int]bool)
		re = regexp.MustCompile(`do\(\)`)
		do := re.FindAllStringSubmatchIndex(line, -1)

		re = regexp.MustCompile(`don't\(\)`)
		dont := re.FindAllStringSubmatchIndex(line, -1)

		for _, elem := range do {
			mapDoDont[elem[0]] = true
		}
		for _, elem := range dont {
			mapDoDont[elem[0]] = false
		}

		for _, match := range matches {
			lastStatus = fineNearstDoDont(mapDoDont, match[0], lastStatus)
			if lastStatus {
				num1, _ := strconv.Atoi(line[match[2]:match[3]])
				num2, _ := strconv.Atoi(line[match[4]:match[5]])
				ans += num1 * num2
			}
		}
	}

	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}

func fineNearstDoDont(doDont map[int]bool, idx int, lastStatus bool) bool {
	nearst := math.MaxInt
	status := true
	for elem := range doDont {
		if elem < idx {
			if idx-elem < nearst {
				nearst = idx - elem
				status = doDont[elem]
			}
		}
	}
	if nearst == math.MaxInt {
		return lastStatus
	}
	return status
}
