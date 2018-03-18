package pkmn

import (
	"fmt"
	"strconv"
)

// Pokemon represents a Pokemon
type Pokemon struct {
	Num   int      // number in the pokedex
	Name  string   // Pokemon name
	CP    int      // total number of combat points
	Types []string // all types
}

// CreatePokemon creates a new Pokemon using the given infomation
func CreatePokemon(numstr string, name string, cpstr string, types []string) (*Pokemon, error) {
	pkmn := &Pokemon{
		Name:  name,
		Types: types,
	}

	// parse the Pokemon's dex number
	num, err := strconv.ParseInt(numstr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pokemon: %v", err)
	}
	pkmn.Num = int(num)

	// parse the Pokemon's CP
	cp, err := strconv.ParseInt(cpstr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pokemon: %v", err)
	}
	pkmn.CP = int(cp)

	return pkmn, nil
}
