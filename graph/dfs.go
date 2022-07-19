package graph

import "fmt"

const MAX_NEIGHBORS = 4

type World [][]Point

type Point struct {
	tileType rune
	column   int
	row      int
}

type DiscreteWorld interface {
	At(row int, column int) *Point
	Rows() int
	Columns() int
	isLand(point *Point) bool
	isLandAt(row int, column int) bool
}

func (w World) At(row, column int) *Point {
	if len(w) <= 0 || len(w[0]) <= 0 {
		return nil
	}
	if row < 0 || column < 0 {
		return nil
	}
	if row >= len(w) || column >= len(w[0]) {
		return nil
	}
	return &w[row][column]
}

func (w World) Rows() int {
	return len(w)
}

func (w World) Columns() int {
	return len(w[0])
}

func (w World) isLandAt(row int, column int) bool {
	at := w.At(row, column)
	return at != nil && at.tileType != 'W'
}

func (w World) isLand(p *Point) bool {
	return p != nil && w.isLandAt(p.row, p.column)
}

func (p *Point) GetNeighbors(world DiscreteWorld) []*Point {
	neighbors := make([]*Point, 0, MAX_NEIGHBORS)
	top := world.At(p.row-1, p.column)
	right := world.At(p.row, p.column+1)
	down := world.At(p.row+1, p.column)
	left := world.At(p.row, p.column-1)

	if world.isLand(top) {
		neighbors = append(neighbors, top)
	}
	if world.isLand(right) {
		neighbors = append(neighbors, right)
	}
	if world.isLand(down) {
		neighbors = append(neighbors, down)
	}
	if world.isLand(left) {
		neighbors = append(neighbors, left)
	}

	return neighbors
}

func hasCycles(graph [][]int, visited []bool, node int) bool {
	traversalPath = append(traversalPath, node)
	visited[node] = true
	for _, neighbor := range graph[node] {
		if visited[neighbor] || hasCycles(graph, visited, neighbor) {
			return true
		}
	}
	return false
}

func DFS(node int) {
	visit(node)
	traversalPath = append(traversalPath, node)
	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			components[neighbour] = components[node]
			DFS(neighbour)
		} else {
			alreadyVisited(neighbour)
		}
	}
	endVisit(node)
}

func ParseMap(world [][]rune) World {
	rows := make([][]Point, len(world))
	for row, w := range world {
		if rows[row] == nil {
			rows[row] = make([]Point, len(world[0]))
		}
		for column := range w {
			rows[row][column] = Point{row: row, column: column, tileType: world[row][column]}
		}
	}
	return rows
}

func IslandCount(world DiscreteWorld) (int, int) {
	columns := world.Columns()
	rows := world.Rows()
	count := 0
	minimum := world.Columns() * world.Rows()

	visited := make(map[int]bool)

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if _, ok := visited[columns*row+column]; !ok && world.isLandAt(row, column) {
				count++
				islandSize := ExploreIslandFrom(visited, world, row, column)
				if islandSize <= minimum {
					minimum = islandSize
				}
			}
		}
	}

	return count, minimum
}

func ExploreIslandFrom(visited map[int]bool, world DiscreteWorld, row int, column int) int {
	count := 1
	visited[world.Columns()*row+column] = true
	neighbors := world.At(row, column).GetNeighbors(world)
	for _, neighbor := range neighbors {
		if _, ok := visited[world.Columns()*neighbor.row+neighbor.column]; !ok {
			count += ExploreIslandFrom(visited, world, neighbor.row, neighbor.column)
		}
	}
	return count
}

func DFSTraversalPath(node int, path *[]int) int {
	visit(node)
	size := 1
	*path = append(*path, node)
	// fmt.Println(len(path))
	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			size += DFSTraversalPath(neighbour, path)
		} else {
			alreadyVisited(neighbour)
		}
	}
	endVisit(node)
	return size
}

func DFSUntil(node int, dest int) bool {
	visit(node)
	if node == dest {
		return true
	}
	for _, neighbour := range graph[node] {
		if !visited[neighbour] {
			DFSUntil(neighbour, dest)
		}
		alreadyVisited(neighbour)
	}
	endVisit(node)
	return false
}

func DFSFrom(node int) {
	DFS(node)
}

func FindComponents() {
	count := 0
	for node := range graph {
		if !visited[node] {
			components[node] = count
			DFS(node)
			count++
		}
	}
}

func FindLargestComponent() {
	largest := 0
	for node := range graph {
		if !visited[node] {
			path := make([]int, 0, len(graph))
			componentSize := DFSTraversalPath(node, &path)
			if componentSize >= largest {
				largest = componentSize
			}
			fmt.Println(componentSize, len(path))
		}
	}
}

func HasPath(from int, to int) bool {
	return DFSUntil(from, to)
}
