package main

import (
	"log"
	"net/http"
	"sample/test_driven/server/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
