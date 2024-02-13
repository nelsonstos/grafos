package graph

type Node struct {
	Id    string           // Agregamos información sobre el ID del nodo
	Edges map[string]*Edge // Agregamos un mapa de aristas
	//Distance float64          // Agregamos  distacia para el balaceamiento del grafo
	Visited bool // Agregamos información sobre nodo visitado
}

func NewNode(id string) *Node {
	return &Node{
		Id:    id,
		Edges: make(map[string]*Edge), // Inicializamos el mapa de aristas
		//Distance: 0,                      // Inicializar la distancia en 0 por defecto
		Visited: false,
	}
}

func (n *Node) AddDirectedEdge(to string, weight float64) {
	// Creamos la arista y la agregamos al mapa de aristas del nodo
	n.Edges[to] = &Edge{To: to, Weight: weight}
}

func (n *Node) AddUndirectedEdge(to string, weight float64) {
	// Creamos la arista y la agregamos al mapa de aristas del nodo
	n.Edges[to] = &Edge{To: to, Weight: weight}
}

func (n *Node) RemoveEdge(to string) {
	// Eliminamos la arista del mapa de aristas del nodo
	delete(n.Edges, to)
}
