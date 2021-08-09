package usecase

import (
	"fmt"

	"github.com/ryanlemes/graph-data-structure/domain"
)

type graphUsecase struct {
	gRepository domain.GraphRepository
}

func NewGraphUsecase(g domain.GraphRepository) domain.GraphUsecase {
	return &graphUsecase{
		gRepository: g,
	}
}

func (g *graphUsecase) AddVertex(k int) {
	if g.gRepository.Contains(g.gRepository.GetVertices(), k) {
		err := fmt.Errorf("vertex %v already exists", k)
		fmt.Println(err.Error())
		return
	}

	g.gRepository.AddVertex(k)
}

func (g *graphUsecase) AddEdge(from, to, value int) {
	fromVertex := g.gRepository.GetVertex(from)
	toVertex := g.gRepository.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
		return
	}

	if g.gRepository.Contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("already exists edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
		return
	}
	g.gRepository.AddEdge(fromVertex, toVertex, value)
}

func (g *graphUsecase) Print() {
	for _, v := range g.gRepository.GetVertices() {
		fmt.Printf("%v: \n", v.Key)
		for _, a := range v.Adjacent {
			fmt.Printf("  (%v -> %v", v.Key, a.Key)

			if data := g.gRepository.GetEdge(v.Key, a.Key); data != nil {
				fmt.Printf(" weigh: %v", data.Value)
			}

			fmt.Printf(")\n")
		}
	}
}
