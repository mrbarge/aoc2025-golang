package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Machine struct {
	Expected     []bool
	State        []bool
	Instructions []Instruction
}

type State struct {
	LightState []bool
	Turns      int
}

func (m Machine) ApplyInstruction(instruction Instruction) []bool {
	r := slices.Clone(m.State)
	for _, pos := range instruction {
		r[pos] = !r[pos]
	}
	return r
}

type Instruction []int

func stateAsKey(i []bool) (r string) {
	r = ""
	for _, v := range i {
		if v {
			r += "#"
		} else {
			r += "."
		}
	}
	return r
}

func (m Machine) Done() bool {
	return slices.Equal(m.State, m.Expected)
}

func (m Machine) Set(pos int) bool {
	m.State[pos] = !m.State[pos]
	return m.Expected[pos] == m.State[pos]
}

func readMachine(line string) (r Machine) {
	lp := strings.Index(line, "[")
	rp := strings.Index(line, "]")
	for i := lp + 1; i < rp; i++ {
		if line[i] == '.' {
			r.Expected = append(r.Expected, false)
		} else if line[i] == '#' {
			r.Expected = append(r.Expected, true)
		}
		r.State = append(r.State, false)
	}

	elems := strings.Split(line, " ")
	for _, elem := range elems {
		if strings.Contains(elem, "[") || strings.Contains(elem, "{") {
			continue
		}
		i := make(Instruction, 0)
		nums := strings.Split(elem[1:len(elem)-1], ",")
		for _, num := range nums {
			inum, _ := strconv.Atoi(num)
			i = append(i, inum)

		}
		r.Instructions = append(r.Instructions, i)
	}

	return r
}

func readInstructions(line string) (r []Instruction) {
	return r
}

func readData(lines []string) (m []Machine) {
	m = make([]Machine, 0)
	for _, line := range lines {
		machine := readMachine(line)
		m = append(m, machine)
	}
	return m
}

func (m Machine) Simulate() int {
	queue := []State{{LightState: slices.Clone(m.State), Turns: 0}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if slices.Equal(next.LightState, m.Expected) {
			return next.Turns
		}
		for _, instruction := range m.Instructions {
			stateTurn := slices.Clone(next.LightState)
			for _, pos := range instruction {
				stateTurn[pos] = !stateTurn[pos]
			}
			stateKey := stateAsKey(stateTurn)
			if _, ok := visited[stateKey]; !ok {
				visited[stateKey] = true
				queue = append(queue, State{LightState: slices.Clone(stateTurn), Turns: next.Turns + 1})
			}
		}
	}
	return -1
}

func partOne(lines []string) (r int, err error) {
	machines := readData(lines)
	for _, machine := range machines {
		r += machine.Simulate()
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
