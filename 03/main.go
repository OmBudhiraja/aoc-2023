package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines, part := utils.Setup()

	if part == 1 {
		part1(lines)
	} else {
		part2(lines)
	}

}

func part1(lines []string) {
	sumOfPartNumber := 0
	for lineIndex, line := range lines {

		for charIndex := 0; charIndex < len(line); charIndex++ {
			if _, err := strconv.Atoi(string(line[charIndex])); err == nil {
				partNumber, charIndexEnd, err := getPartNumber(line, lines, lineIndex, charIndex)
				charIndex += charIndexEnd
				if err == nil {
					sumOfPartNumber += partNumber
				}

			}
		}

	}

	fmt.Println("Part 1: ", sumOfPartNumber)
}

func part2(lines []string) {}

func getPartNumber(line string, lines []string, lineIndex, charIndex int) (int, int, error) {

	partNumber := ""

	line = strings.TrimSpace(line)

	for i := charIndex; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err != nil {
			break
		}
		partNumber += string(line[i])
	}

	isPartNumberValid := false

	// check for any adjacent symbol

	if charIndex > 0 { // left
		if IsSymbol(string(line[charIndex-1])) {
			isPartNumberValid = true
		}

	}

	if charIndex+len(partNumber) < len(line) { // right
		if IsSymbol(string(line[charIndex+len(partNumber)])) {
			isPartNumberValid = true
		}
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

			if IsSymbol(string(lines[lineIndex-1][i])) {
				isPartNumberValid = true
				break
			}
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

			if IsSymbol(string(lines[lineIndex+1][i])) {
				isPartNumberValid = true
				break
			}
		}
	}

	if isPartNumberValid {
		partNr, _ := strconv.Atoi(partNumber)

		return partNr, len(partNumber), nil
	}

	return 0, len(partNumber), errors.New("invalid part number")
}

func IsSymbol(s string) bool {
	if s == "." {
		return false
	}

	if _, err := strconv.Atoi(s); err == nil {
		return false
	}

	return true
}
