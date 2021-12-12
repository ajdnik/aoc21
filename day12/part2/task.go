package main

import (
	"log"

	"github.com/ajdnik/aoc21/day12"
	"github.com/ajdnik/aoc21/utils"
)

func HasLowerDuplicate(path []string) bool {
	lowers := map[string]bool{}
	for _, node := range path {
		if day12.IsUpper(node) {
			continue
		}
		if _, ok := lowers[node]; !ok {
			lowers[node] = true
			continue
		}
		return true
	}
	return false
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	lines := []string{}
	for scanner.Scan() {
		data := scanner.Text()
		lines = append(lines, data)
	}
	graph := day12.ToGraph(lines)
	paths := day12.TraverseGraph(graph, []string{"start"}, func(path []string, node string) bool {
		return day12.IsIncluded(path, node) && !day12.IsUpper(node) && HasLowerDuplicate(path)
	})
	log.Printf("numPaths=%d\n", len(paths))
}
