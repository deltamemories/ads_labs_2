package main

import (
	"fmt"
)

// Graph represents an adjacency list.
type Graph struct {
	V   int
	Adj [][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		V:   v,
		Adj: make([][]int, v),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
	g.Adj[v] = append(g.Adj[v], u)
}

// GreedyColoring assigns colors to vertices such that no two adjacent vertices have the same color.
func (g *Graph) GreedyColoring() []int {
	result := make([]int, g.V)
	for i := range result {
		result[i] = -1
	}

	// Assign the first color to the first vertex
	result[0] = 0

	// available[i] = false means color i is assigned to one of its adjacent vertices
	available := make([]bool, g.V)
	for i := range available {
		available[i] = true
	}

	for u := 1; u < g.V; u++ {
		// Process all adjacent vertices and flag their colors as unavailable
		for _, v := range g.Adj[u] {
			if result[v] != -1 {
				available[result[v]] = false
			}
		}

		// Find the first available color
		cr := 0
		for cr < g.V {
			if available[cr] {
				break
			}
			cr++
		}

		result[u] = cr

		// Reset available colors for the next iteration
		for _, v := range g.Adj[u] {
			if result[v] != -1 {
				available[result[v]] = true
			}
		}
	}

	return result
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	colors := g.GreedyColoring()

	fmt.Println("Graph Coloring (Greedy):")
	for i, color := range colors {
		fmt.Printf("Vertex %d ---> Color %d\n", i, color)
	}

	maxColor := 0
	for _, c := range colors {
		if c > maxColor {
			maxColor = c
		}
	}
	fmt.Printf("Total colors used: %d\n", maxColor+1)
}
