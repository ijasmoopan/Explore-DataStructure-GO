package main

import "fmt"

// Graph
type Graph struct {
	matrix  [][]int
	visited [][]bool
}

// CheckRiverSize
func (g *Graph) CheckRiverSize() []int {
	g.visited = make([][]bool, len(g.matrix))

	sizes := []int{}
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix[0]); j++ {
			if g.visited[i][j] {
				continue
			}
			sizes = g.TraverseThroughRiver(i, j, sizes)
		}
	}
	return sizes
}

// TraverseThroughRiver
func (g *Graph) TraverseThroughRiver(i int, j int, sizes []int) []int {
	// g.visited = [len(g.matrix)][len(g.matrix[0])]bool {}
	currentRiverSize := 0
	stack := [][]int{}
	stack = append(stack, []int{i, j})

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i = top[0]
		j = top[1]
		if g.visited[i][j] {
			continue
		}
		g.visited[i][j] = true
		if g.matrix[i][j] == 0 {
			continue
		}
		currentRiverSize++
		nearest := g.GetNearestOfNodes(i, j)
		for _, v := range nearest {
			stack = append(stack, v)
		}
	}
	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}
	return sizes
}

// GetNearestOfNodes
func (g *Graph) GetNearestOfNodes(i int, j int) [][]int {
	nearest := [][]int{}
	if i > 0 && !g.visited[i-1][j] {
		nearest = append(nearest, []int{i - 1, j})
	}
	if i < len(g.matrix)-1 && !g.visited[i+1][j] {
		nearest = append(nearest, []int{i + 1, j})
	}
	if j > 0 && !g.visited[i][j-1] {
		nearest = append(nearest, []int{i, j - 1})
	}
	if j < len(g.matrix[0])-1 && !g.visited[i][j+1] {
		nearest = append(nearest, []int{i, j + 1})
	}
	return nearest
}

func main() {
	graph := &Graph{}

	graph.matrix = [][]int{{1, 0, 0, 1, 1},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 1},
		{1, 0, 0, 0, 1}}

	// graph.visited = [len(graph.matrix)][len(graph.matrix[0])]bool{}
	// graph.visited = [][]bool{{false, false, false, false, false},
	// 	{false, false, false, false, false},
	// 	{false, false, false, false, false},
	// 	{false, false, false, false, false}}
	fmt.Println(graph.matrix)
	result := graph.CheckRiverSize()
	fmt.Println(result)
}
