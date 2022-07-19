package graph

type Node struct {
	id       int
	pathSize int
}

func ShortestPath(graph [][]int, src int, dst int) int {
	next := make([]Node, 0, len(graph))
	shortest := -1

	for next = append(next, Node{id: src, pathSize: 0}); len(next) != 0; {
		node := next[0]
		visited[node.id] = true
		visit(node.id)

		if node.id == dst {
			return node.pathSize
		}

		for _, neighbour := range graph[node.id] {

			if !visited[neighbour] {
				next = append(next, Node{id: neighbour, pathSize: node.pathSize + 1})
				visited[neighbour] = true
			}
		}
		endVisit(node.id)
		next = next[1:]
	}

	return shortest
}

func FindShortestPath(src int, dst int) int {
	return ShortestPath(graph, src, dst)
}
