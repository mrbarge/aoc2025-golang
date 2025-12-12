package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Machine struct {
	Expected []bool
	State    []bool
}

type Instruction []int

func (m Machine) Done() bool {
	return slices.Equal(m.State, m.Expected)
}

func (m Machine) Set(pos int) bool {
	m.State[pos] = !m.State[pos]
	return m.Expected[pos] == m.State[pos]
}

func runInstruction(m *Machine, instruction []int) {
	for _, pos := range instruction {
		m.Set(pos)
	}
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
	return r
}

func readInstructions(line string) (r []Instruction) {
	r = make([]Instruction, 0)
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
		r = append(r, i)
	}
	return r
}

func readData(lines []string) (m []Machine, i [][]Instruction) {
	i = make([][]Instruction, 0)
	m = make([]Machine, 0)
	for _, line := range lines {
		inst := readInstructions(line)
		machine := readMachine(line)

		m = append(m, machine)
		i = append(i, inst)
	}
	return m, i
}

func simulate(state []bool, expected []bool, turns int, i []Instruction) int {
	// Are we there yet
	if slices.Equal(state, expected) {
		return turns
	}

	minturns := math.MaxInt
	for _, inst := range i {
		closer := false
		newstate := slices.Clone(state)
		for _, pos := range inst {
			newstate[pos] = !newstate[pos]
			closer = closer || newstate[pos] == expected[pos]
		}
		// did this bring us any closer
		if closer {
			nt := simulate(newstate, expected, turns+1, i)
			if nt < minturns {
				minturns = nt
			}
		}
	}
	return minturns
}

func partOne(lines []string) (r int, err error) {
	machines, instructions := readData(lines)
	for i, machine := range machines {
		mi := instructions[i]
		turns := simulate(machine.State, machine.Expected, 0, mi)
		fmt.Printf("turns: %d\n", turns)
	}
	return r, err
}

func main() {
	fh, _ := os.Open("test.txt")
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
