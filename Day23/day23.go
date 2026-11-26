package Day23

import (
	"AOD2024/Utils"
	"fmt"
	"strings"
)

const day = 23
const title = "--- Day 23: LAN Party ---\n"

type nodeTract struct {
	node    string
	preNode int
	lastT   int
}

func LAN(edges map[[2]string]bool, nodes []string) int {
	solves := 0
	for i := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			for k := j + 1; k < len(nodes); k++ {
				if nodes[i][0] == 't' || nodes[j][0] == 't' || nodes[k][0] == 't' {
					if edges[[2]string{nodes[i], nodes[j]}] && edges[[2]string{nodes[i], nodes[k]}] && edges[[2]string{nodes[j], nodes[k]}] {
						solves += 1
					}
				}
			}
		}
	}
	return solves
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	edges := map[[2]string]bool{}
	nodesMap := map[string]bool{}
	for _, line := range lines {
		machines := strings.Split(line, "-")
		edges[[2]string{machines[0], machines[1]}] = true
		edges[[2]string{machines[1], machines[0]}] = true
		nodesMap[machines[0]] = true
		nodesMap[machines[1]] = true
	}
	nodes := make([]string, len(nodesMap))
	i := 0
	for node := range nodesMap {
		nodes[i] = node
		i += 1
	}
	ans := LAN(edges, nodes)
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
