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

func gardenTravelPart2(n, m int, graph [][]int) int {
	visited := make(map[[2]int]bool)
	directs := [][4]int{
		{0, 1, '|', 'r'},
		{0, -1, '|', 'l'},
		{1, 0, '-', 'd'},
		{-1, 0, '-', 'u'},
	}
	corner := [][2]int{
		{-1, -1},
		{1, 1},
		{-1, 1},
		{1, -1},
	}
	var helper func(i, j int) int
	helper = func(i, j int) int {
		perimeterGraph := make([][]int, n*2+1)
		perimeterMap := make(map[[2]int]int)
		for z := range n*2 + 1 {
			perimeterGraph[z] = make([]int, m*2+1)
		}
		queue := Utils.NewQueue()
		queue.Push([2]int{i, j})
		visited[[2]int{i, j}] = true
		plant := graph[i][j]
		perimeter := 1
		for queue.Length() > 0 {
			curr := queue.Pop().([2]int)
			perimeterGraph[curr[0]*2+1][curr[1]*2+1] = plant
			for _, direct := range directs {
				nextI := curr[0] + direct[0]
				nextJ := curr[1] + direct[1]
				if nextI > -1 && nextI < n && nextJ > -1 && nextJ < m && graph[nextI][nextJ] == plant {
					if visited[[2]int{nextI, nextJ}] != true {
						perimeter++
						visited[[2]int{nextI, nextJ}] = true
						queue.Push([2]int{nextI, nextJ})
					}
				} else {
					perimeterMap[[2]int{curr[0]*2 + 1 + direct[0], curr[1]*2 + 1 + direct[1]}] = direct[3]
					perimeterGraph[curr[0]*2+1+direct[0]][curr[1]*2+1+direct[1]] = direct[2]
				}
			}
			for _, direct := range corner {
				nextI := curr[0]*2 + 1 + direct[0]
				nextJ := curr[1]*2 + 1 + direct[1]
				perimeterGraph[nextI][nextJ] = '+'
			}
		}
		region := 0
		visitedRegion := make(map[[2]int]bool)
		for l := range n*2 + 1 {
			for k := range m*2 + 1 {
				if perimeterGraph[l][k] == '-' {
					if visitedRegion[[2]int{l, k}] != true {
						visitedRegion[[2]int{l, k}] = true
						vectorMap := perimeterMap[[2]int{l, k}]
						region++
						nextK := k + 1
						for nextK < m*2+1 && (perimeterGraph[l][nextK] == '+' || perimeterGraph[l][nextK] == '-' && perimeterMap[[2]int{l, nextK}] == vectorMap) {
							visitedRegion[[2]int{l, nextK}] = true
							nextK++
						}
					}
				} else if perimeterGraph[l][k] == '|' {
					if visitedRegion[[2]int{l, k}] != true {
						visitedRegion[[2]int{l, k}] = true
						vectorMap := perimeterMap[[2]int{l, k}]
						region++
						nextL := l + 1
						for nextL < n*2+1 && (perimeterGraph[nextL][k] == '+' || perimeterGraph[nextL][k] == '|' && perimeterMap[[2]int{nextL, k}] == vectorMap) {
							visitedRegion[[2]int{nextL, k}] = true
							nextL++
						}
					}
				}
			}
		}
		for l := range n*2 + 1 {
			for k := range m*2 + 1 {
				if perimeterGraph[l][k] == 0 {
					fmt.Print(".")
				} else {
					fmt.Printf("%s", string(perimeterGraph[l][k]))
				}
			}
			fmt.Println()
		}
		fmt.Printf("PLANT %v, Region %v, Perimeter %v", string(graph[i][j]), region, perimeter)
		fmt.Println()
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
	ans := gardenTravelPart2(n, m, graph)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
