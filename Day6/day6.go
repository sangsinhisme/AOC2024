package Day6

import (
	"AOD2024/Utils"
	"fmt"
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
	n := len(lines)
	m := len(lines[0])
	graph := make([][]uint8, n)
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	startI, startJ := 0, 0
	for i, line := range lines {
		for j := range line {
			if line[j] == '^' {
				startI = i
				startJ = j
			}
			graph[i][j] = line[j]
		}
	}

	direct := map[int][2]int{
		0: {-1, 0},
		1: {0, 1},
		2: {1, 0},
		3: {0, -1},
	}
	visited := make(map[[2]int]int)
	currDirect := 0
	currI, currJ := startI, startJ
	for currI > -1 && currI < n && currJ > -1 && currJ < m {
		if visited[[2]int{currI, currJ}] != 1 {
			visited[[2]int{currI, currJ}] = 1
		}
		nextI, nextJ := currI+direct[currDirect][0], currJ+direct[currDirect][1]
		if nextI > -1 && nextI < n && nextJ > -1 && nextJ < m && graph[nextI][nextJ] == '#' {
			currDirect = (currDirect + 1) % 4
		}
		currI, currJ = currI+direct[currDirect][0], currJ+direct[currDirect][1]
	}
	ans := 0
	for i := range n {
		for j := range m {
			if graph[i][j] == '.' {
				graph[i][j] = '#'
				if deathLoop(startI, startJ, n, m, graph) {
					ans++
				}
				graph[i][j] = '.'
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	//if submit {
	//	Utils.Submit(day, 2, ans)
	//}
}

func deathLoop(currI, currJ, n, m int, graph [][]uint8) bool {
	directions := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	visited := make(map[[2]int]int)
	currDirection := 0

	for currI >= 0 && currI < n && currJ >= 0 && currJ < m {
		pos := [2]int{currI, currJ}
		if visited[pos] > 4 {
			return true
		}
		visited[pos]++

		nextI, nextJ := currI+directions[currDirection][0], currJ+directions[currDirection][1]

		if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && graph[nextI][nextJ] == '#' {
			currDirection = (currDirection + 1) % 4
		} else {
			currI, currJ = nextI, nextJ
		}
	}
	return false
}
