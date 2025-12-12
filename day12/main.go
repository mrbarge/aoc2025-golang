package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Shape [][]bool
type Region struct {
	X          int
	Y          int
	Quantities []int
}

func NewShape() Shape {
	return [][]bool{
		make([]bool, 3),
		make([]bool, 3),
		make([]bool, 3),
	}
}

func NewRegion(x int, y int) Region {
	return Region{X: x, Y: y, Quantities: make([]int, 0)}
}

func readData(lines []string) (r []Shape, pr []Region) {
	r = make([]Shape, 0)
	pr = make([]Region, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "x") {
			x, _ := strconv.Atoi(strings.Split(line, "x")[0])
			y, _ := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], "x")[1])
			presents := strings.Split(line, " ")[1:]
			present := NewRegion(x, y)
			for _, pres := range presents {
				p, _ := strconv.Atoi(pres)
				present.Quantities = append(present.Quantities, p)
			}
			pr = append(pr, present)
		} else if strings.Contains(line, ":") {
			// Shape
			s := NewShape()
			i++
			y := 0
			line = lines[i]
			for line != "" {
				for x, c := range line {
					if c == '#' {
						s[y][x] = true
					}
				}
				y++
				i++
				line = lines[i]
			}
			r = append(r, s)
		}
	}
	return r, pr
}

func partOne(lines []string) (r int, err error) {
	shapes, regions := readData(lines)
	fmt.Printf("shapes: %v\n", shapes)
	fmt.Printf("regions: %v\n", regions)

	for _, region := range regions {
		area := region.X * region.Y
		sum := 0
		for _, q := range region.Quantities {
			sum += q * 9
		}
		if sum <= area {
			r++
		}
	}
	return r, err
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, false)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines)
	fmt.Printf("Part one: %d\n", ans)

	//ans, err = partTwo(lines)
	//fmt.Printf("Part two: %d\n", ans)

}
