package ds

func DAG() Graph {
	g1 := &GraphImpl{}

	g1.New(10)
	g1.NewEdge(3, 4)
	g1.NewEdge(0, 1)
	g1.NewEdge(9, 5)
	g1.NewEdge(0, 2)
	g1.NewEdge(0, 3)
	g1.NewEdge(2, 3)
	g1.NewEdge(4, 1)
	g1.NewEdge(4, 5)
	g1.NewEdge(7, 5)
	g1.NewEdge(5, 6)
	g1.NewEdge(1, 7)
	g1.NewEdge(7, 8)
	g1.NewEdge(3, 9)

	return g1
}

func DirectedCyclicGraph() Graph {
	g1 := &GraphImpl{}

	g1.New(10)
	g1.NewEdge(3, 4)
	g1.NewEdge(0, 1)
	g1.NewEdge(9, 5)
	g1.NewEdge(0, 2)
	g1.NewEdge(0, 3)
	g1.NewEdge(2, 3)
	g1.NewEdge(4, 1)
	g1.NewEdge(6, 9)
	g1.NewEdge(4, 5)
	g1.NewEdge(7, 5)
	g1.NewEdge(5, 6)
	g1.NewEdge(1, 7)
	g1.NewEdge(7, 8)
	g1.NewEdge(3, 9)

	return g1
}
