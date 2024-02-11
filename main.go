package main

import (
	"fmt"
	graph "github.com/nelsonstos/grafos/graph"
)

func main() {
	// Crear un grafo dirigido
	directedGraph := graph.NewGraph(true)

	// Agregar nodos y aristas con pesos
	directedGraph.AddEdge("A", "B", 5)
	directedGraph.AddEdge("A", "C", 3)
	directedGraph.AddEdge("B", "C", 2)
	directedGraph.AddEdge("C", "D", 4)

	// Imprimir el grafo dirigido antes de eliminar un nodo
	fmt.Println("Grafo Dirigido antes de eliminar un nodo:")
	directedGraph.Print()

	// Eliminar el nodo "B" del grafo dirigido
	directedGraph.RemoveNode("B")

	// Imprimir el grafo dirigido después de eliminar un nodo
	fmt.Println("\nGrafo Dirigido después de eliminar un nodo:")
	directedGraph.Print()

	// Crear un grafo no dirigido
	undirectedGraph := graph.NewGraph(false)

	// Agregar nodos y aristas con pesos
	undirectedGraph.AddEdge("A", "B", 5)
	undirectedGraph.AddEdge("A", "C", 3)
	undirectedGraph.AddEdge("B", "C", 2)
	undirectedGraph.AddEdge("C", "D", 4)

	// Imprimir el grafo no dirigido antes de eliminar un nodo
	fmt.Println("\nGrafo No Dirigido antes de eliminar un nodo:")
	undirectedGraph.Print()

	// Eliminar el nodo "B" del grafo no dirigido
	undirectedGraph.RemoveNode("B")

	// Imprimir el grafo no dirigido después de eliminar un nodo
	fmt.Println("\nGrafo No Dirigido después de eliminar un nodo:")
	undirectedGraph.Print()
}
