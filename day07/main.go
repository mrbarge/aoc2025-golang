package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

func simulateBeam(beam helper.Coord, lines []string, seenPos map[helper.Coord]int) (r int) {

	if _, ok := seenPos[beam]; ok {
		return seenPos[beam]
	}

	if beam.Y == len(lines)-1 {
		return 1
	}

	nextBeam := helper.Coord{X: beam.X, Y: beam.Y + 1}
	if lines[beam.Y+1][beam.X] == '^' {
		left := helper.Coord{X: beam.X - 1, Y: beam.Y + 1}
		right := helper.Coord{X: beam.X + 1, Y: beam.Y + 1}
		r += simulateBeam(left, lines, seenPos)
		r += simulateBeam(right, lines, seenPos)
	} else {
		r += simulateBeam(nextBeam, lines, seenPos)
	}
	seenPos[nextBeam] = r
	return r
}

func simulate(beams []helper.Coord, lines []string) (splits int, next []helper.Coord) {
	seen := make(map[helper.Coord]bool)
	next = make([]helper.Coord, 0)

	// Are we at the bottom?
	if beams[0].Y == len(lines)-1 {
		return -1, beams
	}

	for _, beam := range beams {
		nextBeam := helper.Coord{X: beam.X, Y: beam.Y + 1}
		if _, ok := seen[nextBeam]; ok {
			continue
		}
		if lines[nextBeam.Y][nextBeam.X] == '^' {
			// split
			splits++
			leftBeam := helper.Coord{X: beam.X - 1, Y: beam.Y + 1}
			rightBeam := helper.Coord{X: beam.X + 1, Y: beam.Y + 1}
			if _, ok := seen[leftBeam]; !ok {
				next = append(next, leftBeam)
				seen[leftBeam] = true
			}
			if _, ok := seen[rightBeam]; !ok {
				next = append(next, rightBeam)
				seen[rightBeam] = true
			}
		} else {
			next = append(next, nextBeam)
		}
		seen[nextBeam] = true
	}
	return splits, next
}

func partOne(lines []string) (r int, err error) {
	startPos := helper.Coord{X: strings.Index(lines[0], "S")}
	beams := []helper.Coord{startPos}
	done := false
	for !done {
		turnsplits, nextbeams := simulate(beams, lines)
		if turnsplits == -1 {
			done = true
			break
		}
		beams = nextbeams
		r += turnsplits
	}
	return r, err
}

func partTwo(lines []string) (r int, err error) {
	startPos := helper.Coord{X: strings.Index(lines[0], "S")}
	beams := []helper.Coord{startPos}
	seenPos := make(map[helper.Coord]int)
	r = simulateBeam(beams[0], lines, seenPos)
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

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %d\n", ans)

}
