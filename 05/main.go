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
	part2Result := math.MaxInt64

	seeds := utils.Mapper(strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " "), func(s string) int {
		n, err := strconv.Atoi(s)
		utils.CheckError(err)
		return n
	})

	conversions := make([]int, 0)

	conversions = append(conversions, seeds...)

	groupedLines := strings.Split(strings.Join(lines[2:], "\n"), "\n\n")

	for _, data := range groupedLines {
		lines := strings.Split(data, "\n")

		for seedIndex := range seeds {
			for lineIndex, line := range lines {
				if lineIndex == 0 {
					continue
				}
				conversionData := strings.Split(strings.TrimSpace(line), " ")
				destinationRangeStart, _ := strconv.Atoi(conversionData[0])
				sourceRangeStart, _ := strconv.Atoi(conversionData[1])
				rangeLength, _ := strconv.Atoi(conversionData[2])

				if sourceRangeStart <= conversions[seedIndex] && conversions[seedIndex] <= sourceRangeStart+rangeLength-1 {
					conversions[seedIndex] = conversions[seedIndex] - sourceRangeStart + destinationRangeStart
					break
				} else {
					continue
				}
			}
		}

	}

	for i, seed := range seeds {
		if i%2 != 0 {
			continue
		}

		for j := seed; j < seed+seeds[i+1]; j++ {

			finalConvertedValue := j

			for _, data := range groupedLines {
				lines := strings.Split(data, "\n")

				for lineIndex, line := range lines {
					if lineIndex == 0 {
						continue
					}

					conversionData := strings.Split(strings.TrimSpace(line), " ")
					destinationRangeStart, _ := strconv.Atoi(conversionData[0])
					sourceRangeStart, _ := strconv.Atoi(conversionData[1])
					rangeLength, _ := strconv.Atoi(conversionData[2])

					if sourceRangeStart <= finalConvertedValue && finalConvertedValue <= sourceRangeStart+rangeLength-1 {
						finalConvertedValue = finalConvertedValue - sourceRangeStart + destinationRangeStart
						break
					} else {
						continue
					}
				}

			}

			if finalConvertedValue < part2Result {
				part2Result = finalConvertedValue
			}
		}
	}

	for _, conversion := range conversions {
		if conversion < part1Result {
			part1Result = conversion
		}
	}

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)

}
