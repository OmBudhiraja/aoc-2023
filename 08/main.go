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

	stepCounts := make([]int, len(startingNodes))

	for i, startingNode := range startingNodes {
		currentNode = nodes[startingNode]
		totalSteps := 0
		currentIdx = 0

		for {

			next := currentNode[pattern[currentIdx]]
			currentIdx = (currentIdx + 1) % len(pattern)
			totalSteps++

			if strings.HasSuffix(next, "Z") {
				break
			} else {
				currentNode = nodes[next]
			}
		}

		stepCounts[i] = totalSteps
	}

	fmt.Println("Part 1 -> ", part1Result)
	fmt.Println("Part 2 -> ", LCM(stepCounts...))

}

func LCM(nums ...int) int {

	result := 1

	for _, num := range nums {
		result = (result * num) / GCD(result, num)
	}

	return result
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}
