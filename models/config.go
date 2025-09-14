package models

type Config struct {
	Next          int32
	Prev          int32
	CurrentOffset int32
	Limit         int32
	Pokedex       map[string]Pokemon
}
