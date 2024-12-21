package Day16

import (
	"AOD2024/Utils"
	"container/heap"
	"fmt"
)

const day = 16
const title = "--- Day 16: Reindeer Maze ---"

func reindeerOlympics(start, end [2]int, n, m int, graph [][]uint8) int {
	directs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	pq := &Utils.PriorityQueue{}
	heap.Init(pq)
	startItem := &Utils.Item{
		Node:     start,
		Distance: 0,
		Vector:   [2]int{0, 1},
	}
	heap.Push(pq, startItem)
	visited := make(map[[2]int]map[[2]int]int)
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Utils.Item)
		if curr.Node == end {
			return curr.Distance
		}
		if _, exists := visited[curr.Node]; !exists {
			visited[curr.Node] = make(map[[2]int]int)
		}
		if minDist, exists := visited[curr.Node][curr.Vector]; exists && curr.Distance >= minDist {
			continue
		}
		visited[curr.Node][curr.Vector] = curr.Distance
		for _, direct := range directs {
			nextN, nextM := curr.Node[0]+direct[0], curr.Node[1]+direct[1]
			if nextN > -1 && nextN < n && nextM > -1 && nextM < m && graph[nextN][nextM] != '#' {
				rotate := 0
				if direct != curr.Vector {
					rotate = 1000
				}
				nextItem := &Utils.Item{
					Node:     [2]int{nextN, nextM},
					Distance: curr.Distance + 1 + rotate,
					Vector:   direct,
				}
				heap.Push(pq, nextItem)
			}
		}
	}
	return -1
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
	for i := range graph {
		graph[i] = make([]uint8, m)
	}
	start := [2]int{0, 0}
	end := [2]int{0, 0}
	for i, line := range lines {
		for j := range line {
			graph[i][j] = line[j]
			if line[j] == 'S' {
				start = [2]int{i, j}
			}
			if line[j] == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	for i := range n {
		for j := range m {
			fmt.Printf("%s", string(graph[i][j]))
		}
		fmt.Println()
	}
	ans := reindeerOlympics(start, end, n, m, graph)
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
