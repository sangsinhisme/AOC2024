package Day10

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 10

func trailheadPath(n, m int, graph [][]int, init [3]int) int {
	queue := Utils.NewQueue()
	queue.Push(init)
	dicts := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	path := 0
	visited := make(map[[2]int]int)
	for queue.Length() > 0 {
		position := queue.Pop().([3]int)
		visited[[2]int{position[0], position[1]}] = 1
		if position[2] == 9 {
			path++
			continue
		}
		for _, dict := range dicts {
			nextN, nextM := position[0]+dict[0], position[1]+dict[1]
			if nextN > -1 && nextN < n && nextM > -1 && nextM < m {
				if graph[nextN][nextM] == position[2]+1 && visited[[2]int{nextN, nextM}] == 0 {
					queue.Push([3]int{nextN, nextM, position[2] + 1})
					visited[[2]int{nextN, nextM}] = 1
				}
			}
		}
	}
	return path
}

func trailheadPath2(n, m int, graph [][]int, init [3]int) int {
	queue := Utils.NewQueue()
	queue.Push(init)
	dicts := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	path := 0
	for queue.Length() > 0 {
		position := queue.Pop().([3]int)
		if position[2] == 9 {
			path++
			continue
		}
		for _, dict := range dicts {
			nextN, nextM := position[0]+dict[0], position[1]+dict[1]
			if nextN > -1 && nextN < n && nextM > -1 && nextM < m && graph[nextN][nextM] == position[2]+1 {
				queue.Push([3]int{nextN, nextM, position[2] + 1})
			}
		}
	}
	return path
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
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	var startPath [][2]int
	for i := range n {
		for j := range m {
			graph[i][j] = int(lines[i][j]) - '0'
			if graph[i][j] == 0 {
				startPath = append(startPath, [2]int{i, j})
			}
		}
	}
	ans := 0
	for _, init := range startPath {
		ans += trailheadPath(n, m, graph, [3]int{init[0], init[1], 0})
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
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	var startPath [][2]int
	for i := range n {
		for j := range m {
			graph[i][j] = int(lines[i][j]) - '0'
			if graph[i][j] == 0 {
				startPath = append(startPath, [2]int{i, j})
			}
		}
	}
	ans := 0
	for _, init := range startPath {
		ans += trailheadPath2(n, m, graph, [3]int{init[0], init[1], 0})
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
