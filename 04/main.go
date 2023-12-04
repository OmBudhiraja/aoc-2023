package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 0

	for _, line := range lines {
		matches := 0

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
		// fmt.Println("matches", matches)

		if matches > 0 {
			part1Result += int(math.Pow(2, float64(matches-1)))
		}
	}

	fmt.Println(part1Result)

}
