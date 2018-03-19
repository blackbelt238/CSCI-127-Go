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
	p := &Pokemon{
		Name:  name,
		Types: types,
	}

	// parse the Pokemon's dex number
	num, err := strconv.ParseInt(numstr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pokemon: %v", err)
	}
	p.Num = int(num)

	// parse the Pokemon's CP
	cp, err := strconv.ParseInt(cpstr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pokemon: %v", err)
	}
	p.CP = int(cp)

	return p, nil
}

// Print allows the pokemon to be printed
func (p *Pokemon) Print() {
	fmt.Printf("Number: %d, Name: %s, CP: %d, type: %s ", p.Num, p.Name, p.CP, p.Types[0])
	if len(p.Types) > 1 {
		fmt.Print("and ", p.Types[1])
	}
	fmt.Println()
}
