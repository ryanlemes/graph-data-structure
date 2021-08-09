package domain

type Graph struct {
	Vertices []*Vertex
	Edge     []*Edge
}

// GraphUsecase represent the graph's usecases
type GraphUsecase interface {
	AddVertex(k int)
	AddEdge(from, to, value int)
	Print()
}

// GraphRepository represent the graph's repository contract
type GraphRepository interface {
	GetVertices() []*Vertex
	GetVertex(k int) *Vertex
	GetEdge(keyFrom, keyTo int) *Edge
	AddVertex(k int)
	AddEdge(fromVertex, toVertex *Vertex, value int)
	Contains(s []*Vertex, k int) bool
}
