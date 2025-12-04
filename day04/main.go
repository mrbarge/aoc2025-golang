package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func countRolls(lines []string) (r int, newRolls []string) {
	newRolls = make([]string, 0)
	for y, row := range lines {
		newRow := ""
		for x, _ := range row {
			c := helper.Coord{X: x, Y: y}
			rolls := 0
			neighbours := c.GetNeighboursPos(true)
			for _, neighbour := range neighbours {
				if neighbour.X >= len(row) || neighbour.Y >= len(lines) {
					continue
				}
				if lines[neighbour.Y][neighbour.X] == '@' {
					rolls++
				}
			}
			if lines[y][x] == '@' && rolls < 4 {
				r++
				newRow += "."
			} else {
				newRow += string(lines[y][x])
			}
		}
		newRolls = append(newRolls, newRow)
	}
	return r, newRolls
}

func partOne(lines []string) (r int, err error) {
	answer, _ := countRolls(lines)
	return answer, nil
}

func partTwo(lines []string) (r int, err error) {
	result, newGrid := countRolls(lines)
	for result != 0 {
		r += result
		result, newGrid = countRolls(newGrid)
	}
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %d\n", ans)

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %d\n", ans)

}
