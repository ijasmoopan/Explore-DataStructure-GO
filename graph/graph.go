package main

import (
	"fmt"
)

// Graph represents an adjacency list graph
type Graph struct {
	Vertices []*Vertex
}
// Vertex represents a graph vertex
type Vertex struct {
	Key int
	Edge []*Vertex
}

// AddVertex adds a vertex to the Graph
func (g *Graph) AddVertex (k int) {
	if Contains(g.Vertices, k) {
		err := fmt.Errorf("%v is already existing", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

// AddEdge
func (g *Graph) AddEdge (from, to int) {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if Contains(fromVertex.Edge, to) {
		err := fmt.Errorf("Already existing (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
	// add edge
		fromVertex.Edge = append(fromVertex.Edge, toVertex)
	}
}

// GetVertex
func (g *Graph) GetVertex (k int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}

// Contains
func Contains (s []*Vertex, k int) bool {
	for _, v := range s {
		if v.Key == k {
			return true
		}
	}
	return false
} 

// Print
func (g *Graph) Print () {
	for _, v := range g.Vertices {
		fmt.Printf("\nvertex %v : ", v.Key)
		for _, v := range v.Edge {
			fmt.Printf(" %v ", v.Key)
		}
	}
}

func main() {
	graph := &Graph{}
	for i:= 0; i<5;i++ {
		graph.AddVertex(i)
	}
	graph.AddVertex(0)
	graph.AddVertex(4)
	// fmt.Println(graph.Vertices)
	graph.AddEdge(5, 1)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 2)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 4)
	graph.Print()
}