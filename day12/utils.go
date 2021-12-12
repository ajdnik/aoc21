package day12

import "strings"

func ToGraph(data []string) map[string][]string {
	graph := map[string][]string{}
	for _, line := range data {
		nodes := strings.Split(line, "-")
		if _, ok := graph[nodes[0]]; !ok {
			graph[nodes[0]] = []string{}
		}
		graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		if _, ok := graph[nodes[1]]; !ok {
			graph[nodes[1]] = []string{}
		}
		graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
	}
	return graph
}

func IsIncluded(data []string, val string) bool {
	for _, itm := range data {
		if itm == val {
			return true
		}
	}
	return false
}

func IsUpper(itm string) bool {
	return strings.ToUpper(itm) == itm
}

func TraverseGraph(graph map[string][]string, path []string, skip func(path []string, node string) bool) [][]string {
	nodes := graph[path[len(path)-1]]
	paths := [][]string{}
	for _, node := range nodes {
		if node == "start" {
			continue
		}
		if node == "end" {
			paths = append(paths, append(path, node))
			continue
		}
		if skip(path, node) {
			continue
		}
		res := TraverseGraph(graph, append(path, node), skip)
		paths = append(paths, res...)
	}
	return paths
}
