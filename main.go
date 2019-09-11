package main

import (
	"log"
	"net/http"g
)

//InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct{}

//GetPlayerStore gets player store information
func (i *InMemoryPlayerStore) GetPlayerStore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not start server on port 5000: %v", err)
	}
}
