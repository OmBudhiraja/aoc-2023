package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := math.MaxInt64

	seeds := utils.Mapper(strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " "), func(s string) int {
		n, err := strconv.Atoi(s)
		utils.CheckError(err)
		return n
	})

	conversions := seeds

	groupedLines := strings.Split(strings.Join(lines[2:], "\n"), "\n\n")

	for _, conversionMap := range groupedLines {
		lines := strings.Split(conversionMap, "\n")

		for seedIndex, seed := range seeds {
			for lineIndex, line := range lines {
				if lineIndex == 0 {
					continue
				}
				conversionData := strings.Split(strings.TrimSpace(line), " ")
				destinationRangeStart, _ := strconv.Atoi(conversionData[0])
				sourceRangeStart, _ := strconv.Atoi(conversionData[1])
				rangeLength, _ := strconv.Atoi(conversionData[2])

				if sourceRangeStart <= seed && seed <= sourceRangeStart+rangeLength-1 {
					// found = true
					conversions[seedIndex] = seed - sourceRangeStart + destinationRangeStart
				} else {
					continue
				}
			}
		}

	}

	for _, conversion := range conversions {
		if conversion < part1Result {
			part1Result = conversion
		}
	}

	fmt.Println("Part 1:", part1Result)

}
