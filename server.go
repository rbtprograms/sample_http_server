package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	getPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

//Server currently only returns 10 or 20
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.getPlayerScore(player)
	fmt.Fprint(w, score)
}

// func getPlayerScore(name string) int {
// 	var value int
// 	if name == "Pepper" {
// 		value = 20
// 		} else {
// 		value = 10
// 	}
// 	return value
// }