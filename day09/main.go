package main

import (
	"fmt"
	"math"
	"os"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func readData(lines []string) (r []helper.Coord) {
	r = make([]helper.Coord, 0)
	for _, line := range lines {
		nums, _ := helper.StrCsvToIntArray(line, ",")
		r = append(r, helper.Coord{X: nums[0], Y: nums[1]})
	}
	return r
}

func partOne(lines []string) (r int, err error) {
	coords := readData(lines)
	r = math.MinInt
	for i, coord := range coords {
		others := make([]helper.Coord, 0)
		others = append(others, coords[:i]...)
		others = append(others, coords[i+1:]...)

		for _, c := range others {
			area := coord.Area(c)
			if area > r {
				r = area
			}
		}
	}
	return r, err
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

	//ans, err = partTwo(lines)
	//fmt.Printf("Part two: %d\n", ans)

}
