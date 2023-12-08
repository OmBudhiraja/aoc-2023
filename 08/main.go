package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 0

	nodes := map[string][]string{}

	pattern := utils.Mapper(strings.Split(lines[0], ""), func(s string) int {
		if s == "L" {
			return 0
		} else {
			return 1
		}
	})
	currentIdx := 0

	for _, line := range lines[2:] {
		line = strings.TrimSpace(line)
		node := strings.Split(line, " = ")[0]
		values := strings.Split(line, " = ")[1]

		values = strings.TrimLeft(values, "(")
		values = strings.TrimRight(values, ")")

		nodes[node] = strings.Split(values, ", ")

		// parts := strings.Split(line, " ")
		// nodes[parts[0]] = append(nodes[parts[0]], parts[2])
	}

	currentNode := nodes["AAA"]

	for {
		next := currentNode[pattern[currentIdx]]
		currentIdx = (currentIdx + 1) % len(pattern)
		part1Result++

		if next == "ZZZ" {
			break
		} else {
			currentNode = nodes[next]
		}
	}

	fmt.Println(part1Result)

}
