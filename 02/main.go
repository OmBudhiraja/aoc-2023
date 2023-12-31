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

	part1(lines)
	part2(lines)

}

func part1(lines []string) {

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	result := 0

	for _, line := range lines {
		validGame := true

		split := strings.Split(line, ":")
		gameIdStr := strings.Split(split[0], " ")[1]
		gameId, _ := strconv.Atoi(gameIdStr)

		sets := strings.Split(split[1], ";")

		for _, set := range sets {
			records := strings.Split(set, ",")

			for _, record := range records {

				trimmed := strings.TrimSpace(record)
				nr, _ := strconv.Atoi(strings.Split(trimmed, " ")[0])
				color := strings.Split(trimmed, " ")[1]

				if color == "red" {
					if nr > maxRed {
						validGame = false
						break
					}
				}

				if color == "green" {
					if nr > maxGreen {
						validGame = false
						break
					}
				}

				if color == "blue" {
					if nr > maxBlue {
						validGame = false
						break
					}
				}

			}

			if !validGame {
				break
			}

		}

		if validGame {
			fmt.Println("Valid game", gameId, result)
			result += gameId
		}
	}

	fmt.Print("Part 1 -> ", result)
}

func part2(lines []string) {

	totalPower := 0

	for _, line := range lines {
		fewestRedNeeded := math.MinInt
		fewestGreenNeeded := math.MinInt
		fewestBlueNeeded := math.MinInt

		sets := strings.Split(strings.Split(line, ":")[1], ";")

		for _, set := range sets {
			records := strings.Split(set, ",")

			for _, record := range records {
				trimmed := strings.TrimSpace(record)
				nr, _ := strconv.Atoi(strings.Split(trimmed, " ")[0])
				color := strings.Split(trimmed, " ")[1]

				if color == "red" {
					if nr > fewestRedNeeded {
						fewestRedNeeded = nr
					}
				}

				if color == "green" {
					if nr > fewestGreenNeeded {
						fewestGreenNeeded = nr
					}
				}

				if color == "blue" {
					if nr > fewestBlueNeeded {
						fewestBlueNeeded = nr
					}
				}
			}
		}

		totalPower += fewestRedNeeded * fewestGreenNeeded * fewestBlueNeeded
	}

	fmt.Println("Part 2 -> ", totalPower)
}
