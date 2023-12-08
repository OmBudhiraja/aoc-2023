package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 0
	part2Result := 0

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

	}

	currentNode := nodes["AAA"]

	for {
		if currentNode == nil {
			break
		}
		next := currentNode[pattern[currentIdx]]
		currentIdx = (currentIdx + 1) % len(pattern)
		part1Result++

		if next == "ZZZ" {
			break
		} else {
			currentNode = nodes[next]
		}
	}

	currentIdx = 0
	startingNodes := []string{}

	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	fmt.Println(startingNodes)

	fonundAllWithZ := true
	for {
		fonundAllWithZ = true
		for i, node := range startingNodes {
			next := nodes[node][pattern[currentIdx]]

			startingNodes[i] = next

			if !strings.HasSuffix(next, "Z") {
				fonundAllWithZ = false
			}
		}

		currentIdx = (currentIdx + 1) % len(pattern)
		part2Result++

		if fonundAllWithZ {
			break
		}
	}

	fmt.Println("Part 1 -> ", part1Result)
	fmt.Println("Part 2 -> ", part2Result)

}
