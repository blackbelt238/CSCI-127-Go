package main

import (
	pkmn "CSCI-127-Go/Programs/4/pkmn"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	dex, _ := pkmn.CreatePokedex("pokedex.txt")

	scanner := bufio.NewScanner(os.Stdin)
	printMenu()
	fmt.Print("Enter a menu option: ")
	for scanner.Scan() {

		input = scanner.Text()

		if input == "1" {
			dex.Print()
		} else if input == "2" {
			fmt.Print("Enter a Pokemon name: ")
			scanner.Scan()
			dex.LookupByName(scanner.Text())
		} else if input == "3" {
			fmt.Print("Enter a Pokemon number: ")
			scanner.Scan()
			num, err := strconv.ParseInt(scanner.Text(), 10, 0)
			if err != nil {
				fmt.Println("Invalid input. A number must be entered.")
			} else {
				dex.LookupByNum(int(num))
			}
		} else if input == "4" {
			dex.Count()
		} else if input == "5" {
			dex.CP()
		} else {
			break
		}

		printMenu()
		fmt.Print("Enter a menu option: ")
	}
	fmt.Println("Thank you. Goodbye!")
}

func printMenu() {
	fmt.Println()
	fmt.Println("1. Print Pokedex")
	fmt.Println("2. Lookup Pokemon by Name")
	fmt.Println("3. Lookup Pokemon by Number")
	fmt.Println("4. Print Number of Pokemon")
	fmt.Println("5. Print total Combat Points of All Pokemon")
	fmt.Println("6. Quit")
}
