package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/ombudhiraja/aoc-2023/utils"
)

type Point struct {
	i int
	j int
}

func main() {
	lines := utils.Setup()

	part1(lines)
	part2(lines)

}

func part1(lines []string) {
	sumOfPartNumber := 0
	for lineIndex, line := range lines {

		for charIndex := 0; charIndex < len(line); charIndex++ {
			if _, err := strconv.Atoi(string(line[charIndex])); err == nil {
				partNumber, charIndexEnd, _, err := getPartNumber(line, lines, lineIndex, charIndex)
				charIndex += charIndexEnd
				if err == nil {
					sumOfPartNumber += partNumber
				}

			}
		}

	}

	fmt.Println("Part 1: ", sumOfPartNumber)
}

func part2(lines []string) {

	sumOfGearRatios := 0
	gearMap := make(map[Point][]int)
	var gearPos *Point
	partNumber := 0

	for lineIndex, line := range lines {
		for charIndex := 0; charIndex < len(line); charIndex++ {

			if unicode.IsDigit(rune(line[charIndex])) {
				nr, charIndexEnd, p, err := getPartNumber(line, lines, lineIndex, charIndex)
				charIndex += charIndexEnd

				if err == nil {
					gearPos = p
					partNumber = nr
				}
			} else {
				if gearPos != nil {
					if _, ok := gearMap[*gearPos]; !ok {
						gearMap[*gearPos] = make([]int, 0)
					}

					gearMap[*gearPos] = append(gearMap[*gearPos], partNumber)
				}

				gearPos = nil
				partNumber = 0
			}

		}

	}

	for _, gear := range gearMap {
		if len(gear) == 2 {
			sumOfGearRatios += gear[0] * gear[1]
		}
	}

	fmt.Println("Part 2: ", sumOfGearRatios)
}

func getPartNumber(line string, lines []string, lineIndex, charIndex int) (int, int, *Point, error) {

	partNumber := ""
	var gearPos *Point

	for i := charIndex; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			break
		}
		partNumber += string(line[i])
	}

	isPartNumberValid := false

	neighbours := make([]Point, 0)

	// check for any adjacent symbol
	if charIndex > 0 { // left
		neighbours = append(neighbours, Point{lineIndex, charIndex - 1})
	}

	if charIndex+len(partNumber) < len(line) { // right
		neighbours = append(neighbours, Point{lineIndex, charIndex + len(partNumber)})
	}

	if lineIndex > 0 { // upper line
		start := charIndex
		end := charIndex + len(partNumber)

		if charIndex > 0 {
			start--
		}

		if charIndex+len(partNumber) < len(line) {
			end++
		}

		for i := start; i < end; i++ {
			neighbours = append(neighbours, Point{i: lineIndex - 1, j: i})
		}
	}

	if lineIndex < len(lines)-1 { // lower line
		start := charIndex
		end := charIndex + len(partNumber)

		if charIndex > 0 {
			start--
		}

		if charIndex+len(partNumber) < len(line) {
			end++
		}

		for i := start; i < end; i++ {
			neighbours = append(neighbours, Point{i: lineIndex + 1, j: i})
		}
	}

	for _, neighbour := range neighbours {
		char := lines[neighbour.i][neighbour.j]

		if IsSymbol(rune(char)) {
			isPartNumberValid = true
			if string(char) == "*" {
				gearPos = &Point{
					i: neighbour.i,
					j: neighbour.j,
				}
			}
			break
		}
	}

	if isPartNumberValid {
		partNr, _ := strconv.Atoi(partNumber)

		return partNr, len(partNumber) - 1, gearPos, nil
	}

	return 0, len(partNumber) - 1, gearPos, errors.New("invalid part number")
}

func IsSymbol(s rune) bool {
	if s == '.' || unicode.IsDigit((s)) {
		return false
	}

	return true
}
