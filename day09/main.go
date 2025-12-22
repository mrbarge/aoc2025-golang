package main

import (
	"fmt"
	"math"
	"os"
	"slices"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type XRange struct {
	X1 int
	X2 int
}

func readData(lines []string) (r []helper.Coord) {
	r = make([]helper.Coord, 0)
	for _, line := range lines {
		nums, _ := helper.StrCsvToIntArray(line, ",")
		r = append(r, helper.Coord{X: nums[0], Y: nums[1]})
	}
	return r
}

func readDataTwo(lines []string) (red []helper.Coord, yranges map[int]XRange) {
	reds := readData(lines)

	nxt := slices.Clone(reds)
	nxt = append(nxt, reds[0])

	yranges = make(map[int]XRange)
	for i, v := range nxt[:len(nxt)-1] {
		nv := nxt[i+1]
		var minx, miny, maxx, maxy int = nv.X, nv.Y, v.X, v.Y
		minx = min(nv.X, v.X)
		miny = min(nv.Y, v.Y)
		maxx = max(nv.X, v.X)
		maxy = max(nv.Y, v.Y)

		for iy := miny; iy < maxy+1; iy++ {
			if yr, ok := yranges[iy]; ok {
				yranges[iy] = XRange{min(minx, yr.X1), max(maxx, yr.X2)}
			} else {
				yranges[iy] = XRange{minx, maxx}
			}
		}
	}

	return reds, yranges
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

func valid(c1 helper.Coord, c2 helper.Coord, yranges map[int]XRange) bool {
	minx := min(c1.X, c2.X)
	miny := min(c1.Y, c2.Y)
	maxx := max(c1.X, c2.X)
	maxy := max(c1.Y, c2.Y)

	for y := miny; y < maxy+1; y++ {
		if _, ok := yranges[y]; ok {
			if minx < yranges[y].X1 || minx > yranges[y].X2 ||
				maxx < yranges[y].X1 || maxx > yranges[y].X2 {
				return false
			}
		} else {
			fmt.Printf("What happened here?")
			return false
		}
	}
	return true
}

func partTwo(lines []string) (r int, err error) {
	coords, yranges := readDataTwo(lines)
	r = math.MinInt
	for i, coord := range coords {
		others := make([]helper.Coord, 0)
		others = append(others, coords[:i]...)
		others = append(others, coords[i+1:]...)

		for _, c := range others {
			area := coord.Area(c)
			if area > r && valid(coord, c, yranges) {
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

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %d\n", ans)

}
