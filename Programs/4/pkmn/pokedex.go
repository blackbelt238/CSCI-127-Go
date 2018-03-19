package pkmn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Pokedex represents an index of Pokemon
type Pokedex struct {
	pokemon []*Pokemon
}

// CreatePokedex uses the file at the given location to create and return a new Pokedex
func CreatePokedex(fname string) (*Pokedex, error) {
	// open the file containing the Pokedex
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pokedex: %v", err)
	}
	defer file.Close()

	// start with an empty dex
	dex := &Pokedex{pokemon: make([]*Pokemon, 0, 0)}
	entry := bufio.NewScanner(file)

	// for each entry in the file, create and add the Pokemon
	for entry.Scan() {
		pkstr := entry.Text()
		pklist := strings.Split(pkstr, ",")

		types := make([]string, 1, 2)
		for i := 3; i < len(pklist); i++ {
			types = append(types, pklist[i])
		}

		// make a new Pokemon using the entry
		pkmn, err := CreatePokemon(pklist[0], pklist[1], pklist[2], types)
		if err != nil {
			return nil, fmt.Errorf("cannot add entry \"%s\" to Pokedex: %v", pkstr, err)
		}

		dex.pokemon = append(dex.pokemon, pkmn)
	}

	return dex, nil
}
