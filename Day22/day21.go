package Day22

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
)

const day = 22
const title = "--- Day 22: Monkey Market ---\n"

const pruneModule = 16777216

func mix(secret, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % pruneModule
}

func pseudorandom(secret int) int {
	secret = mix(secret, secret*64)
	secret = prune(secret)
	secret = mix(secret, secret/32)
	secret = prune(secret)
	secret = prune(mix(secret, secret*2048))
	return secret
}

func solvePart1(secret, loop int) int {
	for range loop {
		secret = pseudorandom(secret)
	}
	return secret
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	ans := 0
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		solved := solvePart1(secret, 2000)
		ans += solved
		fmt.Printf("Pseudorandom %d: %d\n", secret, solved)
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	ans := len(lines)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
