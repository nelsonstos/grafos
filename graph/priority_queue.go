package graph

import (
	"container/heap"
)

// Item representa un nodo en la cola de prioridad
type Item struct {
	node     *Node   // Puntero al nodo
	priority float64 // Prioridad del nodo (distancia desde el nodo fuente)
	index    int     // Índice del elemento en la cola de prioridad
}

// PriorityQueue es una cola de prioridad implementada como un montículo mínimo
type PriorityQueue []*Item

// Len retorna la longitud de la cola de prioridad
func (pq PriorityQueue) Len() int { return len(pq) }

// Less determina si el elemento en el índice i es menor que el elemento en el índice j
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// Swap intercambia los elementos en los índices i y j
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push agrega un elemento a la cola de prioridad
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop elimina y devuelve el elemento más pequeño de la cola de prioridad
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // Marcar como eliminado
	*pq = old[0 : n-1]
	return item
}

// Update modifica la prioridad y el valor de un elemento en la cola de prioridad
func (pq *PriorityQueue) Update(item *Item, priority float64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}
