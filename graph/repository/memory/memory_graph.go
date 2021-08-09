package memory

import (
	"github.com/ryanlemes/graph-data-structure/domain"
)

type graphRepository struct {
	graph *domain.Graph
}

func NewMemoryGraphRepository(g *domain.Graph) domain.GraphRepository {
	return &graphRepository{
		graph: g,
	}
}

func addAdjacent(fromVertex, toVertex *domain.Vertex) {
	fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
}

func newEdge(fromVertex, toVertex *domain.Vertex, value int) *domain.Edge {
	return &domain.Edge{
		From:  fromVertex,
		To:    toVertex,
		Value: value,
	}
}

func (g *graphRepository) GetVertices() []*domain.Vertex {
	return g.graph.Vertices
}

func (g *graphRepository) GetVertex(k int) *domain.Vertex {
	for i, v := range g.graph.Vertices {
		if v.Key == k {
			return g.graph.Vertices[i]
		}
	}
	return nil
}

func (g *graphRepository) GetEdge(keyFrom, keyTo int) *domain.Edge {
	for i, v := range g.graph.Edge {
		if v.From.Key == keyFrom && v.To.Key == keyTo {
			return g.graph.Edge[i]
		}
	}
	return nil
}

func (g *graphRepository) Contains(s []*domain.Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

func (g *graphRepository) AddEdge(fromVertex, toVertex *domain.Vertex, value int) {
	addAdjacent(fromVertex, toVertex)
	g.graph.Edge = append(g.graph.Edge, newEdge(fromVertex, toVertex, value))
}

func (g *graphRepository) AddVertex(k int) {
	g.graph.Vertices = append(g.graph.Vertices, &domain.Vertex{Key: k})
}
