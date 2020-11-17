import sort.Slice

//Edges have a cost and 2 Vertices
struct Edges{
	v1,v2 *Vertex
	weight uint
}

//vertices have a list of edges
struct Vertex{
	neighbors []*Vertex
}

//returns true if the vertex is contained in the list
func contains(list []*Vertex, v *Vertex) bool {

	for _, vertex := range list{
		if vertex == v {
			return true
		}
	}

	return false
}


//returns the edges that were not used in BFS
func notContained(list []*Edge, discovered []*Vertex) []*Edge {

	var res []*Edge

	for _, e := range list {

		//if edge not used
		if !contains(e.v1, discovered) && !contains(e.v2, discovered) {
			res = append(res, e)
		}
	}

	return res
}

//return neighbors to a given vertex
func getNeighbors(v *Vertex, edges []*Edge) []*Vertex {

	neighbors []*Vertex

	for _, e := range edges {
		if e.v1 == v {
			neighbors = append(neighbors, e.v2)
		}
		else {
			if e.v2 == v {
				neighbors = append(neighbors, e.v1)
			}
		}
	}

	return neighbors

}

//returns vertices corresponding to the set of edges
func getVertices(edges []*Edge) []*Vertex {

	var res []*Vertex

	for _, e := range edges {

		res := append(res, e.v1)
		res := append(res, e.v2)
	}

	return res
}

//check if there is a cycle in the list of edges
//BFS to find cycles
func noCycles(edges []*Edges) bool {

	if len(edges) < 1{
		return true
	}

	vertices := getVertices(edges)

	//queue of vertices we want to run bfs from
	queue := list.New()

	discovered := []*Vertex

	//root of BFS
	root := edges[0].v1
	discovered = append(discovered, root)

	queue.PushBack(root)

	for len(queue) > 0 {

		//dequeue
		v := queue.Front()
		queue.Remove(v)

		for _, v := range getNeighbors(v) {

			//cycle found
			if contains(discovered, v){
				return false
			}

			//mark as discovered and enqueue
			queue.PushBack(v)
			discovered = append(discovered, root)

		}

		//if not all vertices discovered
		//launch with remaining edges
		if len(discovered) < len(vertices){

			remaining := notContained(edges, discovered)

			return noCycles(remaining)
		}

	}

	return true
}


//Returns a MST for an undirected weighted graph
func MinimumSpanningTree(edges []*Edges, vertices []*Vertex) ([]*Edges){

	var res []*Edges

	//first sort the edges in increasing weight order
	sort.Slice(edges, 
		func(i, j int) bool {
			edges[i].weight < edges[j].weight
		}
	)

	//since MST is a matroid, greedy finds the right solution
	for _, e := range edges {

		tmp := append(res, e)

		if noCycles(tmp){
			res := append(res, e)
		}

	}

}