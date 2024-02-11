package graph

type Node struct {
    id    string
    edges map[string]*Edge 
}

func NewNode(id string) *Node {
    return &Node{
        id:    id,
        edges: make(map[string]*Edge), // Inicializamos el mapa de aristas
    }
}

func (n *Node) AddDirectedEdge(to string, weight float64) {
    // Creamos la arista y la agregamos al mapa de aristas del nodo
    n.edges[to] = &Edge{to: to, weight: weight}
}

func (n *Node) AddUndirectedEdge(to string, weight float64) {
    // Creamos la arista y la agregamos al mapa de aristas del nodo
    n.edges[to] = &Edge{to: to, weight: weight}
}

func (n *Node) RemoveEdge(to string) {
    // Eliminamos la arista del mapa de aristas del nodo
    delete(n.edges, to)
}

