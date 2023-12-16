package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

type Slot struct {
	label       string
	focalLength int
}

func main() {
	lines := utils.Setup()

	part1Result := 0
	part2Result := 0

	boxMap := make(map[int][]Slot)

	words := strings.Split(lines[0], ",")

	for _, word := range words {
		part1Result += calcHashValue(word)
	}

	for _, word := range words {
		var boxNr int
		var label string
		var focalLength int

		if len(strings.Split(word, "=")) == 2 {
			label = strings.Split(word, "=")[0]
			focalLength, _ = strconv.Atoi(strings.Split(word, "=")[1])
			boxNr = calcHashValue(label)

			if _, ok := boxMap[boxNr]; !ok {
				boxMap[boxNr] = make([]Slot, 0)
			}

			hasLabel := utils.Filter(boxMap[boxNr], func(b Slot) bool {
				return b.label == label

			})

			if len(hasLabel) == 0 {
				boxMap[boxNr] = append(boxMap[boxNr], Slot{label: label, focalLength: focalLength})
			} else {
				for i, slot := range boxMap[boxNr] {
					if slot.label == label {
						boxMap[boxNr][i].focalLength = focalLength
					}
				}
			}

		} else {
			label = strings.Split(word, "-")[0]
			boxNr = calcHashValue(label)

			if _, ok := boxMap[boxNr]; !ok {
				boxMap[boxNr] = make([]Slot, 0)
			}

			hasLabel := utils.Filter(boxMap[boxNr], func(b Slot) bool {
				return b.label == label
			})

			if len(hasLabel) != 0 {
				// remove that index
				for i, slot := range boxMap[boxNr] {
					if slot.label == label {
						boxMap[boxNr] = append(boxMap[boxNr][:i], boxMap[boxNr][i+1:]...)
					}
				}
			}
		}

	}

	for box, slots := range boxMap {
		for i, slot := range slots {
			part2Result += (box + 1) * (i + 1) * slot.focalLength
		}
	}

	fmt.Println("Part 1:", part1Result)
	fmt.Println("Part 1:", part2Result)
}

func calcHashValue(word string) int {
	value := 0
	for _, char := range word {
		// get ascii value of char
		ascii := int(char)

		value += ascii
		value = value * 17
		value = value % 256
	}

	return value
}
