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
		case '[':
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
		case '#':
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

func lanternfishTwiceWide(mapPuzzle [][]uint8, movePuzzle []uint8, init [2]int) int {
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
		case '#':
			continue
		default:
			if direct[0] == 0 {
				for mapPuzzle[nextN][nextM] == '[' || mapPuzzle[nextN][nextM] == ']' {
					nextN, nextM = nextN+direct[0], nextM+direct[1]
				}
				if mapPuzzle[nextN][nextM] == '#' {
					continue
				} else {
					tempN, tempM := nextN, nextM
					for tempN != init[0] || tempM != init[1] {
						backN, backM := tempN-direct[0], tempM-direct[1]
						mapPuzzle[tempN][tempM] = mapPuzzle[backN][backM]
						tempN, tempM = backN, backM
					}
				}
				init[0], init[1] = init[0]+direct[0], init[1]+direct[1]
			} else {
				var blocks [][2]int
				if mapPuzzle[nextN][nextM] == '[' {
					blocks = [][2]int{{nextN, nextM}, {nextN, nextM + 1}}
				} else {
					blocks = [][2]int{{nextN, nextM}, {nextN, nextM - 1}}
				}
				for block := range blocks {
					fmt.Println(block)
				}
			}
		}
		for l := range mapPuzzle {
			for k := range mapPuzzle[0] {
				fmt.Printf("%s", string(mapPuzzle[l][k]))
			}
			fmt.Println()
		}
	}
	return 0
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
	var init [2]int
	graph := make([][]uint8, len(mapPuzzle))
	moves := make([]uint8, len(movePuzzle)*len(movePuzzle[0]))
	for i := range graph {
		graph[i] = make([]uint8, len(mapPuzzle[0])*2)
	}
	for i := range len(mapPuzzle) {
		for j := range len(mapPuzzle[0]) {
			switch mapPuzzle[i][j] {
			case '#':
				graph[i][j*2] = '#'
				graph[i][j*2+1] = '#'
			case '.':
				graph[i][j*2] = '.'
				graph[i][j*2+1] = '.'
			case '@':
				graph[i][j*2] = '.'
				graph[i][j*2+1] = '.'
				init[0] = i
				init[1] = j * 2
			case 'O':
				graph[i][j*2] = '['
				graph[i][j*2+1] = ']'
			}
		}
	}
	for i, move := range movePuzzle {
		for j := range move {
			moves[i*len(move)+j] = move[j]
		}
	}
	for i := range graph {
		for j := range graph[0] {
			fmt.Printf("%s", string(graph[i][j]))
		}
		fmt.Println()
	}
	ans := lanternfishTwiceWide(graph, moves, init)
	fmt.Printf("Length map %v, Length movement %v", len(mapPuzzle), len(movePuzzle))
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
