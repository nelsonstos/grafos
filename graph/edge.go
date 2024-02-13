package graph

type Edge struct {
	From    string  // Agregamos información sobre el nodo de origen
	To      string  // Agregamos información sobre el nodo de destino
	Weight  float64 // Agregamos información sobre la distancia entre el nodo de origen y el nodo destino
	Visited bool    // Agregar el campo visited
}
