package graph

//A single entity in the graph
//Nodes must be comparable
type Node interface {
   //Get the edges for the node. This operation may be expensive
   Edges() []Edge
   Value() interface{}
}

//Generic edge with no direction or weight
type Edge interface {
   Nodes() [2]Node //TODO(chaten): If Nodes return a slice, this could technically be a hypergraph. Does that break things?
}

//An edge with a direction
type DirectedEdge interface {
   Edge
   Source() uint8
}

//An edge that has a cost to traverse it
type WeightedEdge interface {
   Edge
   Cost() float64
}

//Same interface as go/ast.Visitor
type Visitor interface {
   Visit(node Node) (w Visitor)
}
