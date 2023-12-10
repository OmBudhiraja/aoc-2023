package main

import (
	"fmt"
	"strings"

	"github.com/ombudhiraja/aoc-2023/utils"
)

type Point struct {
	x int
	y int
}

func main() {
	lines := utils.Setup()

	grid := utils.Mapper(lines, func(s string) []string {
		return strings.Split(s, "")
	})

	startingPoint := getStartingPoint(grid)
	visited := map[Point]int{startingPoint: 0}
	notChecked := []Point{startingPoint}

	maxDist := 0
	for len(notChecked) > 0 {
		current := notChecked[0]
		notChecked = notChecked[1:]
		next := nextPoints(grid, current)
		for _, point := range next {

			if !isIndexInGrid(grid, point) {
				continue
			}

			if _, found := visited[point]; !found {
				visited[point] = visited[current] + 1
				maxDist = max(maxDist, visited[current]+1)
				notChecked = append(notChecked, point)
			}
		}
	}

	fmt.Println("Part 1:", maxDist)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getStartingPoint(grid [][]string) Point {
	for y, row := range grid {
		for x, col := range row {
			if col == "S" {
				return Point{x: x, y: y}
			}
		}
	}
	panic("No starting point found")
}

func isIndexInGrid(grid [][]string, idx Point) bool {
	return idx.x >= 0 && idx.x < len(grid[0]) && idx.y >= 0 && idx.y < len(grid)
}

func nextPoints(grid [][]string, p Point) []Point {
	points := []Point{}
	switch grid[p.y][p.x] {
	case "|":
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x, p.y - 1})
	case "-":
		points = append(points, Point{p.x + 1, p.y})
		points = append(points, Point{p.x - 1, p.y})
	case "L":
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x + 1, p.y})
	case "J":
		points = append(points, Point{p.x, p.y - 1})
		points = append(points, Point{p.x - 1, p.y})
	case "7":
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x - 1, p.y})
	case "F":
		points = append(points, Point{p.x, p.y + 1})
		points = append(points, Point{p.x + 1, p.y})
	case ".":
	case "S":
		up := Point{p.x, p.y - 1}
		down := Point{p.x, p.y + 1}
		left := Point{p.x - 1, p.y}
		right := Point{p.x + 1, p.y}

		var pos string

		if isIndexInGrid(grid, up) {
			pos = grid[up.y][up.x]
			if pos == "|" || pos == "7" || pos == "F" {
				points = append(points, up)
			}
		}

		if isIndexInGrid(grid, down) {
			pos = grid[down.y][down.x]
			if pos == "|" || pos == "L" || pos == "J" {
				points = append(points, down)
			}
		}

		if isIndexInGrid(grid, left) {
			pos = grid[left.y][left.x]
			if pos == "-" || pos == "L" || pos == "F" {
				points = append(points, left)
			}
		}

		if isIndexInGrid(grid, right) {
			pos = grid[right.y][right.x]
			if pos == "-" || pos == "7" || pos == "J" {
				points = append(points, right)
			}
		}

		// down, right, up, left := grid[p.y+1][p.x], grid[p.y][p.x+1], grid[p.y-1][p.x], grid[p.y][p.x-1]
		// if down == "|" || down == "L" || down == "J" {
		// 	points = append(points, Point{p.x, p.y + 1})
		// }
		// if right == "-" || right == "7" || right == "J" {
		// 	points = append(points, Point{p.x + 1, p.y})
		// }
		// if up == "|" || up == "7" || up == "F" {
		// 	points = append(points, Point{p.x, p.y - 1})
		// }
		// if left == "-" || left == "L" || left == "F" {
		// 	points = append(points, Point{p.x - 1, p.y})
		// }
	}
	return points
}
