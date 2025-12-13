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
	Expected        []bool
	State           []bool
	Instructions    []Instruction
	JoltRequirement []int
}

type State struct {
	LightState []bool
	Turns      int
}

type StateTwo struct {
	JoltState []int
	Turns     int
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

func stateAsKeyTwo(i []int) (r string) {
	r = ""
	for _, v := range i {
		r += strconv.Itoa(v)
	}
	return r
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

	lp = strings.Index(line, "{")
	rp = strings.Index(line, "}")
	js := line[lp+1 : rp]
	for _, j := range strings.Split(js, ",") {
		jolt, _ := strconv.Atoi(j)
		r.JoltRequirement = append(r.JoltRequirement, jolt)
	}

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

func (m Machine) InvalidNextState(s []int) bool {
	for i, v := range s {
		if v > m.JoltRequirement[i] {
			return true
		}
	}
	return false
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

func (m Machine) SimulateTwo() int {
	queue := []StateTwo{{JoltState: make([]int, len(m.JoltRequirement)), Turns: 0}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if slices.Equal(m.JoltRequirement, next.JoltState) {
			return next.Turns
		}

		for _, instruction := range m.Instructions {
			stateTurn := slices.Clone(next.JoltState)
			for _, pos := range instruction {
				stateTurn[pos] += 1
			}
			if !m.InvalidNextState(stateTurn) {
				stateKey := stateAsKeyTwo(stateTurn)
				if _, ok := visited[stateKey]; !ok {
					visited[stateKey] = true
					queue = append(queue, StateTwo{JoltState: slices.Clone(stateTurn), Turns: next.Turns + 1})
				}
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

func partTwo(lines []string) (r int, err error) {
	machines := readData(lines)
	for _, machine := range machines {
		r += machine.SimulateTwo()
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
