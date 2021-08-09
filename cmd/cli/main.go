package main

import (
	"fmt"

	"github.com/ryanlemes/graph-data-structure/domain"
	_graphRepo "github.com/ryanlemes/graph-data-structure/graph/repository/memory"
	_graphUCase "github.com/ryanlemes/graph-data-structure/graph/usecase"
)

func main() {

	graph := &domain.Graph{}
	gr := _graphRepo.NewMemoryGraphRepository(graph)
	gu := _graphUCase.NewGraphUsecase(gr)

	for i := 0; i < 5; i++ {
		gu.AddVertex(i)
	}

	gu.AddEdge(0, 1, 10)
	gu.AddEdge(0, 2, 3)
	gu.AddEdge(0, 3, 5)
	gu.AddEdge(2, 1, 7)
	gu.AddEdge(2, 4, 2)
	gu.AddEdge(2, 3, 4)

	fmt.Println(gr.GetEdge(2, 4))

	gu.Print()
}
