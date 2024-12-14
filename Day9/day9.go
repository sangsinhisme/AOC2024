package Day9

import (
	"AOD2024/Utils"
	"fmt"
)

const day = 9

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
	ans := 0
	for _, line := range lines {
		spacesLength := 0
		filesLength := 0
		for i := 0; i < len(line); i++ {
			spacesLength += int(line[i] - '0')
			if i%2 == 0 {
				filesLength += int(line[i] - '0')
			}
		}
		currIdx := 0
		fileIdx := 0
		spaces := make([]int, spacesLength)
		files := make([]int, filesLength)
		for i := 0; i < len(line); i++ {
			currSpace := int(line[i] - '0')
			if i%2 == 0 {
				for j := 0; j < currSpace; j++ {
					spaces[currIdx] = fileIdx
					currIdx++
				}
				fileIdx++
			} else {
				for j := 0; j < currSpace; j++ {
					spaces[currIdx] = -1
					currIdx++
				}
			}

		}
		leftRight := len(spaces) - 1
		for i := 0; i < len(files); i++ {
			if spaces[i] == -1 {
				for spaces[leftRight] == -1 {
					leftRight--
				}
				files[i] = spaces[leftRight]
				leftRight--
			} else {
				files[i] = spaces[i]
			}
		}
		for i := 0; i < len(files); i++ {
			ans += i * files[i]
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
	ans := 0
	for _, line := range lines {
		spacesLength := 0
		for i := 0; i < len(line); i++ {
			spacesLength += int(line[i] - '0')
		}
		currIdx := 0
		fileIdx := 0
		spaces := make([]int, spacesLength)
		for i := 0; i < len(line); i++ {
			currSpace := int(line[i] - '0')
			if i%2 == 0 {
				for j := 0; j < currSpace; j++ {
					spaces[currIdx] = fileIdx
					currIdx++
				}
				fileIdx++
			} else {
				for j := 0; j < currSpace; j++ {
					spaces[currIdx] = -1
					currIdx++
				}
			}

		}
		i := len(spaces) - 1
		for i > 0 {
			if spaces[i] != -1 {
				j := i - 1
				for j > 0 && spaces[j] == spaces[i] {
					j--
				}
				plusI := j
				lengthReplace := i - j

				left := 0
				for left < plusI && spaces[left] != -1 {
					left++
				}

				for left < j {
					for left < j && spaces[left] != -1 {
						left++
					}
					k := left + 1
					for k < len(spaces) && spaces[k] == -1 {
						k++
					}
					lengthSpace := k - left
					if lengthSpace >= lengthReplace {
						for l := 0; l < lengthReplace; l++ {
							spaces[left] = spaces[i]
							left++
						}
						for j < i {
							spaces[j+1] = -1
							j++
						}
						break
					}
					left = k
				}
				i = plusI
			} else {
				i--
			}
		}
		for m := 0; m < len(spaces); m++ {
			if spaces[m] != -1 {
				ans += m * spaces[m]
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
