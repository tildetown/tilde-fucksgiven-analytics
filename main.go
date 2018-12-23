package main

import (
	"fmt"
	
	"github.com/tildetown/tilde-fucksgiven-analytics/helpers"
)

func main() {
	f, err := helpers.ParseFucks()
	if err != nil {
		fmt.Printf("failed to parse unique fucks: %v", err.Error())
	}

	uf := f.UniqueFucks
	numUF := len(uf)

	fmt.Printf("Number of Unique Fucks: %v\n", numUF)

	upperCaseFucks, err := helpers.GetUpperCaseFucks(uf)
	if err != nil {
		fmt.Printf("failed to get upper case fucks: %v", err.Error())
	}

	lenUpperCaseFucks := len(upperCaseFucks)
	fmt.Printf("Number of CAPSLOCK Fucks: %v\n", lenUpperCaseFucks)

	var p float32
	p = float32(lenUpperCaseFucks) / float32(numUF)
	fmt.Printf("Percentage of CAPSLOCK Fucks: %v", p*100)
}
