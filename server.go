package main

import (
	"fmt"
	"net/http"
)

//PlayerStore stores score infor about players
type PlayerStore interface {
	GetPlayerScore(name string) int
}

//PlayerServer is an HTTP interface for player information
type PlayerServer struct {
	store PlayerStore
}

//Server currently only returns 10 or 20
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// func GetPlayerScore(name string) int {
// 	var value int
// 	if name == "Pepper" {
// 		value = 20
// 		} else {
// 		value = 10
// 	}
// 	return value
// }
