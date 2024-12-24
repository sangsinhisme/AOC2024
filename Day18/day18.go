package Day18

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 18
const title = "--- Day 18: RAM Run ---"
const n = 71
const m = 71
const splitAt = 1024

func dijkstra(graph [][]int) int {
	directs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := Utils.Queue{}
	queue.Push([3]int{0, 0, 1})
	visited := make(map[[2]int]int)
	for queue.Length() > 0 {
		curr := queue.Pop().([3]int)
		if curr[0] == n-1 && curr[1] == m-1 {
			return curr[2] - 1
		}
		for _, direct := range directs {
			nextCurrN := curr[0] + direct[0]
			nextCurrM := curr[1] + direct[1]
			if nextCurrN > -1 && nextCurrN < n && nextCurrM > -1 && nextCurrM < m && graph[nextCurrN][nextCurrM] != 1 && visited[[2]int{nextCurrN, nextCurrM}] != 1 {
				queue.Push([3]int{nextCurrN, nextCurrM, curr[2] + 1})
				visited[[2]int{nextCurrN, nextCurrM}] = 1
			}
		}
	}
	return -1
}

func falling(graph [][]int) bool {
	directs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	i1 := 0
	visited := make(map[[2]int]int)
	for j1 := 0; j1 < m; j1++ {
		if graph[i1][j1] == 1 {
			queue := Utils.Queue{}
			queue.Push([2]int{i1, j1})
			for queue.Length() > 0 {
				curr := queue.Pop().([2]int)
				if curr[1] == 0 || curr[0] == n-1 {
					return true
				}
				for _, direct := range directs {
					nextCurrN := curr[0] + direct[0]
					nextCurrM := curr[1] + direct[1]
					if nextCurrN > -1 && nextCurrN < n && nextCurrM > -1 && nextCurrM < m && graph[nextCurrN][nextCurrM] == 1 && visited[[2]int{nextCurrN, nextCurrM}] != 1 {
						queue.Push([2]int{nextCurrN, nextCurrM})
						visited[[2]int{nextCurrN, nextCurrM}] = 1
					}
				}
			}
		}
	}
	j2 := 0
	for i2 := 0; i2 < n; i2++ {
		if graph[i2][j2] == 1 {
			queue := Utils.Queue{}
			queue.Push([2]int{i2, j2})
			for queue.Length() > 0 {
				curr := queue.Pop().([2]int)
				if curr[1] == m-1 {
					return true
				}
				for _, direct := range directs {
					nextCurrN := curr[0] + direct[0]
					nextCurrM := curr[1] + direct[1]
					if nextCurrN > -1 && nextCurrN < n && nextCurrM > -1 && nextCurrM < m && graph[nextCurrN][nextCurrM] == 1 && visited[[2]int{nextCurrN, nextCurrM}] != 1 {
						queue.Push([2]int{nextCurrN, nextCurrM})
						visited[[2]int{nextCurrN, nextCurrM}] = 1
					}
				}
			}
		}
	}
	i3 := n - 1
	for j3 := 0; j3 < m; j3++ {
		if graph[i3][j3] == 1 {
			queue := Utils.Queue{}
			queue.Push([2]int{i3, j3})
			for queue.Length() > 0 {
				curr := queue.Pop().([2]int)
				if curr[1] == m-1 {
					return true
				}
				for _, direct := range directs {
					nextCurrN := curr[0] + direct[0]
					nextCurrM := curr[1] + direct[1]
					if nextCurrN > -1 && nextCurrN < n && nextCurrM > -1 && nextCurrM < m && graph[nextCurrN][nextCurrM] == 1 && visited[[2]int{nextCurrN, nextCurrM}] != 1 {
						queue.Push([2]int{nextCurrN, nextCurrM})
						visited[[2]int{nextCurrN, nextCurrM}] = 1
					}
				}
			}
		}
	}
	return false
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
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	for i := 0; i < splitAt; i++ {
		split := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		graph[y][x] = 1
	}
	for i := range n {
		for j := range m {
			if graph[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	ans := dijkstra(graph)
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
	graph := make([][]int, n)
	for i := range n {
		graph[i] = make([]int, m)
	}
	for i := range lines {
		split := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		graph[y][x] = 1
		for l := range n {
			for k := range m {
				if graph[l][k] == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		if falling(graph) {
			fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v %v. Let's submit this problem.", x, y) + Utils.Reset)
			break
		}
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	//if submit {
	//	Utils.Submit(day, 2, ans)
	//}
}
