package graph

import (
    "fmt"
)

type Graph struct {
    nodes    map[string]*Node
    directed bool
}

func NewGraph(directed bool) *Graph {
    return &Graph{
        nodes:    make(map[string]*Node),
        directed: directed,
    }
}

func (g *Graph) AddNode(id string) {
    if _, exists := g.nodes[id]; !exists {
        g.nodes[id] = NewNode(id)
    }
}

func (g *Graph) RemoveNode(id string) {
    if node, exists := g.nodes[id]; exists {
        for _, edge := range node.edges {
            delete(g.nodes[edge.to].edges, id)
        }
        delete(g.nodes, id)
    }
}

func (g *Graph) AddEdge(from, to string, weight float64) {
    if _, exists := g.nodes[from]; !exists {
        g.AddNode(from)
    }
    if _, exists := g.nodes[to]; !exists {
        g.AddNode(to)
    }

    edge := &Edge{to: to, weight: weight}
    g.nodes[from].edges[to] = edge

    if !g.directed {
        g.nodes[to].edges[from] = edge
    }
}

func (g *Graph) RemoveEdge(from, to string) {
    if _, exists := g.nodes[from]; exists {
        delete(g.nodes[from].edges, to)
        if !g.directed {
            delete(g.nodes[to].edges, from)
        }
    }
}


func (g *Graph) Print() {
    for id, node := range g.nodes {
        fmt.Printf("Node: %s, Edges: ", id)
        for to, edge := range node.edges {
			if !g.directed {
				fmt.Printf("(%s -- %s, Weight: %.2f) ", id, to, edge.weight)
			} else {
				fmt.Printf("(%s -> %s, Weight: %.2f) ", id, to, edge.weight)
			}
        }
        fmt.Println()
    }
}
