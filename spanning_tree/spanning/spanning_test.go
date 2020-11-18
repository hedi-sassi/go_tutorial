package spanning

import (
	"testing"
)


func TestPartialBFS(t *testing.T) {

	//graph is a triangle and a separate component with 2 vertices
	var (
		v1 Vertex
		v2 Vertex
		v3 Vertex
		v4 Vertex
		v5 Vertex
	)

	v1.neighbors = []*Vertex{&v2, &v3}
	v2.neighbors = []*Vertex{&v1, &v3}
	v3.neighbors = []*Vertex{&v2, &v1}

	v4.neighbors = []*Vertex{&v5}
	v4.neighbors = []*Vertex{&v5}


	var e1, e2, e3, e4 Edge

	e1.weight = 1
	e2.weight = 1
	e3.weight = 1
	e4.weight = 1

	e1.v1 = &v1
	e1.v2 = &v2
	e2.v1 = &v2
	e2.v2 = &v3
	e3.v1 = &v3
	e3.v2 = &v1

	e4.v1 = &v4
	e4.v2 = &v5

	//bfs discovered the isolated component but not the triangle

	discovered := []*Vertex{&v4, &v5}
	edges := []*Edge{&e1, &e2, &e3, &e4}

	remaining := NotDiscovered(edges, discovered)

	remaining_vertices := GetVertices(remaining)

	neighbors := GetNeighbors(&v1, []*Edge{&e1})

	if len(neighbors) != 1 || !Contains(neighbors, &v2){
		t.Fatalf("GetNeighbors function not correct")
	}

	if len(remaining) != 3 || len(remaining_vertices) != 3 || !Contains(remaining_vertices, &v1) || !Contains(remaining_vertices, &v2) || !Contains(remaining_vertices, &v3) {
		t.Fatalf("NotDiscovered or GetVertices functions not correct")
	}

	if NoCycles(remaining) || !NoCycles([]*Edge{&e4}) {
		
		t.Fatalf("NoCycles function not correct")
	}
}


func TestMST(t *testing.T) {

	

}