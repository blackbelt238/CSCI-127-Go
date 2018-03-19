package pkmn

import (
	"strings"
	"testing"
)

func TestCreatePokemon(t *testing.T) {
	// the desired Pokemon
	pkmn := poke(0)

	// test using correct input
	prec, err := CreatePokemon("792", "Lunala", "680", []string{"psychic", "ghost"})
	if err != nil {
		t.Errorf("Pokemon couldn't be created: %v", err)
	}

	issues := make([]string, 0, 4)

	if pkmn.Num != prec.Num {
		issues = append(issues, "num")
	}
	if pkmn.Name != prec.Name {
		issues = append(issues, "name")
	}
	if pkmn.CP != prec.CP {
		issues = append(issues, "cp")
	}

	if len(issues) > 0 {
		errstr := "Created Pokemon does not have the expected " + issues[0]
		for i := 1; i < len(issues); i++ {
			errstr += ", " + issues[1]
		}
		errstr += ".\n\texpected: %v\n\trecieved: %v\n"
		t.Errorf(errstr, pkmn, prec)
	}

	// test using incorrect num input
	prec, err = CreatePokemon("seven hundred ninety two", "Lunala", "680", []string{"psychic", "ghost"})
	if !strings.Contains(err.Error(), "invalid syntax") {
		t.Errorf("Invalid syntax error expected, instead recieved:\n%v", err)
	}

	// test using incorrect num input
	prec, err = CreatePokemon("792", "Lunala", "six hundred eighty", []string{"psychic", "ghost"})
	if !strings.Contains(err.Error(), "invalid syntax") {
		t.Errorf("Invalid syntax error expected, instead recieved:\n%v", err)
	}
}

func poke(num int) *Pokemon {
	if num == 0 {
		return &Pokemon{
			Num:   792,
			Name:  "Lunala",
			CP:    680,
			Types: []string{"psychic", "ghost"},
		}
	} else if num == 1 {
		return &Pokemon{
			Num:   681,
			Name:  "Aegislash",
			CP:    520,
			Types: []string{"steel", "ghost"},
		}
	}
	return nil
}
