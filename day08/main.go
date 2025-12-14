package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type JunctionBox struct {
	Id          int
	C           helper.Coord3D
	Connections []*JunctionBox
	Circuits    []*Circuit
}

type Connection struct {
	Distance int
	Boxes    [2]*JunctionBox
}

type Circuit struct {
	Id    int
	Boxes []*JunctionBox
}

func (c *Circuit) HasBox(jb *JunctionBox) bool {
	if c.Boxes == nil || len(c.Boxes) == 0 {
		return false
	}
	return slices.Contains(c.Boxes, jb)
}

func (c *Circuit) AddBox(jb *JunctionBox) {
	for _, b := range c.Boxes {
		if b.Id == jb.Id {
			return
		}
	}
	c.Boxes = append(c.Boxes, jb)
}

func readData(lines []string) (r []*JunctionBox) {
	r = make([]*JunctionBox, 0)
	for i, line := range lines {
		elems := strings.Split(line, ",")
		ix, _ := strconv.Atoi(elems[0])
		iy, _ := strconv.Atoi(elems[1])
		iz, _ := strconv.Atoi(elems[2])
		jb := JunctionBox{Id: i, C: helper.Coord3D{X: ix, Y: iy, Z: iz}, Connections: make([]*JunctionBox, 0), Circuits: make([]*Circuit, 0)}
		r = append(r, &jb)
	}
	return r
}

func makeDistances(boxes []*JunctionBox) []Connection {
	distances := make([]Connection, 0)
	for i, b1 := range boxes {
		for _, b2 := range boxes[i+1:] {
			distance := b1.C.DistanceTo(b2.C)
			distances = append(distances, Connection{
				Distance: distance,
				Boxes:    [2]*JunctionBox{b1, b2},
			})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})
	return distances
}

func makeCircuits(boxes []*JunctionBox) map[int]*Circuit {
	r := make(map[int]*Circuit, 0)
	for i, b := range boxes {
		r[i] = &Circuit{Id: i, Boxes: []*JunctionBox{b}}
	}
	return r
}

func getBoxCircuit(circuits map[int]*Circuit, box *JunctionBox) *Circuit {
	for _, c := range circuits {
		if c.HasBox(box) {
			return c
		}
	}
	return nil
}

func partOne(lines []string, connections int) (r int, err error) {
	boxes := readData(lines)
	distances := makeDistances(boxes)
	circuits := makeCircuits(boxes)

	p2Answer := 0
	for connections > 0 && len(circuits) > 1 {

		closest := distances[0]
		distances = distances[1:]
		p2Answer = closest.Boxes[0].C.X * closest.Boxes[1].C.X

		b1Circuit := getBoxCircuit(circuits, closest.Boxes[0])
		b2Circuit := getBoxCircuit(circuits, closest.Boxes[1])

		if b1Circuit != b2Circuit {
			// Merge and remove the second circuit
			for _, b := range b2Circuit.Boxes {
				b1Circuit.AddBox(b)
			}
			delete(circuits, b2Circuit.Id)
		}
		connections--
	}

	circuitLengths := make([]int, 0)
	for _, circuit := range circuits {
		circuitLengths = append(circuitLengths, len(circuit.Boxes))
	}

	// crude part 1 or part 2 test
	if len(circuitLengths) > 1 {
		sort.Ints(circuitLengths)
		return circuitLengths[len(circuitLengths)-1] * circuitLengths[len(circuitLengths)-2] * circuitLengths[len(circuitLengths)-3], nil
	} else {
		return p2Answer, nil
	}
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines, 1000)
	fmt.Printf("Part one: %d\n", ans)

	ans, err = partOne(lines, math.MaxInt32)
	fmt.Printf("Part two: %d\n", ans)

}
