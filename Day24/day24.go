package Day24

import (
	"AOD2024/Utils"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

const day = 24
const title = "--- Day 24: Crossed Wires ---\n\n"

type Wire struct {
	value int
}

type Gate struct {
	input [2]string
	kind  Kind // 0: AND, 1: OR, 2: XOR
}

type Kind int

const (
	AND Kind = iota
	OR
	XOR
)

var kinds = map[string]Kind{
	"AND": AND,
	"OR":  OR,
	"XOR": XOR,
}

func findWireValue(wires map[string]int, gates map[string]Gate, name string) int {
	if wire, ok := wires[name]; ok {
		return wire
	}
	var out int
	gate := gates[name]
	a := findWireValue(wires, gates, gate.input[0])
	b := findWireValue(wires, gates, gate.input[1])

	switch gate.kind {
	case AND:
		out = a & b
	case OR:
		out = a | b
	case XOR:
		out = a ^ b
	default:
		panic("bad gate kind")
	}
	wires[name] = out
	return wires[name]
}

func Part1(submit bool) {
	lines, _ := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	wires := map[string]int{}
	gates := map[string]Gate{}
	wiresZ := map[string]bool{}
	i := 0
	for i < len(lines) {
		if lines[i] == "" {
			i += 1
			break
		}
		line := strings.Split(lines[i], ": ")
		value, _ := strconv.Atoi(line[1])
		wires[line[0]] = value
		i += 1
	}
	for i < len(lines) {
		line := strings.Split(lines[i], " ")
		gates[line[4]] = Gate{input: [2]string{line[0], line[2]}, kind: kinds[line[1]]}
		if line[4][0] == 'z' {
			wiresZ[line[4]] = true
		}
		i += 1
	}

	keys := slices.Collect(maps.Keys(wiresZ))
	slices.Sort(keys)
	binary := ""
	for _, wire := range keys {
		binary = strconv.Itoa(findWireValue(wires, gates, wire)) + binary
	}
	ans, _ := strconv.ParseInt(binary, 2, 64)
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, int(ans))
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
