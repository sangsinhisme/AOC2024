package Day4

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 4

func XMAS(n, m int, graph [][]int32) int {
	var helper func(i, j, ii, jj int) bool
	helper = func(i, j, ii, jj int) bool {
		for _, elem := range "XMAS" {
			if i < 0 || i > n-1 || j < 0 || j > m-1 || graph[i][j] != elem {
				return false
			}
			i = i + ii
			j = j + jj
		}
		return true
	}
	direct := make([][2]int, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i != 0 || j != 0 {
				direct = append(direct, [2]int{i, j})
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range direct {
				if helper(i, j, d[0], d[1]) {
					ans++
				}
			}
		}
	}
	return ans
}

func MAS(n, m int, graph [][]int32) int {
	direct := [][3]int{
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}

	var helper func(i, j int) bool
	helper = func(i, j int) bool {
		idxMap := make(map[int]int32)
		for idx, elem := range direct {
			nextI, nextJ := i+elem[0], j+elem[1]
			if nextI > -1 && nextI < n && nextJ > -1 && nextJ < m {
				idxMap[idx] = graph[nextI][nextJ]
			}
		}
		if len(idxMap) != 4 {
			return false
		}
		mCountFirst := 0
		sCountFirst := 0
		mCountSecond := 0
		sCountSecond := 0

		for i, val := range idxMap {
			switch val {
			case 'M':
				if i < 2 {
					mCountFirst++
				} else {
					mCountSecond++
				}
			case 'S':
				if i < 2 {
					sCountFirst++
				} else {
					sCountSecond++
				}
			}
		}

		return mCountFirst == 1 && sCountFirst == 1 && mCountSecond == 1 && sCountSecond == 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 'A' && helper(i, j) {
				ans++
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
	ans := MAS(n, m, graph)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
