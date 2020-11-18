
package spanning

import (
	"sort"
	"container/list"
)

//Edges have a cost and 2 Vertices
type Edge struct {
	v1,v2 *Vertex
	weight uint
}

//vertices have a list of edges
type Vertex struct {
	neighbors []*Vertex
}

//returns true if the vertex is contained in the list
func Contains(list []*Vertex, v *Vertex) bool {

	for _, vertex := range list{
		if vertex == v {
			return true
		}
	}

	return false
}


//returns the edges that were not used in BFS
func NotDiscovered(list []*Edge, discovered []*Vertex) []*Edge {

	var res []*Edge

	for _, e := range list {

		//if edge not used one of its endpoints has not been discovered
		if !Contains(discovered, e.v1) || !Contains(discovered, e.v2) {
			res = append(res, e)
		}
	}

	return res
}

//return neighbors reachable with a given set of edges
func GetNeighbors(v *Vertex, edges []*Edge) []*Vertex {

	var neighbors []*Vertex

	for _, e := range edges {

		if e.v1 == v {
			neighbors = append(neighbors, e.v2)
		} else if e.v2 == v {	
			neighbors = append(neighbors, e.v1)	
		}
	}

	return neighbors

}

//returns vertices corresponding to the set of edges
func GetVertices(edges []*Edge) []*Vertex {

	var res []*Vertex

	for _, e := range edges {

		if !Contains(res, e.v1) {
			res = append(res, e.v1)
		}
		if !Contains(res, e.v2) {
			res = append(res, e.v2)
		}
	}

	return res
}

//check if there is a cycle in the list of edges
//BFS to find cycles
func NoCycles(edges []*Edge) bool {

	if len(edges) < 2 {
		return true
	}

	vertices := GetVertices(edges)

	//queue of vertices we want to run bfs from
	queue := list.New()

	var discovered []*Vertex

	//root of BFS
	root := edges[0].v1
	discovered = append(discovered, root)

	//undiscovered edges remaining
	var remaining []*Edge

	queue.PushBack(root)

	for queue.Len() > 0 {

		//dequeue
		v := queue.Front()
		queue.Remove(v)

		remaining = NotDiscovered(edges, discovered)

		for _, v := range GetNeighbors(v.Value.(*Vertex), remaining) {

			//cycle found
			if Contains(discovered, v){
				return false
			}

			//mark as discovered and enqueue
			queue.PushBack(v)
			discovered = append(discovered, root)

		}

		//if not all vertices discovered
		//launch with remaining edges
		if len(discovered) < len(vertices){

			remaining = NotDiscovered(edges, discovered)

			return NoCycles(remaining)
		}

	}

	return true
}


//Returns a MST for an undirected weighted graph
func MinimumSpanningTree(edges []*Edge, vertices []*Vertex) ([]*Edge){

	var res []*Edge
	
	order := func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	}

	//first sort the edges in increasing weight order
	sort.Slice(edges, order)

	//since MST is a matroid, greedy finds the right solution
	for _, e := range edges {

		tmp := append(res, e)

		if NoCycles(tmp){
			res = append(res, e)
		}

	}

	return res
}