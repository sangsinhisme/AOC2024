package Day23

import (
	"AOD2024/Utils"
	"fmt"
	"slices"
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

func findCircleGraph(origin []string, edges map[[2]string]bool) []string {
	var best, cur []string

	var dfs func(start int)
	dfs = func(start int) {
		if len(cur) > len(best) {
			best = slices.Clone(cur)
		}
		for j := start; j < len(origin); j++ {
			if len(cur)+len(origin)-j <= len(best) {
				return // can't beat best even taking everything left
			}
			ok := true
			for _, c := range cur {
				if !edges[[2]string{c, origin[j]}] {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			cur = append(cur, origin[j])
			dfs(j + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return best
}

func findGraph(edges map[string][]string, nodes []string, edgesMap map[[2]string]bool) string {
	visited := make(map[string]bool)
	maxSolve := 0
	var solved []string
	for _, node := range nodes {
		if !visited[node] {
			queue := Utils.NewQueue()
			queue.Push(node)
			visited[node] = true
			var graph []string
			for queue.Length() > 0 {
				curr := queue.Pop().(string)
				graph = append(graph, curr)
				for _, nextNode := range edges[curr] {
					if !visited[nextNode] {
						queue.Push(nextNode)
						visited[nextNode] = true
					}
				}
			}
			maxCircle := findCircleGraph(graph, edgesMap)
			if len(maxCircle) > maxSolve {
				solved = maxCircle
			}
		}
	}
	slices.Sort(solved)
	return strings.Join(solved, ",")
}

func Part2(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	edges := map[string][]string{}
	nodesMap := map[string]bool{}
	edgesMap := map[[2]string]bool{}
	for _, line := range lines {
		machines := strings.Split(line, "-")
		edgesMap[[2]string{machines[0], machines[1]}] = true
		edgesMap[[2]string{machines[1], machines[0]}] = true
		edges[machines[0]] = append(edges[machines[0]], machines[1])
		edges[machines[1]] = append(edges[machines[1]], machines[0])
		nodesMap[machines[0]] = true
		nodesMap[machines[1]] = true
	}
	nodes := make([]string, len(nodesMap))
	i := 0
	for node := range nodesMap {
		nodes[i] = node
		i += 1
	}
	ans := findGraph(edges, nodes, edgesMap)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, len(ans))
	}
}
