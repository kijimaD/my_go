package main

import (
	"log"
	"net/http"
	"sample/test_driven/server/server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &server.PlayerServer{Store: &InMemoryPlayerStore{}}

	log.Fatal(http.ListenAndServe(":5000", server))
}
