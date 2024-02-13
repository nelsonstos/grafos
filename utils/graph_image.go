package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nelsonstos/grafos/graph"
)

// Función para generar un archivo DOT a partir de un grafo
func GenerateDot(g *graph.Graph) string {
	var dot string

	if g.Directed {
		dot = "digraph G {\n"
	} else {
		dot = "graph G {\n"
	}

	for id, node := range g.Nodes {
		for to, edge := range node.Edges {
			if g.Directed {
				dot += fmt.Sprintf("\t%s -> %s [label=\"%.2f\"];\n", id, to, edge.Weight)
			} else {
				if id < to {
					dot += fmt.Sprintf("\t%s -- %s [label=\"%.2f\"];\n", id, to, edge.Weight)
				}
			}
		}
	}

	dot += "}\n"

	return dot
}

// Función para generar una imagen del grafo utilizando Graphviz
func GenerateGraphImage(filename string, dot string) {
	file, err := os.Create("graph.dot")
	if err != nil {
		fmt.Println("Error creando el archivo DOT:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(dot)
	if err != nil {
		fmt.Println("Error escribiendo en el archivo DOT:", err)
		return
	}

	cmd := exec.Command("dot", "-Tpng", "graph.dot", "-o", filename)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error generando la imagen del grafo:", err)
		return
	}
	fmt.Println("Imagen del grafo generada correctamente:", filename)
}
