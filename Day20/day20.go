package Day20

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 20
const title = "--- Day 20: Race Condition ---"

func racingCheat(n, m int, graph [][]uint8, start, end [2]int) int {
	directs := [][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	queue := Utils.Queue{}
	visited := make(map[[2]int]int)
	raceLength := 0
	queue.Push(start)
	visited[start] = 0
	var history [][2]int
	for queue.Length() > 0 {
		curr := queue.Pop().([2]int)
		if curr == end {
			break
		}
		raceLength++
		for _, direct := range directs {
			nextN, nextM := curr[0]+direct[0], curr[1]+direct[1]
			if nextN > -1 && nextN < n && nextM > -1 && nextM < m && graph[nextN][nextM] != '#' && visited[[2]int{nextN, nextM}] == 0 {
				queue.Push([2]int{nextN, nextM})
				visited[[2]int{nextN, nextM}] = raceLength
				history = append(history, [2]int{nextN, nextM})
			}
		}
	}
	ans := 0
	for i := range history {
		curr := history[i]
		for _, direct := range directs {
			next1N, next1M := curr[0]+direct[0], curr[1]+direct[1]
			nextN, nextM := curr[0]+direct[0]*2, curr[1]+direct[1]*2
			if nextN > -1 && nextN < n && nextM > -1 && nextM < m && graph[nextN][nextM] != '#' && visited[[2]int{nextN, nextM}] != 0 {
				if graph[next1N][next1M] == '#' {
					currCost := visited[[2]int{curr[0], curr[1]}]
					nextCost := visited[[2]int{nextN, nextM}]
					if nextCost-currCost-2 > 99 {
						ans++
					}
					fmt.Printf("Possible cheat node [%v, %v] reduce from %v to %v", curr[0], curr[1], currCost, nextCost)
					fmt.Println()
				}
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
	graph := make([][]uint8, n)
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for i := range n {
		for j := range n {
			graph[i][j] = lines[i][j]
			if graph[i][j] == 'S' {
				start = [2]int{i, j}
			}
			if graph[i][j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	ans := racingCheat(n, m, graph, start, end)
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
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
