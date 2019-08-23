package main

import (
	"fmt"
	"net/http"
)

//PlayerServer currently only returns 20
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprintf(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	var value string
	if name == "Pepper" {
		value = "20"
		} else {
		value = "10"
	}
	return value
}