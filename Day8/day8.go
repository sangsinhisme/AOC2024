package Day8

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 8

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
	antenna := make(map[uint8][][2]int)
	n := len(lines)
	m := len(lines[0])
	for i := range n {
		for j := range m {
			if lines[i][j] != '.' {
				if antenna[lines[i][j]] == nil {
					antenna[lines[i][j]] = [][2]int{{i, j}}
				} else {
					antenna[lines[i][j]] = append(antenna[lines[i][j]], [2]int{i, j})
				}
			}
		}
	}
	ans := 0
	antinodes := make(map[[2]int]int)
	for _, value := range antenna {
		for i := 0; i < len(value); i++ {
			antenaI := value[i]
			for j := i + 1; j < len(value); j++ {
				antenaJ := value[j]
				freqI, freqJ := antenaI[0]-antenaJ[0], antenaI[1]-antenaJ[1]
				leftI, leftJ := antenaI[0]+freqI, antenaI[1]+freqJ
				rightI, rightJ := antenaJ[0]-freqI, antenaJ[1]-freqJ
				if leftI > -1 && leftI < n && leftJ > -1 && leftJ < m && antinodes[[2]int{leftI, leftJ}] != 1 {
					ans++
					antinodes[[2]int{leftI, leftJ}] = 1
				}
				if rightI > -1 && rightI < n && rightJ > -1 && rightJ < m && antinodes[[2]int{rightI, rightJ}] != 1 {
					ans++
					antinodes[[2]int{rightI, rightJ}] = 1
				}
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
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
	antenna := make(map[uint8][][2]int)
	n := len(lines)
	m := len(lines[0])
	for i := range n {
		for j := range m {
			if lines[i][j] != '.' {
				if antenna[lines[i][j]] == nil {
					antenna[lines[i][j]] = [][2]int{{i, j}}
				} else {
					antenna[lines[i][j]] = append(antenna[lines[i][j]], [2]int{i, j})
				}
			}
		}
	}
	ans := 0
	antinodes := make(map[[2]int]int)
	for _, value := range antenna {
		for i := 0; i < len(value); i++ {
			antenaI := value[i]
			for j := i + 1; j < len(value); j++ {
				antenaJ := value[j]
				freqI, freqJ := antenaI[0]-antenaJ[0], antenaI[1]-antenaJ[1]
				leftI, leftJ := antenaI[0]+freqI, antenaI[1]+freqJ
				rightI, rightJ := antenaJ[0]-freqI, antenaJ[1]-freqJ
				for leftI > -1 && leftI < n && leftJ > -1 && leftJ < m {
					if antinodes[[2]int{leftI, leftJ}] != 1 {
						ans++
						antinodes[[2]int{leftI, leftJ}] = 1
					}
					leftI, leftJ = leftI+freqI, leftJ+freqJ
				}
				for rightI > -1 && rightI < n && rightJ > -1 && rightJ < m {
					if antinodes[[2]int{rightI, rightJ}] != 1 {
						ans++
						antinodes[[2]int{rightI, rightJ}] = 1
					}
					rightI, rightJ = rightI-freqI, rightJ-freqJ
				}
			}
		}
	}
	for _, value := range antenna {
		if len(value) > 1 {
			for i := 0; i < len(value); i++ {
				if antinodes[value[i]] != 1 {
					ans++
				}
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
