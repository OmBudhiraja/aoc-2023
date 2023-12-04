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

	part1Result := 0
	part2Result := 0

	cardCopiesCount := map[int]int{}

	for _, line := range lines {
		matches := 0

		cardNrStr := strings.Split(strings.Split(line, ":")[0], " ")

		cardNr, err := strconv.Atoi(cardNrStr[len(cardNrStr)-1])
		utils.CheckError(err)

		numbers := strings.TrimSpace(strings.Split(line, ":")[1])
		winningNums := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		ourNums := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")

		for _, winningNum := range winningNums {

			if winningNum == "" {
				continue
			}

			for _, ourNum := range ourNums {

				if ourNum == "" {
					continue
				}

				if winningNum == ourNum {
					matches++
					break
				}
			}

		}

		if _, ok := cardCopiesCount[cardNr]; !ok {
			cardCopiesCount[cardNr] = 1
		} else {
			cardCopiesCount[cardNr]++
		}

		if matches > 0 {
			part1Result += int(math.Pow(2, float64(matches-1)))

			for i := 1; i <= matches; i++ {
				cardCopiesCount[cardNr+i] += cardCopiesCount[cardNr]
			}
		}
	}

	for _, copies := range cardCopiesCount {
		part2Result += copies
	}

	fmt.Println("Part 1 -> ", part1Result)
	fmt.Println("Part 2 -> ", part2Result)

}
