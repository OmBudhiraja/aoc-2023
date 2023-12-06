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
	part2Result := 0

	raceTimes := filterEmptySpaces(strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " "))
	distanceRecords := filterEmptySpaces(strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " "))

	joinedRaceTime, _ := strconv.Atoi(strings.Join(raceTimes, ""))
	joinedDistanceRecord, _ := strconv.Atoi(strings.Join(distanceRecords, ""))

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

	for hold := 1; hold < joinedRaceTime; hold++ {
		distanceCovered := hold * (joinedRaceTime - hold)
		if distanceCovered > joinedDistanceRecord {
			part2Result++
		}
	}

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 2:", part2Result)
}

func filterEmptySpaces(str []string) []string {
	return utils.Filter(str, func(s string) bool {
		return s != ""
	})
}
