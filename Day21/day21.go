package Day21

import (
	"AOD2024/Utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = 21
const title = "--- Day 21: Keypad Conundrum ---"

func makeShortedPath(keyboard [][]rune, order []rune) map[[2]rune]string {
	moveMap := map[rune][2]int{
		'>': {0, 1},
		'v': {1, 0},
		'^': {-1, 0},
		'<': {0, -1},
	}
	shorted := map[[2]rune]string{}
	for i := range len(keyboard) {
		for j := range len(keyboard[1]) {
			start := keyboard[i][j]
			queue := Utils.NewQueue()
			queue.Push([3]any{i, j, ""})
			visited := map[rune]bool{}
			visited[start] = true
			for queue.Length() > 0 {
				curr := queue.Pop().([3]any)
				currI, currJ := curr[0].(int), curr[1].(int)
				moveString := curr[2].(string)

				for _, moveChar := range order {
					move := moveMap[moveChar]
					nextI, nextJ := currI+move[0], currJ+move[1]
					if nextI >= 0 && nextI < len(keyboard) && nextJ >= 0 && nextJ < len(keyboard[1]) {
						nextNumeric := keyboard[nextI][nextJ]
						if nextNumeric != '-' && !visited[nextNumeric] {
							nextMove := moveString + string(moveChar)
							queue.Push([3]any{nextI, nextJ, nextMove})
							shorted[[2]rune{start, nextNumeric}] = nextMove
							visited[nextNumeric] = true
						}
					}
				}
			}
		}
	}
	return shorted
}

func convertCode(line string, shorted map[[2]rune]string) string {
	line = "A" + line
	var findPath strings.Builder
	for i := range (len(line)) - 1 {
		findPath.WriteString(shorted[[2]rune{rune(line[i]), rune(line[i+1])}] + "A")
	}
	return findPath.String()
}

func solvePart1(line string, orderNumeric []rune) int {
	numeric := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'-', '0', 'A'},
	}
	direction := [][]rune{
		{'-', '^', 'A'},
		{'<', 'v', '>'},
	}
	shortedNumeric := makeShortedPath(numeric, orderNumeric)
	shortedDirection := makeShortedPath(direction, orderNumeric)
	firstConv := convertCode(line, shortedNumeric)
	secondConv := convertCode(firstConv, shortedDirection)
	thirdConv := convertCode(secondConv, shortedDirection)
	return len(thirdConv)
}

func permutations(order []rune) [][]rune {
	var res [][]rune
	var cur []rune
	used := make([]bool, len(order))

	var dfs func()
	dfs = func() {
		if len(cur) == len(order) {
			res = append(res, append([]rune(nil), cur...)) // copy, cur gets reused
			return
		}
		for i, r := range order {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, r)
			dfs()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	order := []rune{'>', 'v', '^', '<'}
	ans := 0
	for _, line := range lines {
		shorted := math.MaxInt
		for _, p := range permutations(order) {
			shorted = min(solvePart1(line, p), shorted)
		}
		fmt.Printf("Solved %s with %d moves.\n", line, shorted)
		numericPart, _ := strconv.Atoi(strings.ReplaceAll(line, "A", ""))
		ans += shorted * numericPart
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	n := len(lines)
	m := len(lines[0])
	graph := make([][]uint8, n)
	for i := range n {
		graph[i] = make([]uint8, m)
	}
	ans := 0
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
