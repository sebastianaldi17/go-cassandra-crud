package main

import (
	"go-cassandra-crud/handler"
	"go-cassandra-crud/repo"
	"go-cassandra-crud/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "store"

	r := chi.NewRouter()

	repo := repo.New(cluster)
	usecase := usecase.New(*repo)
	handler := handler.New(*usecase)

	r.Get("/cartCounts", handler.FetchAll)
	r.Get("/cartCount/{userID}", handler.FetchOne)
	r.Post("/cartCount", handler.Insert)
	r.Delete("/cartCount/{userID}", handler.Delete)
	http.ListenAndServe(":3333", r)
}
