package graph

import "fmt"

func makeAdjList(n int, prerequisites [][]int) [][]int {
	adjList := make([][]int, n)
	for _, pair := range prerequisites {
		from := pair[1]
		to := pair[0]

		if len(adjList[from]) == 0 {
			adjList[from] = make([]int, 0, n-1)
		}
		adjList[from] = append(adjList[from], to)

	}

	for i, adj := range adjList {
		if adj == nil {
			adjList[i] = make([]int, 0)
		}
	}

	return adjList
}

func CanFinishCourses(n int, prerequisites [][]int) bool {
	adjList := makeAdjList(n, prerequisites)
	fmt.Println(adjList)
	visited := make([]bool, n)
	for node := 0; node < n; node++ {
		if !visited[node] && hasCycles(adjList, visited, node) {
			return false
		}
	}
	return true
}
