package Day19

import (
	"AOD2024/Utils"
	"fmt"
	"strings"
)

const day = 19
const title = "--- Day 19: Linen Layout ---"

func possibleTowels(towels map[string]int, target string) bool {
	n := len(target)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && towels[target[j:i]] == 1 {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

func totalPossibleTowels(towels map[string]int, target string) int {
	n := len(target)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] != 0 && towels[target[j:i]] != 0 {
				dp[i] += dp[j]
			}
		}
	}
	return dp[n]
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
	towelsInput := strings.Split(lines[0], ", ")
	towelMapping := make(map[string]int)
	for i := range towelsInput {
		towelMapping[towelsInput[i]] = 1
	}
	ans := 0
	for i := 2; i < len(lines); i++ {
		fmt.Println("Process", lines[i])
		if possibleTowels(towelMapping, lines[i]) {
			ans++
		}
	}
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
	towelsInput := strings.Split(lines[0], ", ")
	towelMapping := make(map[string]int)
	for i := range towelsInput {
		towelMapping[towelsInput[i]] = 1
	}
	ans := 0
	for i := 2; i < len(lines); i++ {
		possible := totalPossibleTowels(towelMapping, lines[i])
		fmt.Println("Process", lines[i], " Possible: ", possible)
		ans += possible
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
