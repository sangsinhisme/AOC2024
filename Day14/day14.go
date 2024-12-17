package Day14

import (
	"AOD2024/Utils"
	"fmt"
	"regexp"
	"strconv"
)

const day = 14
const title = "--- Day 14: Restroom Redoubt ---"

func teleport(n, m, pN, pM, vN, vM int) (int, int) {
	pN = (pN + vN + n) % (n)
	pM = (pM + vM + m) % (m)
	return pN, pM
}

func timeFindEasterEgg(n, m int, robots [][4]int) int {
	time := 1
	update, state := easterEgg(n, m, robots)
	for !state {
		update, state = easterEgg(n, m, update)
		time++
	}

	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	for _, robot := range update {
		graph[robot[0]][robot[1]]++
	}
	for i := range n {
		for j := range m {
			if graph[i][j] != 0 {
				fmt.Printf("%v", graph[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return time
}

func easterEgg(n, m int, robots [][4]int) ([][4]int, bool) {
	state := true
	graph := make(map[[2]int]int)
	newRobots := make([][4]int, len(robots))

	for i, robot := range robots {
		N, M := teleport(n, m, robot[0], robot[1], robot[2], robot[3])
		newRobots[i] = [4]int{N, M, robot[2], robot[3]}
		graph[[2]int{N, M}]++
		if graph[[2]int{N, M}] > 1 {
			state = false
		}
	}
	return newRobots, state
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
	n := 103
	m := 101
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	for _, line := range lines {
		re := regexp.MustCompile(`[pv]=(?P<i>-?\d+),(?P<j>-?\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		pM, _ := strconv.Atoi(matches[0][1])
		pN, _ := strconv.Atoi(matches[0][2])
		vM, _ := strconv.Atoi(matches[1][1])
		vN, _ := strconv.Atoi(matches[1][2])
		x, y := teleport(n, m, pN, pM, vN, vM)
		graph[x][y]++
	}
	midN := (n - 1) / 2
	midM := (m - 1) / 2
	quadrant1, quadrant2, quadrant3, quadrant4 := 0, 0, 0, 0
	for i := range n {
		for j := range m {
			if i < midN {
				if j < midM {
					quadrant1 += graph[i][j]
				}
				if j > midM {
					quadrant2 += graph[i][j]
				}
			}
			if i > midN {
				if j < midM {
					quadrant3 += graph[i][j]
				}
				if j > midM {
					quadrant4 += graph[i][j]
				}
			}
		}
	}
	ans = quadrant1 * quadrant2 * quadrant3 * quadrant4
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
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
	n := 103
	m := 101
	robots := make([][4]int, len(lines))
	for i, line := range lines {
		re := regexp.MustCompile(`[pv]=(?P<i>-?\d+),(?P<j>-?\d+)`)
		matches := re.FindAllStringSubmatch(line, -1)
		pM, _ := strconv.Atoi(matches[0][1])
		pN, _ := strconv.Atoi(matches[0][2])
		vM, _ := strconv.Atoi(matches[1][1])
		vN, _ := strconv.Atoi(matches[1][2])
		robots[i] = [4]int{pN, pM, vN, vM}
	}
	ans := timeFindEasterEgg(n, m, robots)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
