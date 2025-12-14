package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"

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

func buildMatrix(machine Machine) ([][]int, []int) {
	numCounters := len(machine.JoltRequirement)
	numButtons := len(machine.Instructions)
	matrix := make([][]int, numCounters)
	for i := range matrix {
		matrix[i] = make([]int, numButtons)
	}
	for buttonIdx, button := range machine.Instructions {
		for _, counterIdx := range button {
			matrix[counterIdx][buttonIdx] = 1
		}
	}
	return matrix, machine.JoltRequirement
}

func solveWithGonum(matrix [][]int, targets []int) []int {
	numEquations := len(matrix)
	numVariables := len(matrix[0])

	// Convert your int matrix to float64 for gonum
	data := make([]float64, numEquations*numVariables)
	for i := 0; i < numEquations; i++ {
		for j := 0; j < numVariables; j++ {
			data[i*numVariables+j] = float64(matrix[i][j])
		}
	}

	// Create gonum matrix
	A := mat.NewDense(numEquations, numVariables, data)

	// Create target vector
	fl := make([]float64, len(targets))
	for i, v := range targets {
		fl[i] = float64(v)
	}
	b := mat.NewVecDense(numEquations, fl)

	// Solve the system A*x = b
	var x mat.VecDense
	err := x.SolveVec(A, b)
	if err != nil {
		// Handle error - might not have a solution
	}

	// Extract solution and round to integers
	solution := make([]int, numVariables)
	for i := 0; i < numVariables; i++ {
		solution[i] = int(math.Round(x.AtVec(i)))
	}

	return solution
}

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
		matrix, _ := buildMatrix(machine)
		rs := solveWithGonum(matrix, machine.JoltRequirement)
		r += len(rs)
		fmt.Printf("Solved machine %v\n", rs)
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

	ans, err = partTwo(lines)
	fmt.Printf("Part two: %d\n", ans)

}
