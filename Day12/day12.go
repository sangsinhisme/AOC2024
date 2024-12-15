package Day12

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 12
const title = "--- Day 12: Garden Groups ---"

func gardenTravel(n, m int, graph [][]int) int {
	visited := make(map[[2]int]bool)
	directs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var helper func(i, j int) int
	helper = func(i, j int) int {
		queue := Utils.NewQueue()
		queue.Push([2]int{i, j})
		visited[[2]int{i, j}] = true
		plant := graph[i][j]
		region := 0
		perimeter := 0
		for queue.Length() > 0 {
			curr := queue.Pop().([2]int)
			neighbor := 0
			for _, direct := range directs {
				nextI := curr[0] + direct[0]
				nextJ := curr[1] + direct[1]
				if nextI > -1 && nextI < n && nextJ > -1 && nextJ < m && graph[nextI][nextJ] == plant {
					if visited[[2]int{nextI, nextJ}] != true {
						visited[[2]int{nextI, nextJ}] = true
						queue.Push([2]int{nextI, nextJ})
					}
					neighbor++
				}
			}
			region += 4 - neighbor
			perimeter++
		}
		return region * perimeter
	}
	ans := 0
	for i := range n {
		for j := range m {
			if visited[[2]int{i, j}] != true {
				ans += helper(i, j)
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
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	for i := range n {
		for j := range m {
			graph[i][j] = int(lines[i][j])
		}
	}

	ans := gardenTravel(n, m, graph)
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
	ans := len(lines)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
