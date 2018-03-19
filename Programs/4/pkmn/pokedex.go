package pkmn

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		types := make([]string, 0, 2)
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

// Count outputs the number of Pokemon present in the Pokedex
func (p *Pokedex) Count() {
	fmt.Printf("There are %d different Pokemon in the Pokedex", len(p.pokemon))
}

// CP outputs the total number of combat points across all Pokemon in the Pokedex
func (p *Pokedex) CP() {
	var pcp int // total cp in the Pokedex
	for _, pkmn := range p.pokemon {
		pcp += pkmn.CP
	}
	fmt.Printf("The total number of combat points for all Pokemon in the Pokedex is %d", pcp)
}

// LookupByName outputs whether a Pokemon of the given name is in the dex
func (p *Pokedex) LookupByName(name string) {
	for _, pkmn := range p.pokemon {
		if pkmn.Name == name {
			pkmn.Print()
			return
		}
	}
	fmt.Printf("No Pokemon named %s in the Pokedex\n", name)
}

// LookupByNum outputs whether a Pokemon of the given num is in the dex
func (p *Pokedex) LookupByNum(num int) {
	p.sort() // sort the dex in numerical order
	for _, pkmn := range p.pokemon {
		if pkmn.Num == num {
			pkmn.Print()
			return
		}
	}
	fmt.Printf("Pokemon number %d in not in the Pokedex\n", num)
}

// Print outputs the contents of the pokedex in
func (p *Pokedex) Print() {
	// sort the pokedex
	p.sort()

	fmt.Println("The Pokedex")
	fmt.Println("-----------")
	for _, pkmn := range p.pokemon {
		pkmn.Print()
	}
	fmt.Print("-----------\n\n")
}

func (p *Pokedex) sort() {
	sort.Slice(p.pokemon, func(i, j int) bool {
		return p.pokemon[i].Num < p.pokemon[j].Num
	})
}
