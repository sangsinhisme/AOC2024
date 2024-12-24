package Day16

import (
	"AOD2024/Utils"
	"container/heap"
	"fmt"
	"math"
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
	n := len(lines)
	m := len(lines[0])
	graph := make([][]uint8, n)
	for i := range graph {
		graph[i] = make([]uint8, m)
	}
	for i, line := range lines {
		for j := range line {
			graph[i][j] = line[j]
		}
	}
	ans := solve(graph)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}

const (
	East = iota
	inf  = math.MaxInt
)

type State struct {
	pos   Pos
	score int
	path  string
}

type Pos struct {
	x, y, d int
}

type PathHeap []State

func (h PathHeap) Len() int            { return len(h) }
func (h PathHeap) Less(i, j int) bool  { return h[i].score < h[j].score }
func (h PathHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PathHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *PathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func solve(maze [][]uint8) int {
	visited := map[Pos]int{}
	var paths []string

	hp := &PathHeap{}
	heap.Init(hp)

	n, m := len(maze), len(maze[0])
	start := Pos{n - 2, 1, East}
	end := Pos{x: 1, y: m - 2}

	heap.Push(hp, State{start, 0, ""})
	visited[start] = 0
	maxScore := inf

	for hp.Len() > 0 {
		s := heap.Pop(hp).(State)

		if s.score > maxScore {
			break
		}

		if scr, ok := visited[s.pos]; ok && scr < s.score {
			continue
		}

		visited[s.pos] = s.score
		if s.pos.x == end.x && s.pos.y == end.y {
			maxScore = s.score
			paths = append(paths, s.path)
		}

		dir := s.pos.d
		nx, ny := s.pos.x+dx[dir], s.pos.y+dy[dir]
		if maze[nx][ny] != '#' {
			heap.Push(hp, State{Pos{nx, ny, dir}, s.score + 1, s.path + "F"})
		}

		l, r := (dir+1)%4, ((dir-1)+4)%4
		heap.Push(hp, State{Pos{s.pos.x, s.pos.y, l}, s.score + 1000, s.path + "R"})
		heap.Push(hp, State{Pos{s.pos.x, s.pos.y, r}, s.score + 1000, s.path + "L"})
	}

	tiles := map[Pos]struct{}{start: {}}
	for _, p := range paths {
		tile := start
		dir := East
		for _, c := range p {
			switch c {
			case 'L':
				dir = (((dir - 1) % 4) + 4) % 4
			case 'R':
				dir = (((dir + 1) % 4) + 4) % 4
			case 'F':
				tile.x += dx[dir]
				tile.y += dy[dir]
				tiles[tile] = struct{}{}
			}
		}
	}

	return len(tiles)
}
