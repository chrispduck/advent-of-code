package algorithms

import "math"

type GraphNode struct {
	name  string
	edges map[string]int
}

type Graph struct {
	Nodes map[string]GraphNode
}

func Dijkstra(graph Graph, start string, end string) (pathLength int, path []string) {
	visited := make(map[string]bool)
	distance := make(map[string]int)
	for _, node := range graph.Nodes {
		// initialize distance to infinity
		distance[node.name] = math.MaxInt
	}

	distance[start] = 0
	visited[start] = true
	for len(visited) < len(graph.Nodes) {
		// find the node with the smallest distance
		var smallestNode string
		smallestEdge := math.MaxInt
		fromNode := ""
		// check all the nodes that have been visited
		for v, _ := range visited {
			for nodeName, edgeLength := range graph.Nodes[v].edges {
				if !visited[nodeName] && edgeLength < smallestEdge {
					smallestEdge = edgeLength
					smallestNode = nodeName
					fromNode = v
				}
			}
		}
		distance[smallestNode] = distance[fromNode] + smallestEdge
		visited[smallestNode] = true
	}

	// backtrack to find path
	path = []string{end}
	currentNode := end
	for {
		for nodeName, edgeLength := range graph.Nodes[currentNode].edges {
			if distance[nodeName]+edgeLength == distance[currentNode] {
				path = append(path, nodeName)
				currentNode = nodeName
				break
			}
		}
		if currentNode == start {
			break
		}
	}
	return distance[end], path
}
