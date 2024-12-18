package Day15

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 15
const title = "--- Day 15: Warehouse Woes ---"

func lanternfish(mapPuzzle [][]uint8, movePuzzle []uint8, init [2]int) int {
	directs := map[uint8][2]int{
		'<': {0, -1},
		'v': {1, 0},
		'>': {0, 1},
		'^': {-1, 0},
	}
	for i := range movePuzzle {
		direct := directs[movePuzzle[i]]
		nextN, nextM := init[0]+direct[0], init[1]+direct[1]
		switch mapPuzzle[nextN][nextM] {
		case '.':
			init[0] = nextN
			init[1] = nextM
		case 'O':
			for mapPuzzle[nextN][nextM] == 'O' {
				nextN, nextM = nextN+direct[0], nextM+direct[1]
			}
			if mapPuzzle[nextN][nextM] == '#' {
				continue
			} else {
				mapPuzzle[nextN][nextM] = 'O'
				init[0], init[1] = init[0]+direct[0], init[1]+direct[1]
				mapPuzzle[init[0]][init[1]] = '.'
			}
		default:
			continue
		}
	}
	ans := 0
	for i := range mapPuzzle {
		for j := range mapPuzzle[i] {
			if mapPuzzle[i][j] == 'O' {
				ans += (100 * i) + j
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
	var mapPuzzle []string
	var movePuzzle []string
	split := false
	for _, line := range lines {
		if len(line) == 0 {
			split = true
			continue
		}
		if !split {
			mapPuzzle = append(mapPuzzle, line)
		} else {
			movePuzzle = append(movePuzzle, line)
		}
	}
	var init [2]int
	graph := make([][]uint8, len(mapPuzzle))
	moves := make([]uint8, len(movePuzzle)*len(movePuzzle[0]))
	for i := range graph {
		graph[i] = make([]uint8, len(mapPuzzle[0]))
	}
	for i := range graph {
		for j := range graph[i] {
			graph[i][j] = mapPuzzle[i][j]
			if graph[i][j] == '@' {
				init = [2]int{i, j}
				graph[i][j] = '.'
			}
		}
	}
	for i, move := range movePuzzle {
		for j := range move {
			moves[i*len(move)+j] = move[j]
		}
	}
	ans := lanternfish(graph, moves, init)
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
	var mapPuzzle []string
	var movePuzzle []string
	split := false
	for _, line := range lines {
		if len(line) == 0 {
			split = true
			continue
		}
		if !split {
			mapPuzzle = append(mapPuzzle, line)
		} else {
			movePuzzle = append(movePuzzle, line)
		}
	}
	ans := 0
	fmt.Printf("Length map %v, Length movement %v", len(mapPuzzle), len(movePuzzle))
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
