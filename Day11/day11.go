package Day11

import (
	"AOD2024/Utils"
	"fmt"
	"strconv"
	"strings"
)

const day = 11
const title = "--- Day 11: Plutonian Pebbles ---"
const MaxDepth = 75

var cache *Utils.LRUCache

func init() {
	cache = Utils.NewLRUCache(10000)
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
	split := strings.Split(lines[0], " ")
	init := make([]int, len(split))
	for i, elem := range split {
		init[i], _ = strconv.Atoi(elem)
	}
	for i := 0; i < 25; i++ {
		var stones []int
		for _, stone := range init {
			lengthStone := Utils.NumDigits(stone)
			if stone == 0 {
				stones = append(stones, 1)
			} else if lengthStone%2 == 0 {
				initString := strconv.Itoa(stone)
				left, _ := strconv.Atoi(initString[:lengthStone/2])
				right, _ := strconv.Atoi(initString[lengthStone/2:])
				stones = append(stones, left)
				stones = append(stones, right)
			} else {
				stones = append(stones, stone*2024)
			}
		}
		init = stones
	}
	ans := len(init)
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
	split := strings.Split(lines[0], " ")
	stones := make([]int, len(split))
	for i, elem := range split {
		stones[i], _ = strconv.Atoi(elem)
	}
	ans := 0
	for _, stone := range stones {
		ans += expandStone(stone, 0)
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}

func expandStone(stone, depth int) int {
	// Create a unique key for this (stone, depth) pair
	cacheKey := fmt.Sprintf("%d,%d", stone, depth)
	// Check if the result is already cached
	if result, found := cache.Get(cacheKey); found {
		return result
	}
	for depth < MaxDepth {
		if stone == 0 {
			stone = 1
			depth++
			continue
		}
		lengthStone := Utils.NumDigits(stone)
		if lengthStone%2 == 0 {
			divisor := 1
			for j := 0; j < lengthStone/2; j++ {
				divisor *= 10
			}
			left := stone / divisor
			right := stone % divisor
			result := expandStone(left, depth+1) + expandStone(right, depth+1)
			cache.Put(cacheKey, result)
			return result
		}
		stone = stone * 2024
		depth += 1
	}
	cache.Put(cacheKey, 1)
	return 1
}
