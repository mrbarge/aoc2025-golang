package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Graph struct {
	Nodes map[string][]string
}

func (g *Graph) AddNode(from string, to string) {
	g.Nodes[from] = append(g.Nodes[from], to)
}

func (g *Graph) GetNodes() []string {
	return slices.Collect(maps.Keys(g.Nodes))
}

func (g *Graph) DFS(current string, target string, path []string, visited map[string]bool, allPaths *[][]string) {
	visited[current] = true

	if current == target {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
	} else {
		for _, neighbor := range g.Nodes[current] {
			if !visited[neighbor] {
				path = append(path, neighbor)
				g.DFS(neighbor, target, path, visited, allPaths)
				path = path[:len(path)-1]
			}
		}
	}

	visited[current] = false
}

func readData(lines []string) (g *Graph) {
	g = &Graph{Nodes: make(map[string][]string)}

	for _, line := range lines {
		id := strings.Split(line, ":")[0]
		next := strings.Split(line, " ")[1:]
		for _, n := range next {
			g.AddNode(id, n)
		}
	}
	return g
}

func (g *Graph) simulate(from string, to string) [][]string {
	var allPaths [][]string
	currentPath := []string{from}
	visited := make(map[string]bool)

	g.DFS(from, to, currentPath, visited, &allPaths)
	return allPaths
}

func partOne(lines []string) (r int, err error) {
	g := readData(lines)

	return len(g.simulate("you", "out")), err
}

func partTwo(lines []string) (r int, err error) {
	g := readData(lines)

	ap := g.simulate("svr", "out")
	return len(ap), err
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
