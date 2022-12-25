package main

import (
	"log"
	"net/http"
	"sample/test_driven/server/server"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func main() {
	server := &server.PlayerServer{NewInMemoryPlayerStore()}

	log.Fatal(http.ListenAndServe(":5000", server))
}
