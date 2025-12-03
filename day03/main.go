package main

import (
	"fmt"
	"math"
	"os"
	"sort"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func joltage(line []int) (r int) {
	largest := make([]int, 0)
	for i := 0; i < len(line)-1; i++ {
		for j := i + 1; j < len(line); j++ {
			largest = append(largest, line[i]*10+line[j])
		}
	}
	sort.Ints(largest)
	return largest[len(largest)-1]
}

func joltageTwo(line []int) (r int) {
	largest := make([]int, 0)
	for i := 0; i < len(line)-12; i++ {
		currPos := i
		jolt := line[currPos] * int(math.Pow10(11))
		for digitsRemaining := 11; digitsRemaining > 0; digitsRemaining-- {
			// scan forward - look for the closest, largest digit between i and len-digitsRemaining
			largestNext := line[currPos+1]
			largestPos := currPos + 1
			for j := currPos + 1; j < len(line)-digitsRemaining; j++ {
				if line[j] > largestNext {
					largestNext = line[j]
					largestPos = j
				}
			}
			jolt += largestNext * int(math.Pow10(digitsRemaining-1))
			currPos = largestPos
		}
		largest = append(largest, jolt)
	}
	sort.Ints(largest)
	return largest[len(largest)-1]
}

func partOne(lines [][]int) (r int, err error) {
	for _, line := range lines {
		r += joltage(line)
	}
	return r, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLinesAsIntArray(fh)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %d\n", ans)

	//ans, err = partTwo(lines[0])
	//fmt.Printf("Part two: %d\n", ans)

}
