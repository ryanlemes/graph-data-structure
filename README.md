# Graph data structure
Graph data structure implemented in Golang

## Folder structure

This project was created based on [Uncle's bob Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
and [bexcodec golang clean architecture](https://github.com/bxcodec/go-clean-arch).

Is structured like this:

```
root/
├── cmd/
│   └── cli/
│       └── main.go
├── domain/
│   ├── mocks/
│   │   ├── ArticleRepository.go
│   │   ├── AuthorRepository.go
│   │   └── ArticleUsecase.go
│   ├── edge.go
│   ├── graph.go
│   └── vertex.go 
└── graph/
    ├── repository/
    │   └── memory/
    │       ├── memory_graph.go
    │       └── memory_graph_test.go
    └── usecase/
        ├── graph_usecase.go
        └── graph_usecase_test.go
```