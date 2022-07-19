package scanner

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func BuildAdjList(input io.ReadCloser, directed bool) (adjList [][]int) {
	scanner := bufio.NewScanner(input)

	size := 0

	if scanner.Scan() {
		size = readIntOrPanic(scanner.Text())
		adjList = make([][]int, size)
	} else {
		return
	}

	for scanner.Scan() {
		edge := strings.Split(scanner.Text(), " ")
		if len(edge) != 2 {
			continue
		}

		nodes := readInts(edge)
		if len(nodes) != 2 {
			continue
		}

		from := nodes[0]
		to := nodes[1]

		neighbours := getNeighbourSlice(adjList, from)
		adjList[from] = append(neighbours, to)

		if !directed {
			neighbours = getNeighbourSlice(adjList, to)
			adjList[to] = append(neighbours, from)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, adj := range adjList {
		if adj == nil {
			adjList[i] = make([]int, 0)
		}
	}

	return
}

func getNeighbourSlice(adjList [][]int, node int) []int {
	if cap(adjList[node]) == 0 {
		adjList[node] = make([]int, 0, 2)
	}
	return adjList[node]
}

func readInts(numbers []string) []int {
	var emptyValues []int
	values := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if value, err := strconv.Atoi(number); err == nil {
			values = append(values, value)
			continue
		}
	}
	if len(values) != len(numbers) {
		return emptyValues
	}
	return values
}

func readIntOrPanic(number string) int {
	var err error
	if value, err := strconv.Atoi(number); err == nil {
		return value
	}
	log.Panic(err)
	return 0
}
