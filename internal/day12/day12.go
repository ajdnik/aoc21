package day12

import (
	"slices"
	"strings"
)

func toGraph(data []string) map[string][]string {
	graph := map[string][]string{}
	for _, line := range data {
		nodes := strings.Split(line, "-")
		graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
	}
	return graph
}

func isUpper(itm string) bool {
	return strings.ToUpper(itm) == itm
}

func traverseGraph(graph map[string][]string, path []string, skip func(path []string, node string) bool) [][]string {
	nodes := graph[path[len(path)-1]]
	var paths [][]string
	for _, node := range nodes {
		if node == "start" {
			continue
		}
		if node == "end" {
			paths = append(paths, append(append([]string{}, path...), node))
			continue
		}
		if skip(path, node) {
			continue
		}
		res := traverseGraph(graph, append(append([]string{}, path...), node), skip)
		paths = append(paths, res...)
	}
	return paths
}

func hasLowerDuplicate(path []string) bool {
	lowers := map[string]bool{}
	for _, node := range path {
		if isUpper(node) {
			continue
		}
		if lowers[node] {
			return true
		}
		lowers[node] = true
	}
	return false
}

func Part1(lines []string) int {
	graph := toGraph(lines)
	paths := traverseGraph(graph, []string{"start"}, func(path []string, node string) bool {
		return slices.Contains(path, node) && !isUpper(node)
	})
	return len(paths)
}

func Part2(lines []string) int {
	graph := toGraph(lines)
	paths := traverseGraph(graph, []string{"start"}, func(path []string, node string) bool {
		return slices.Contains(path, node) && !isUpper(node) && hasLowerDuplicate(path)
	})
	return len(paths)
}
