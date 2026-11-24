package Day21

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 21
const title = "--- Day 21: Keypad Conundrum ---"

type pos struct{ r, c int }

func layout(rows []string) map[rune]pos {
	m := map[rune]pos{}
	for r, row := range rows {
		for c, ch := range row {
			m[ch] = pos{r, c}
		}
	}
	return m
}

var numeric = layout([]string{"789", "456", "123", "-0A"})
var direction = layout([]string{"-^A", "<v>"})

// bestMove: group same arrows, prefer '<' first, then '^'/'v', then '>'.
// If that order would pass the gap, use the reverse order.
func bestMove(pad map[rune]pos, from, to rune) string {
	a, b, gap := pad[from], pad[to], pad['-']
	dr, dc := b.r-a.r, b.c-a.c
	var h, v string
	if dc < 0 {
		h = strings.Repeat("<", -dc)
	} else {
		h = strings.Repeat(">", dc)
	}
	if dr < 0 {
		v = strings.Repeat("^", -dr)
	} else {
		v = strings.Repeat("v", dr)
	}
	horizontalFirst := dc < 0
	if horizontalFirst && a.r == gap.r && b.c == gap.c {
		horizontalFirst = false // corner (a.r, b.c) is the gap
	}
	if !horizontalFirst && a.c == gap.c && b.r == gap.r {
		horizontalFirst = true // corner (b.r, a.c) is the gap
	}
	if horizontalFirst {
		return h + v + "A"
	}
	return v + h + "A"
}

type memoKey struct {
	from, to rune
	depth    int
}

var memo = map[memoKey]int{}

// cost = presses on the human keypad needed to make the robot at `depth`
// move from `from` to `to` and press it.
func cost(from, to rune, depth int) int {
	if depth == 0 {
		return 1
	}
	k := memoKey{from, to, depth}
	if v, ok := memo[k]; ok {
		return v
	}
	total := 0
	prev := 'A'
	for _, ch := range bestMove(direction, from, to) {
		total += cost(prev, ch, depth-1)
		prev = ch
	}
	memo[k] = total
	return total
}

func solve(code string, robots int) int {
	total := 0
	prev := 'A'
	for _, ch := range code {
		p := 'A'
		for _, s := range bestMove(numeric, prev, ch) {
			total += cost(p, s, robots)
			p = s
		}
		prev = ch
	}
	return total
}

func resolve(lines []string, robots int) int {
	ans := 0
	for _, line := range lines {
		moves := solve(line, robots)
		fmt.Printf("Solved %s with %d moves.\n", line, moves)
		numericPart, _ := strconv.Atoi(strings.TrimSuffix(line, "A"))
		ans += moves * numericPart
	}
	return ans
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	ans := resolve(lines, 2)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	ans := resolve(lines, 25)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
