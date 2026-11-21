package Day20

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 20
const title = "--- Day 20: Race Condition ---"

var moves = [][]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func solvePath(graph [][]uint8, start, end [2]int, n, m int) ([][2]int, map[[2]int]int) {
	stack := Utils.NewStack()
	stack.Push([3]int{start[0], start[1], 0})
	visited := map[[2]int]int{}
	for stack.Length() > 0 {
		curr := stack.Pop().([3]int)
		visited[[2]int{curr[0], curr[1]}] = curr[2]
		for _, move := range moves {
			nextI, nextJ := curr[0]+move[0], curr[1]+move[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && visited[[2]int{nextI, nextJ}] == 0 {
				if graph[nextI][nextJ] == '.' || graph[nextI][nextJ] == 'E' {
					stack.Push([3]int{nextI, nextJ, curr[2] + 1})
				}
			}
		}
	}
	graph[end[0]][end[1]] = '.'
	pathLength := len(visited)
	path := make([][2]int, pathLength)
	for visitPoint, counter := range visited {
		path[counter] = visitPoint
	}

	return path, visited
}

func resolvePart1(graph [][]uint8, start, end [2]int, n, m int) int {
	path, visited := solvePath(graph, start, end, n, m)
	totalReduce := map[int]int{}
	for i, point := range path {
		for _, move := range moves {
			nextI, nextJ := point[0]+move[0], point[1]+move[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && graph[nextI][nextJ] == '#' {
				nextI1, nextJ1 := nextI+move[0], nextJ+move[1]
				if nextI1 >= 0 && nextI1 < n && nextJ1 >= 0 && nextJ1 < m && graph[nextI1][nextJ1] == '.' {
					reduce := visited[[2]int{nextI1, nextJ1}] - i - 2
					if reduce > 0 {
						totalReduce[reduce] += 1
					}
				}
			}
		}
	}
	ans := 0
	for numberReduce, count := range totalReduce {
		if numberReduce >= 100 {
			ans += count
		}
	}
	return ans
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	n := len(lines)
	m := len(lines[0])
	graph := make([][]uint8, n)
	start, end := [2]int{0, 0}, [2]int{0, 0}
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	for i, line := range lines {
		for j, char := range line {
			graph[i][j] = uint8(char)
			if char == 'S' {
				start = [2]int{i, j}
			}
			if char == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	ans := resolvePart1(graph, start, end, n, m)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
}

func helperLoop(graph [][]uint8, curr [2]int, n, m int, records map[[2]int]int, totalReduce map[int]int) {
	queue := Utils.NewQueue()
	queue.Push([3]int{curr[0], curr[1], 0})
	visited := map[[2]int]bool{
		curr: true,
	}
	for queue.Length() > 0 {
		point := queue.Pop().([3]int)
		if point[2] == 20 {
			continue
		}
		cost := point[2] + 1
		for _, move := range moves {
			nextI, nextJ := point[0]+move[0], point[1]+move[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m && !visited[[2]int{nextI, nextJ}] {
				if graph[nextI][nextJ] != '#' {
					nextRecord := records[[2]int{nextI, nextJ}] - records[curr] - cost
					if nextRecord > 0 {
						totalReduce[nextRecord] += 1
					}
				}
				queue.Push([3]int{nextI, nextJ, cost})
				visited[[2]int{nextI, nextJ}] = true
			}
		}
	}
}

func resolvePart2(graph [][]uint8, start, end [2]int, n, m int) int {
	path, records := solvePath(graph, start, end, n, m)
	totalReduce := map[int]int{}
	for _, point := range path {
		helperLoop(graph, point, n, m, records, totalReduce)
	}
	ans := 0
	for numberReduce, count := range totalReduce {
		if numberReduce >= 100 {
			fmt.Printf("There are %d cheats that save %d picoseconds.\n", count, numberReduce)
			ans += count
		}
	}
	return ans
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	n := len(lines)
	m := len(lines[0])
	graph := make([][]uint8, n)
	start, end := [2]int{0, 0}, [2]int{0, 0}
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	for i, line := range lines {
		for j, char := range line {
			graph[i][j] = uint8(char)
			if char == 'S' {
				start = [2]int{i, j}
			}
			if char == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	ans := resolvePart2(graph, start, end, n, m)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
