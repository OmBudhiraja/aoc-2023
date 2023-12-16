package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 0

	words := strings.Split(lines[0], ",")

	for _, word := range words {
		value := 0
		for _, char := range word {
			// get ascii value of char
			ascii := int(char)

			value += ascii
			value = value * 17
			value = value % 256
		}

		part1Result += value
	}

	fmt.Println("Part 1:", part1Result)
}
