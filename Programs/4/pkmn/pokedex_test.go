package pkmn

import (
	"strings"
	"testing"
)

// TestCreatePokedex ensures CreatePokedex is properly contacting files and populating Pokemon
func TestCreatePokedex(t *testing.T) {
	// test a valid file
	p, err := createTestDex()
	if err != nil {
		t.Errorf("Making Pokedex from \"test.txt\": %v", err)
	}

	// ensure proper dex importing
	if p.pokemon[2].Name != "Lunala" {
		t.Errorf("created pokedex does not have the pokemon. found %s", p.pokemon[2].Name)
	}
	if p.pokemon[2].Types[0] != "psychic" {
		t.Errorf("Types don't match. want psychic, found %s", p.pokemon[2].Types[0])
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

// TestCount checks if the number of pokemon in the dex is properly counted
func TestCount(t *testing.T) {
	p, _ := createTestDex()
	ct := p.Count()

	if ct != 4 {
		t.Errorf("number of total pokemon improperly counted: counted %d instead of 4", ct)
	}
}

// TestCP ensures proper calculation of total CP
func TestCP(t *testing.T) {
	p, _ := createTestDex()
	cp := p.CP()
	cpexp := 525 + 520 + 680 + 500

	if cp != cpexp {
		t.Errorf("number of total pokemon improperly counted: counted %d instead of %d", cp, cpexp)
	}
}

func TestLookupByName(t *testing.T) {

}

func TestLookupByNum(t *testing.T) {

}

func TestSort(t *testing.T) {

}

// enables the test file to easily be changed
func createTestDex() (*Pokedex, error) {
	return CreatePokedex("test.txt")
}
