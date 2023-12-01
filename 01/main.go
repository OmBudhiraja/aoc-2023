package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func setup() (string, int) {
	fileName := flag.String("f", "input.txt", "-f <filename>")
	part := flag.Int("p", 1, "-p <partNr>")
	flag.Parse()

	return *fileName, *part
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename, part := setup()

	data, err := os.ReadFile(filename)
	checkError(err)

	lines := strings.Split(string(data), "\n")

	if part == 1 {
		part1(lines)
	} else {
		part2(lines)
	}

}

func part1(lines []string) {
	result := 0

	for _, line := range lines {
		chars := strings.Split(line, "")
		digit := make([]*int, 2)
		for _, char := range chars {
			val, err := strconv.Atoi(char)
			if err == nil {
				if digit[0] == nil {
					digit[0] = &val
				} else {
					digit[1] = &val
				}
			}
		}

		// join the two digits [1, 2] -> 12
		result += *digit[0] * 10

		if digit[1] == nil {
			result += *digit[0]
		} else {
			result += *digit[1]
		}
	}

	fmt.Println(result)
}
func part2(lines []string) {

}
