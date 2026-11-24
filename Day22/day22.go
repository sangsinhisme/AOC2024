package Day22

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
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

func convert(arr []int) string {
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func trace(secret int, loop int, cache map[string]int) map[string]int {
	visited := make(map[string]bool)
	prev := make([]int, 4)
	for i := range 3 {
		nextSecret := pseudorandom(secret)
		lastDigit := nextSecret % 10
		prev[i+1] = lastDigit - secret%10
		secret = nextSecret
	}
	for range loop - 3 {
		nextSecret := pseudorandom(secret)
		lastDigit := nextSecret % 10
		prev = prev[1:4]
		prev = append(prev, lastDigit-secret%10)
		prevString := convert(prev)
		if !visited[prevString] {
			cache[prevString] += lastDigit
			visited[prevString] = true
		}
		secret = nextSecret
	}
	return cache
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	ans := 0
	cache := make(map[string]int)
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		cache = trace(secret, 2000, cache)
	}
	for _, value := range cache {
		ans = max(ans, value)
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
