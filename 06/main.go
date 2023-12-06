package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 1

	raceTimes := filterEmptySpaces(strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " "))
	distanceRecords := filterEmptySpaces(strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " "))

	for i := 0; i < len(raceTimes); i++ {
		time, _ := strconv.Atoi(raceTimes[i])
		distance, _ := strconv.Atoi(distanceRecords[i])

		nrOfWaysToWin := 0

		for hold := 1; hold < time; hold++ {
			distanceCovered := hold * (time - hold)
			if distanceCovered > distance {
				nrOfWaysToWin++
			}
		}

		part1Result *= nrOfWaysToWin
	}

	fmt.Println("Part 1:", part1Result)
}

func filterEmptySpaces(str []string) []string {
	return utils.Filter(str, func(s string) bool {
		return s != ""
	})
}
