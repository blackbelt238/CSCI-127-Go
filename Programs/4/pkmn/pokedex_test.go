package pkmn

import (
	"strings"
	"testing"
)

// TestCreatePokedex ensures CreatePokedex is properly contacting files and populating Pokemon
func TestCreatePokedex(t *testing.T) {
	// test a valid file
	p, err := CreatePokedex("test.txt")
	if err != nil {
		t.Errorf("Making Pokedex from \"test.txt\": %v", err)
	}

	// ensure proper dex importing
	if p.pokemon[2].Name != "Lunala" {
		t.Errorf("created pokedex does not have the pokemon. found %s", p.pokemon[2].Name)
	}

	// test a nonexistent file
	_, err = CreatePokedex("404.txt")
	if !strings.Contains(err.Error(), "The system cannot find the file specified") {
		t.Errorf("file not found error expected, instead recieved:\n%v", err)
	}

	// test an invalid file
	_, err = CreatePokedex("test_err.txt")
	if !strings.Contains(err.Error(), "invalid syntax") {
		t.Errorf("syntax error expected, instead recieved:\n%v", err)
	}
}
