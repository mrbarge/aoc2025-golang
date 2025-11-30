package main

import (
	"fmt"
	"os"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func problem(lines []string, partTwo bool) (r float64, err error) {
	return 0, nil
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := problem(lines, false)
	fmt.Printf("Part one: %f\n", ans)

	ans, err = problem(lines, true)
	fmt.Printf("Part two: %f\n", ans)

}
