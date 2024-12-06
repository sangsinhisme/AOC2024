package Day6

import (
	"AOD2024/Utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = 6

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
	graph := make([][]uint8, n)
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	currI, currJ := 0, 0
	direct := map[int][2]int{
		0: {-1, 0},
		1: {0, 1},
		2: {1, 0},
		3: {0, -1},
	}
	currDirect := 0
	for i, line := range lines {
		for j := range line {
			if line[j] == '^' {
				currI = i
				currJ = j
			}
			graph[i][j] = line[j]
		}
	}
	ans := 0
	visited := make(map[[2]int]int)
	for currI > -1 && currI < n && currJ > -1 && currJ < m {
		if visited[[2]int{currI, currJ}] != 1 {
			ans++
			visited[[2]int{currI, currJ}] = 1
		}
		nextI, nextJ := currI+direct[currDirect][0], currJ+direct[currDirect][1]
		if nextI > -1 && nextI < n && nextJ > -1 && nextJ < m && graph[nextI][nextJ] == '#' {
			currDirect = (currDirect + 1) % 4
		}
		currI, currJ = currI+direct[currDirect][0], currJ+direct[currDirect][1]
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
	pageCondition := make(map[int]map[int]int)
	printPage := make([][]int, 0)
	split := false
	for _, line := range lines {
		if line == "" {
			split = true
			continue
		}
		if !split {
			pages := strings.Split(line, "|")
			left, _ := strconv.Atoi(pages[0])
			right, _ := strconv.Atoi(pages[1])
			if pageCondition[left] == nil {
				pageCondition[left] = make(map[int]int)
			}
			pageCondition[left][right] = 1
		} else {
			page := strings.Split(line, ",")
			tempPage := make([]int, len(page))
			for i, elem := range page {
				tempPage[i], _ = strconv.Atoi(elem)
			}
			printPage = append(printPage, tempPage)
		}
	}
	ans := 0
	for _, line := range printPage {
		correctOrder := true
		for i := 1; i < len(line) && correctOrder; i++ {
			for j := 0; j < i && correctOrder; j++ {
				if pageCondition[line[i]][line[j]] == 1 {
					correctOrder = false
					break
				}
			}
		}
		if !correctOrder {
			sort.Slice(line, func(i, j int) bool {
				return pageCondition[line[i]][line[j]] == 1
			})
			middle := len(line) / 2
			ans += line[middle]
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
