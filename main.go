package main

import (
	"fmt"

	"github.com/nelsonstos/grafos/graph"
	"github.com/nelsonstos/grafos/utils"
)

func main() {
	/*******************************************************************/
	// Crear un grafo dirigido
	/*******************************************************************/
	directedGraph := graph.NewGraph(true)

	// Agregar nodos al grafo
	for i := 1; i <= 10; i++ {
		nodeID := fmt.Sprintf("Node%d", i)
		directedGraph.AddNode(nodeID)
	}

	// Agregar aristas al grafo
	directedGraph.AddEdge("Node1", "Node2", 1.0)
	directedGraph.AddEdge("Node1", "Node3", 2.0)
	directedGraph.AddEdge("Node2", "Node4", 3.0)
	directedGraph.AddEdge("Node2", "Node5", 1.5)
	directedGraph.AddEdge("Node3", "Node6", 2.5)
	directedGraph.AddEdge("Node4", "Node7", 4.0)
	directedGraph.AddEdge("Node5", "Node8", 2.0)
	directedGraph.AddEdge("Node6", "Node9", 3.0)
	directedGraph.AddEdge("Node7", "Node10", 2.0)
	directedGraph.AddEdge("Node8", "Node10", 1.0)

	// Imprimir el grafo dirigido antes de eliminar un nodo
	fmt.Println("GRAFO DIRIGIDO:")
	directedGraph.Print()

	// Eliminar un nodo del grafo
	directedGraph.RemoveNode("Node3")

	// Generar un archivo DOT para el grafo dirigido
	directedDot := utils.GenerateDot(directedGraph)
	// Llamar a Graphviz para generar las imágenes de los grafos
	utils.GenerateGraphImage("directed_graph.png", directedDot)

	/*******************************************************************/
	// Crear un grafo no dirigido
	/*******************************************************************/
	undirectedGraph := graph.NewGraph(false)

	// Agregar nodos al grafo
	for i := 1; i <= 10; i++ {
		nodeID := fmt.Sprintf("Node%d", i)
		undirectedGraph.AddNode(nodeID)
	}

	// Agregar aristas al grafo
	undirectedGraph.AddEdge("Node1", "Node2", 1.0)
	undirectedGraph.AddEdge("Node1", "Node3", 2.0)
	undirectedGraph.AddEdge("Node2", "Node4", 3.0)
	undirectedGraph.AddEdge("Node2", "Node5", 1.5)
	undirectedGraph.AddEdge("Node3", "Node6", 2.5)
	undirectedGraph.AddEdge("Node4", "Node7", 4.0)
	undirectedGraph.AddEdge("Node5", "Node8", 2.0)
	undirectedGraph.AddEdge("Node6", "Node9", 3.0)
	undirectedGraph.AddEdge("Node7", "Node10", 2.0)
	undirectedGraph.AddEdge("Node8", "Node10", 1.0)

	// Imprimir el grafo no dirigido antes de eliminar un nodo
	fmt.Println("\nGRAFO NO DIRIGIDO:")
	undirectedGraph.Print()

	// Eliminar un nodo del grafo
	undirectedGraph.RemoveNode("Node3")

	// Generar un archivo DOT para el grafo no dirigido
	undirectedDot := utils.GenerateDot(undirectedGraph)

	// Llamar a Graphviz para generar las imágenes de los grafos
	utils.GenerateGraphImage("undirected_graph.png", undirectedDot)

}
