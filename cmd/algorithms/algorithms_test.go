package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := Graph{Nodes: map[string]GraphNode{
		"A": GraphNode{name: "A", edges: map[string]int{"B": 1}},
		"B": GraphNode{name: "B", edges: map[string]int{"A": 1, "C": 1, "D": 2}},
		"C": GraphNode{name: "C", edges: map[string]int{"B": 1, "D": 2, "E": 2}},
		"D": GraphNode{name: "D", edges: map[string]int{"B": 2, "C": 2, "E": 3}},
		"E": GraphNode{name: "E", edges: map[string]int{"C": 2, "D": 3}},
	}}
	pathLength, path := Dijkstra(graph, "A", "E")
	assert.Equal(t, 4, pathLength)
	assert.Equal(t, []string{"E", "C", "B", "A"}, path)
}
