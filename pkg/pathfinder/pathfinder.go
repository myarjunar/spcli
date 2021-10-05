package pathfinder

import (
	"errors"
)

type Pathfinder struct {
	adjacencyList AdjacencyList
	path          string
}

var MAX = 9999

func NewPathfinder(graphFile string, path string) Pathfinder {
	adjacencyList := AdjacencyListFromCsv(graphFile)
	return Pathfinder{
		adjacencyList: adjacencyList,
		path:          path,
	}
}

func (p *Pathfinder) FindShortestPath() (int, error) {
	path := []rune(p.path)
	totalDistance, err := p.Dijkstra(string(path[0]), string(path[1]))

	return totalDistance, err
}

func (p *Pathfinder) FindPathLength() (int, error) {
	totalDistance := 0
	totalDistance, err := p.Traverse(p.path, totalDistance)

	return totalDistance, err
}

func (p *Pathfinder) Traverse(sequence string, currentDistance int) (int, error) {
	// for every node in the input path
	// 		if next node exist in the adj list of cur node
	// 			do it again
	//		else
	// 			if current node is not the last node from the input path
	// 				path not exist
	// 			else
	// 				return total distance

	nodeList := []rune(sequence)

	if len(nodeList) > 1 {
		currentNode := string(nodeList[0])
		nextNode := string(nodeList[1])

		neighbors, ok := p.adjacencyList[currentNode]
		if ok {
			distance, ok := neighbors[nextNode]
			if ok {
				return p.Traverse(string(nodeList[1:]), currentDistance+distance)
			} else {
				return currentDistance, errors.New("path not found")
			}
		}
	}

	return currentDistance, nil
}

func (p *Pathfinder) Dijkstra(start string, end string) (int, error) {
	// set initial value: visited & distance map, distance of start node
	// for every node except the start node
	//		mark node with min distance as visited
	//		for every neighbor in visited node
	//			if neighbor-is-not-visited and visited-node-distance-to-starting-node + visited-node-distance-to-neighbor lower than current-neighbor-to-starting-point-distance
	//				update distance from starting point to neighbor

	var visited = map[string]bool{}
	var distance = map[string]int{}

	for node := range p.adjacencyList {
		distance[node] = MAX
		visited[node] = false
	}

	distance[start] = 0

	for i := 0; i < len(p.adjacencyList)-1; i++ {
		visitedNode := p.MinDistance(distance, visited)
		visited[visitedNode] = true
		for neighbor := range p.adjacencyList[visitedNode] {
			if !visited[neighbor] && distance[visitedNode] != MAX && distance[visitedNode]+p.adjacencyList[visitedNode][neighbor] < distance[neighbor] {
				distance[neighbor] = distance[visitedNode] + p.adjacencyList[visitedNode][neighbor]
			}
		}
	}

	totalDistance, ok := distance[end]

	if !ok {
		return totalDistance, errors.New("end node is not in the graph")
	}

	if totalDistance == MAX {
		return totalDistance, errors.New("nodes are not connected")
	}

	return totalDistance, nil
}

func (p *Pathfinder) MinDistance(distance map[string]int, visited map[string]bool) string {
	var visitedNode string
	var min = MAX

	for node := range distance {
		if distance[node] <= min && !visited[node] {
			min = distance[node]
			visitedNode = node
		}
	}

	return visitedNode
}
