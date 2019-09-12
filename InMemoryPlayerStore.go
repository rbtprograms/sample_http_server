package main

//NewInMemoryPlayerStore initialize an empty player store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//InMemoryPlayerStore stores player scores in memory
type InMemoryPlayerStore struct{
	store map[string]int
}

//AddPlayer adds a new player to the player store
func (i *InMemoryPlayerStore) AddPlayer(name string) {
	i.store[name] = 0
}

//RecordWin adds a win to the players score
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//GetPlayerScore gets player store information
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}