package main

import (
	"fmt"
	"os"

	"github.com/tildetown/tilde-fucksgiven-analytics/helpers"
)

func main() {
	f, err := helpers.ParseFucks()
	if err != nil {
		fmt.Printf("failed to parse unique fucks: %v", err.Error())
	}

	uf := f.UniqueFucks

	kind := os.Args[1]

	output, err := helpers.GetKindOfFucks(kind, &uf)
	if err != nil {
		fmt.Printf("failed to get %v fucks: %v", kind, err.Error())
	}

	fmt.Printf(output)
}
