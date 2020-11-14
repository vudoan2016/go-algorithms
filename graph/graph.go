package graph

import (
	"fmt"
	"sort"
)

type Id int
type Weight int

const (
	MAX int = 1<<32 - 1
)

// Directed weighted graph
type Vertex struct {
	id       Id
	last     Id // last vertex on the shortest path
	visited  bool
	edges    map[Id]Weight
	distance int // shortest distance from source
}
type Graph struct {
	vertices map[Id]*Vertex
}

func New() *Graph {
	return &Graph{
		vertices: map[Id]*Vertex{},
	}
}

func (g *Graph) addVertex(id Id) {
	g.vertices[id] = &Vertex{id: id, edges: make(map[Id]Weight)}
}

// Edge points from 'fromId' to 'toId'
func (g *Graph) addEdge(fromId Id, toId Id, weight Weight) {
	g.vertices[fromId].edges[toId] = weight
}

func (g Graph) print() {
	for _, n := range g.vertices {
		fmt.Printf("Vertex %d: ", n.id)
		for toId, w := range n.edges {
			fmt.Printf("[%d, %d] ", toId, w)
		}
		fmt.Println()
	}
}

func (g Graph) traverseHelper(source Id, stack *[]Id) {
	if _, found := g.vertices[source]; !found {
		return
	}
	g.vertices[source].visited = true
	*stack = append(*stack, source)

	for neighbor := range g.vertices[source].edges {
		if !g.vertices[neighbor].visited {
			g.traverseHelper(neighbor, stack)
			if (*stack)[0] == source {
				fmt.Println(*stack)
				*stack = (*stack)[:1]
			}
		}
	}
}

func (g Graph) dfTraverse(root Id) {
	for _, n := range g.vertices {
		n.visited = false
	}
	var stack []Id
	g.traverseHelper(root, &stack)
}

func (g Graph) dfsHelper(root, key Id) bool {
	if root == key {
		return true
	}
	g.vertices[root].visited = true
	for neighbor := range g.vertices[root].edges {
		if !g.vertices[neighbor].visited {
			if g.dfsHelper(neighbor, key) {
				return true
			}
		}
	}
	return false
}

func (g Graph) dfs(root, key Id) bool {
	for _, n := range g.vertices {
		n.visited = false
	}
	return g.dfsHelper(root, key)
}

// BFS is used to find SP between 2 vertices of an unweighted graph
func (g Graph) bfsTraverse(root Id) {
	for _, n := range g.vertices {
		n.visited = false
	}
	q := []Id{}
	q = append(q, root)
	for len(q) > 0 {
		fmt.Printf("%d ", q[0])
		for neighbor := range g.vertices[q[0]].edges {
			if !g.vertices[neighbor].visited {
				g.vertices[q[0]].visited = true
				q = append(q, neighbor)
			}
		}
		q = q[1:]
	}
	fmt.Println()
}

type VertexArray []*Vertex

// Receiver functions
func (a VertexArray) Len() int           { return len(a) }
func (a VertexArray) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a VertexArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a VertexArray) print(name string) {
	fmt.Print(name, ": ")
	for _, vertex := range a {
		fmt.Printf("<id=%d, last=%d, d=%d>, ", vertex.id, vertex.last, vertex.distance)
	}
	fmt.Println()
}

// Display shortest paths from source to other vertices
func (g Graph) dijkstraShow(source Id) {
	sp := []Id{}
	for _, vertex := range g.vertices {
		if vertex.id != source {
			sp = append(sp, vertex.id)
			for g.vertices[sp[0]].last != source {
				sp = append([]Id{g.vertices[sp[0]].last}, sp...) // prepend
			}
			fmt.Println(source, " --> ", sp, "d =", g.vertices[sp[len(sp)-1]].distance)
		}
		sp = sp[:0]
	}
}

// O(E + E*V) for each vertex?
func (g Graph) dijkstra(source Id) {
	done := VertexArray{}
	process := VertexArray{}

	// Initialize vertices with MAX disstance except source
	for _, v := range g.vertices {
		if v.id != source {
			v.distance = MAX
		}
		process = append(process, v)
	}
	for len(process) > 0 {
		// sort remaining vertices
		sort.Slice(process, func(i, j int) bool { return process[i].distance < process[j].distance }) // anonymous function

		// sort.Sort(VertexArray(process))

		process.print("process")
		done.print("done")

		// relax edges with neighbors
		for neighbor, weight := range process[0].edges {
			if neighbor != process[0].id && process[0].distance+int(weight) < g.vertices[neighbor].distance {
				g.vertices[neighbor].distance = process[0].distance + int(weight)
				g.vertices[neighbor].last = process[0].id
			}
		}
		done = append(done, process[0])
		process = process[1:]
	}
}

func GraphTest() {
	/* 	g := New()
	   	g.addVertex(7)
	   	g.addVertex(11)
	   	g.addVertex(9)
	   	g.addVertex(28)
	   	g.addVertex(5)
	   	g.addEdge(7, 11, 4)
	   	g.addEdge(11, 9, 6)
	   	g.addEdge(9, 28, 8)
	   	g.addEdge(7, 5, 6)
	   	g.addEdge(5, 28, 2)
	   	g.print()
	   	g.dfTraverse(7)
	   	root := Id(11)
	   	key := Id(5)
	   	fmt.Printf("Key %d found: %t\n", key, g.dfs(root, key))
	   	g.bfsTraverse(7)
	*/
	g1 := New()
	g1.addVertex(11)
	g1.addVertex(22)
	g1.addVertex(33)
	g1.addVertex(44)
	g1.addVertex(55)
	g1.addVertex(66)
	g1.addVertex(77)

	g1.addEdge(11, 22, 5)
	g1.addEdge(22, 11, 5)

	g1.addEdge(22, 33, 6)
	g1.addEdge(33, 22, 6)

	g1.addEdge(11, 44, 3)
	g1.addEdge(44, 11, 3)

	g1.addEdge(44, 55, 2)
	g1.addEdge(55, 44, 2)

	g1.addEdge(11, 55, 6)
	g1.addEdge(55, 11, 6)

	g1.addEdge(22, 55, 2)
	g1.addEdge(55, 22, 2)

	g1.addEdge(22, 77, 3)
	g1.addEdge(77, 22, 3)

	g1.addEdge(22, 66, 7)
	g1.addEdge(66, 22, 7)

	g1.addEdge(55, 77, 9)
	g1.addEdge(77, 55, 9)

	g1.addEdge(33, 77, 5)
	g1.addEdge(77, 33, 5)

	g1.addEdge(77, 66, 1)
	g1.addEdge(66, 77, 1)

	g1.addEdge(33, 66, 2)
	g1.addEdge(66, 33, 2)

	g1.dijkstra(11)
	g1.dijkstraShow(11)
}
