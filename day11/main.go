package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrbarge/aoc2025-golang/helper"
)

type Graph struct {
	nodes    map[int][]int
	nameToID map[string]int
	idToName map[int]string
	nextID   int
}

func NewGraph() *Graph {
	return &Graph{
		nodes:    make(map[int][]int),
		nameToID: make(map[string]int),
		idToName: make(map[int]string),
		nextID:   0,
	}
}

func (g *Graph) GetOrCreateNodeID(name string) int {
	if id, exists := g.nameToID[name]; exists {
		return id
	}
	id := g.nextID
	g.nameToID[name] = id
	g.idToName[id] = name
	g.nextID++
	return id
}

func (g *Graph) NodeName(id int) string {
	return g.idToName[id]
}

func (g *Graph) GetNodeID(name string) int {
	if id, exists := g.nameToID[name]; exists {
		return id
	}
	return -1
}

func (g *Graph) AddEdge(src, dest int) {
	g.nodes[src] = append(g.nodes[src], dest)
}

func (g *Graph) AddEdgeByName(srcName, destName string) {
	srcID := g.GetOrCreateNodeID(srcName)
	destID := g.GetOrCreateNodeID(destName)
	g.AddEdge(srcID, destID)
}

func (g *Graph) ReadData(lines []string) error {
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		src := strings.TrimSpace(parts[0])
		dests := strings.Fields(parts[1])
		for _, dest := range dests {
			dest = strings.TrimSpace(dest)
			if dest != "" {
				g.AddEdgeByName(src, dest)
			}
		}
		if len(dests) == 0 {
			g.GetOrCreateNodeID(src)
		}
	}
	return nil
}

func (g *Graph) BuildReverseGraph() *Graph {
	reverse := &Graph{
		nodes:    make(map[int][]int),
		nameToID: g.nameToID,
		idToName: g.idToName,
		nextID:   g.nextID,
	}

	for src, neighbours := range g.nodes {
		for _, dest := range neighbours {
			reverse.AddEdge(dest, src)
		}
	}
	return reverse
}

type PathCache struct {
	pathCache  map[string][][]int
	reachCache map[string]bool
}

func NewPathCache() *PathCache {
	return &PathCache{
		pathCache:  make(map[string][][]int),
		reachCache: make(map[string]bool),
	}
}

func (pc *PathCache) getKey(start, end int, forbidden []int) string {
	key := fmt.Sprintf("%d-%d", start, end)
	if len(forbidden) > 0 {
		key += fmt.Sprintf(":forbid%v", forbidden)
	}
	return key
}

func (pc *PathCache) GetPaths(start, end int, forbidden []int) ([][]int, bool) {
	paths, exists := pc.pathCache[pc.getKey(start, end, forbidden)]
	return paths, exists
}

func (pc *PathCache) SetPaths(start, end int, forbidden []int, paths [][]int) {
	pc.pathCache[pc.getKey(start, end, forbidden)] = paths
}

func (g *Graph) GetReachableNodes(start int) map[int]bool {
	reachable := make(map[int]bool)
	queue := []int{start}
	reachable[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, neighs := range g.nodes[vertex] {
			if !reachable[neighs] {
				reachable[neighs] = true
				queue = append(queue, neighs)
			}
		}
	}

	return reachable
}

func (g *Graph) PruneGraphForSegment(start, end int) map[int]bool {
	forwardReachable := g.GetReachableNodes(start)
	reverseGraph := g.BuildReverseGraph()
	backwardReachable := reverseGraph.GetReachableNodes(end)
	relevant := make(map[int]bool)
	for node := range forwardReachable {
		if backwardReachable[node] {
			relevant[node] = true
		}
	}
	if len(relevant) < 20 {
		var nodeNames []string
		for node := range relevant {
			nodeNames = append(nodeNames, g.NodeName(node))
		}
	}
	return relevant
}

func (g *Graph) FindAllPathsWithPruning(start, target int, forbidden map[int]bool, cache *PathCache) [][]int {
	var forbiddenSlice []int
	for node := range forbidden {
		forbiddenSlice = append(forbiddenSlice, node)
	}
	if paths, exists := cache.GetPaths(start, target, forbiddenSlice); exists {
		return paths
	}

	relevantNodes := g.PruneGraphForSegment(start, target)
	var allPaths [][]int
	currentPath := []int{start}
	visited := make(map[int]bool)
	g.dfsWithPruning(start, target, currentPath, visited, forbidden, relevantNodes, &allPaths)
	cache.SetPaths(start, target, forbiddenSlice, allPaths)
	return allPaths
}

func (g *Graph) dfsWithPruning(current, target int, path []int, visited, forbidden, relevant map[int]bool, allPaths *[][]int) {
	visited[current] = true
	if current == target {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*allPaths = append(*allPaths, pathCopy)
	} else {
		for _, neighbour := range g.nodes[current] {
			if forbidden[neighbour] || !relevant[neighbour] || visited[neighbour] {
				continue
			}

			path = append(path, neighbour)
			g.dfsWithPruning(neighbour, target, path, visited, forbidden, relevant, allPaths)
			path = path[:len(path)-1]
		}
	}
	visited[current] = false
}

func (g *Graph) GetTotalPaths(start, end int, requiredNodes []int) int {
	cache := NewPathCache()

	sequence := append([]int{start}, requiredNodes...)
	sequence = append(sequence, end)
	for i := 0; i < len(sequence)-1; i++ {
		forwardReach := g.GetReachableNodes(sequence[i])
		if !forwardReach[sequence[i+1]] {
			return -1
		}
	}
	var segmentPaths [][][]int
	for i := 0; i < len(sequence)-1; i++ {
		from := sequence[i]
		to := sequence[i+1]
		forbidden := make(map[int]bool)
		for j := i + 2; j < len(sequence); j++ {
			forbidden[sequence[j]] = true
		}
		if len(forbidden) > 0 {
			var forbiddenNames []string
			for node := range forbidden {
				forbiddenNames = append(forbiddenNames, g.NodeName(node))
			}
		}
		paths := g.FindAllPathsWithPruning(from, to, forbidden, cache)
		if len(paths) == 0 {
			return -1
		}
		segmentPaths = append(segmentPaths, paths)
	}
	if len(segmentPaths) == 0 {
		return -1
	}
	r := len(segmentPaths[0])
	for _, sp := range segmentPaths[1:] {
		r *= len(sp)
	}
	return r
}

func (g *Graph) combinePaths(segmentPaths [][][]int) [][]int {
	if len(segmentPaths) == 0 {
		return nil
	}

	result := segmentPaths[0]

	for i := 1; i < len(segmentPaths); i++ {
		var newResult [][]int
		for _, existingPath := range result {
			for _, nextSegment := range segmentPaths[i] {
				combined := make([]int, len(existingPath))
				copy(combined, existingPath)
				combined = append(combined, nextSegment[1:]...)
				newResult = append(newResult, combined)
			}
		}
		result = newResult
	}

	return result
}

func (g *Graph) ConvertPathToNames(path []int) []string {
	names := make([]string, len(path))
	for i, id := range path {
		names[i] = g.NodeName(id)
	}
	return names
}

func partOne(lines []string) (r int, err error) {
	g := NewGraph()
	g.ReadData(lines)
	return g.GetTotalPaths(g.GetNodeID("you"), g.GetNodeID("out"), []int{}), nil
}

func partTwo(lines []string) (r int, err error) {
	g := NewGraph()
	g.ReadData(lines)
	return g.GetTotalPaths(g.GetNodeID("svr"), g.GetNodeID("out"), []int{g.GetNodeID("fft"), g.GetNodeID("dac")}), nil
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
