package Day4

import (
	"AOD2024/Utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const day = 4

func XMAS(n, m int, graph [][]int32) int {
	var helper func(i, j int, visited map[[2]int]int, xmas string) int
	helper = func(i, j int, visited map[[2]int]int, xmas string) int {
		ans := 0
		if strings.HasPrefix("XMAS", xmas) {
			if xmas == "XMAS" {
				return 1
			}
			visited[[2]int{i, j}] = 1
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					nextCurrX, nextCurrY := i+x, j+y
					if nextCurrX > -1 && nextCurrX < n && nextCurrY > -1 && nextCurrY < m && visited[[2]int{nextCurrX, nextCurrY}] != 1 {
						nextString := xmas + string(graph[nextCurrX][nextCurrY])
						if strings.HasPrefix("XMAS", nextString) {
							fmt.Println(i, j, nextCurrX, nextCurrY, nextString)
							visited[[2]int{nextCurrX, nextCurrY}] = 1
							ans += helper(nextCurrX, nextCurrY, visited, nextString)
							visited[[2]int{nextCurrX, nextCurrY}] = 0
						}

					}
				}
			}
		}
		return ans
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 'X' {
				ans += helper(i, j, map[[2]int]int{}, "X")
			}
		}
	}
	return ans
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
	n := len(lines)
	m := len(lines[0])
	graph := make([][]int32, n)
	for i := range n {
		graph[i] = make([]int32, m)
	}
	for i, line := range lines {
		for j, char := range line {
			graph[i][j] = char
		}
	}
	ans := XMAS(n, m, graph)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	//if submit {
	//	Utils.Submit(day, 1, ans)
	//}
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
