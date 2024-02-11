# Grafos ponderados dirigidos y no dirigidos

## Resumen

Este repositorio ofrece una implementación de la lógica necesaria para la creación y eliminación de nodos, así como la capacidad de agregar aristas a grafos dirigidos y no dirigidos.

## Características 
El package permite: 
- Representación de nodos,  aristas(conexiones entre nodos) y sus pesos en un grafo a través del package graph.
- Representación de nodos dirigidos y no dirigidos.
- Presenta métodos para agregar ,eliminar nodos y aristas de manera dinámica.
- Facilita la reutilización.

## Estructura del directorio

```
grafos/
|-- main.go
|-- graph/
|   |-- edge.go
|   |-- graph.go
|   |-- nodo.go
|-- go.mod
|-- go.sum
|-- README.md
```

## Ejecución

### Prerequisitos

- Go 1.18+

### Instalación

1. Clonar repositorio

```bash
git clonehttps://github.com/nelsonstos/grafos.git
```

2. Navegue hasta el directorio

```bash
cd grafos
```

3. Ejecutar la aplicación

```bash
go run main.go
```

### Resultados
```bash
PS D:\projects\challenge\grafos> go run main.go                                              
Grafo Dirigido antes de eliminar un nodo:
Node: A, Edges: (A -> B, Weight: 5.00) (A -> C, Weight: 3.00)
Node: B, Edges: (B -> C, Weight: 2.00)
Node: C, Edges: (C -> D, Weight: 4.00)
Node: D, Edges:

Grafo Dirigido después de eliminar un nodo:
Node: A, Edges: (A -> B, Weight: 5.00) (A -> C, Weight: 3.00)
Node: C, Edges: (C -> D, Weight: 4.00)
Node: D, Edges:

Grafo No Dirigido antes de eliminar un nodo:
Node: B, Edges: (B -- A, Weight: 5.00) (B -- C, Weight: 2.00)
Node: C, Edges: (C -- A, Weight: 3.00) (C -- B, Weight: 2.00) (C -- D, Weight: 4.00)
Node: D, Edges: (D -- C, Weight: 4.00)
Node: A, Edges: (A -- B, Weight: 5.00) (A -- C, Weight: 3.00)

Grafo No Dirigido después de eliminar un nodo:
Node: D, Edges: (D -- C, Weight: 4.00)
Node: A, Edges: (A -- B, Weight: 5.00) (A -- C, Weight: 3.00)
Node: C, Edges: (C -- A, Weight: 3.00) (C -- D, Weight: 4.00)
PS D:\projects\challenge\grafos>
```