package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// Estructura de datos para representar un grafo
type Graph struct {
	Nodes    map[string]*Node // Agregamos un mapa de nodos
	Directed bool             // Agregamos un valor booleano que indica si el grafo es dirigido o no
}

// Crea un nuevo grafo
func NewGraph(directed bool) *Graph {
	return &Graph{
		Nodes:    make(map[string]*Node),
		Directed: directed,
	}
}

func (g *Graph) AddNode(id string) {
	if _, exists := g.Nodes[id]; !exists {
		g.Nodes[id] = NewNode(id)
	}
}

func (g *Graph) RemoveNode(id string) {
	if node, exists := g.Nodes[id]; exists {
		for _, edge := range node.Edges {
			delete(g.Nodes[edge.To].Edges, id)
		}
		delete(g.Nodes, id)
	}
}

func (g *Graph) AddEdge(from, to string, weight float64) {
	if _, exists := g.Nodes[from]; !exists {
		g.AddNode(from)
	}
	if _, exists := g.Nodes[to]; !exists {
		g.AddNode(to)
	}

	edge := &Edge{To: to, Weight: weight}
	g.Nodes[from].Edges[to] = edge

	if !g.Directed {
		g.Nodes[to].Edges[from] = edge
	}
}

func (g *Graph) RemoveEdge(from, to string) {
	if _, exists := g.Nodes[from]; exists {
		delete(g.Nodes[from].Edges, to)
		if !g.Directed {
			delete(g.Nodes[to].Edges, from)
		}
	}
}

func (g *Graph) Print() {
	for id, node := range g.Nodes {
		fmt.Printf("Node: %s, Edges: ", id)
		for to, edge := range node.Edges {
			if !g.Directed {
				fmt.Printf("(%s -- %s, Weight: %.2f) ", id, to, edge.Weight)
			} else {
				fmt.Printf("(%s -> %s, Weight: %.2f) ", id, to, edge.Weight)
			}
		}
		fmt.Println()
	}
}

func (g *Graph) Balance(sourceID string) map[string]float64 {
	// Inicializar las distancias de todos los nodos como infinito
	distances := make(map[string]float64)
	for id := range g.Nodes {
		distances[id] = math.Inf(1)
	}

	// Inicializar la distancia del nodo fuente como 0
	distances[sourceID] = 0

	// Inicializar la cola de prioridad
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: g.Nodes[sourceID], priority: 0})

	// Iterar hasta que la cola de prioridad esté vacía
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).node

		// Iterar sobre los vecinos del nodo actual
		for _, edge := range current.Edges {
			var neighbor *Node

			if g.Directed || (edge.To == current.Id) {
				neighbor = g.Nodes[edge.To]
			} else {
				neighbor = g.Nodes[edge.From]
			}

			// Calcular la distancia tentativa desde el nodo fuente hasta este vecino
			tentativeDistance := distances[current.Id] + edge.Weight

			// Si la distancia tentativa es menor que la distancia actual almacenada,
			// actualizar la distancia y agregar el vecino a la cola de prioridad
			if tentativeDistance < distances[neighbor.Id] {
				distances[neighbor.Id] = tentativeDistance
				heap.Push(&pq, &Item{node: neighbor, priority: tentativeDistance})
			}
		}
	}

	return distances
}
