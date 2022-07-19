package algo

import . "turing/ds"

func BFSSimple(root NodeValue, graph Graph) {
	nodeQueue := make([]NodeValue, 0, graph.Size())
	nodeQueue = append(nodeQueue, root)
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		graph.LogVisit(node)
		for _, neighbor := range graph.NeighborsOf(node) {
			if !graph.IsVisited(neighbor) {
				graph.Visit(neighbor)
				nodeQueue = append(nodeQueue, neighbor)
			}
		}
	}
}
