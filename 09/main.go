package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

func main() {
	lines := utils.Setup()

	part1Result := 0

	for _, line := range lines {

		seqs := make([][]int, 0)

		nums := utils.Mapper(strings.Split(line, " "), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})

		seqs = append(seqs, nums)
		currentSeq := 0

		for {
			diff := make([]int, len(seqs[currentSeq])-1)
			for i := 1; i < len(seqs[currentSeq]); i++ {
				diff[i-1] = seqs[currentSeq][i] - seqs[currentSeq][i-1]
			}

			seqs = append(seqs, diff)
			currentSeq++

			if utils.Every(diff, func(n int) bool {
				return n == diff[0]
			}) {
				break
			}
		}

		for i := len(seqs) - 1; i > 0; i-- {
			val := seqs[i][len(seqs[i])-1] + seqs[i-1][len(seqs[i-1])-1]
			seqs[i-1] = append(seqs[i-1], val)
		}

		part1Result += seqs[0][len(seqs[0])-1]

	}

	fmt.Println("Part 1:", part1Result)
}
