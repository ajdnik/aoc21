package main

import (
	"log"

	"github.com/ajdnik/aoc21/day12"
	"github.com/ajdnik/aoc21/utils"
)

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
		return day12.IsIncluded(path, node) && !day12.IsUpper(node)
	})
	log.Printf("numPaths=%d\n", len(paths))
}
