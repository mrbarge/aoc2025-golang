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
	C           helper.Coord3D
	Connections []*JunctionBox
	Circuits    []*Circuit
}
type Circuit struct {
	Id    int
	Boxes []*JunctionBox
}

type ClosestPair struct {
	Boxes    []*JunctionBox
	Distance float64
}

func (c Circuit) HasBox(jb *JunctionBox) bool {
	if c.Boxes == nil || len(c.Boxes) == 0 {
		return false
	}
	return slices.Contains(c.Boxes, jb)
}

func readData(lines []string) (r []*JunctionBox) {
	r = make([]*JunctionBox, 0)
	for _, line := range lines {
		elems := strings.Split(line, ",")
		ix, _ := strconv.Atoi(elems[0])
		iy, _ := strconv.Atoi(elems[1])
		iz, _ := strconv.Atoi(elems[2])
		jb := JunctionBox{C: helper.Coord3D{X: ix, Y: iy, Z: iz}, Connections: make([]*JunctionBox, 0), Circuits: make([]*Circuit, 0)}
		r = append(r, &jb)
	}
	return r
}

func (jb *JunctionBox) FindClosest(boxes []*JunctionBox) (*JunctionBox, float64) {
	distance := math.MaxFloat64
	var closest *JunctionBox
	for _, b := range boxes {
		// Ignore boxes we're already connected to
		if jb.IsConnectedTo(b) {
			continue
		}
		d := jb.C.DistanceTo(b.C)
		if d < distance {
			distance = d
			closest = b
		}
	}
	return closest, distance
}

func (jb *JunctionBox) IsConnectedTo(box *JunctionBox) bool {
	if jb.Connections == nil || len(jb.Connections) == 0 {
		return false
	}
	return slices.Contains(jb.Connections, box)
}

func (jb *JunctionBox) ConnectTo(box *JunctionBox) {
	jb.Connections = append(jb.Connections, box)
	box.Connections = append(box.Connections, jb)
}

func partOne(lines []string, connections int) (r int, err error) {
	boxes := readData(lines)
	circuits := make([]*Circuit, 0)
	circuitId := 0
	connectionsMade := 0
	for connectionsMade < connections {
		closestPair := ClosestPair{Boxes: make([]*JunctionBox, 0), Distance: math.MaxFloat64}
		for j, box := range boxes {
			otherboxes := make([]*JunctionBox, 0)
			otherboxes = append(otherboxes, boxes[:j]...)
			otherboxes = append(otherboxes, boxes[j+1:]...)
			closest, dist := box.FindClosest(otherboxes)
			if dist < closestPair.Distance {
				closestPair.Boxes = []*JunctionBox{box, closest}
				closestPair.Distance = dist
			}
		}
		//fmt.Printf("Connecting %v and %v\n", closestPair.Boxes[0].C, closestPair.Boxes[1].C)
		// Connect the two
		closestPair.Boxes[0].ConnectTo(closestPair.Boxes[1])
		// Is there a circuit they are part of?
		madeConnection := false
		alreadyConnected := false
		for _, circuit := range circuits {
			if circuit.HasBox(closestPair.Boxes[0]) {
				if !circuit.HasBox(closestPair.Boxes[1]) {
					fmt.Printf("Adding to existing circuit %v.. (pre len %v)\n", circuit.Id, len(circuit.Boxes))
					circuit.Boxes = append(circuit.Boxes, closestPair.Boxes[1])
					madeConnection = true
				} else {
					alreadyConnected = true
				}
			} else if circuit.HasBox(closestPair.Boxes[1]) {
				if !circuit.HasBox(closestPair.Boxes[0]) {
					circuit.Boxes = append(circuit.Boxes, closestPair.Boxes[0])
					madeConnection = true
				} else {
					alreadyConnected = true
				}
			}
		}

		if !madeConnection {
			circuits = append(circuits, &Circuit{Id: circuitId, Boxes: []*JunctionBox{closestPair.Boxes[0], closestPair.Boxes[1]}})
			fmt.Printf("Creating new circuit %v..\n", circuitId)
			circuitId++
		}
		if !alreadyConnected {
			connectionsMade++
		}
		printLengths(circuits)
	}

	circuitSizes := make([]int, 0)
	for i, circuit := range circuits {
		fmt.Printf("Circuit %v is %v\n", i, len(circuit.Boxes))
		circuitSizes = append(circuitSizes, len(circuit.Boxes))
	}
	sort.Ints(circuitSizes)
	//return circuitSizes[len(circuitSizes)-1] * circuitSizes[len(circuitSizes)-2] * circuitSizes[len(circuitSizes)-3], nil
	return 0, nil
}

func printLengths(l []*Circuit) {
	for _, circuit := range l {
		fmt.Printf("%v ", len(circuit.Boxes))
	}
	fmt.Println()
}

func main() {
	fh, _ := os.Open("test.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	ans, err := partOne(lines, 10)
	fmt.Printf("Part one: %d\n", ans)

	//ans, err = partTwo(lines)
	//fmt.Printf("Part two: %d\n", ans)

}
