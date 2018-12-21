package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"unicode"
)

// Fucks is a type for parsing a list of unique forms of the word
// "fuck" on tilde.town
type Fucks struct {
	UniqueFucks []string `json:"unique_fucks"`
}

func main() {
	f, err := parseFucks()
	if err != nil {
		fmt.Printf("failed to parse unique fucks: %v", err.Error())
	}

	uf := f.UniqueFucks
	numUF := len(uf)

	fmt.Printf("Number of Unique Fucks: %v\n", numUF)

	upperCaseFucks, err := getUpperCaseFucks(uf)
	if err != nil {
		fmt.Printf("failed to get upper case fucks: %v", err.Error())
	}

	lenUpperCaseFucks := len(*upperCaseFucks)
	fmt.Printf("Number of CAPSLOCK Fucks: %v\n", lenUpperCaseFucks)

	var p float32
	p = float32(lenUpperCaseFucks) / float32(numUF)
	fmt.Printf("Percentage of CAPSLOCK Fucks: %v", p*100)
}

func parseFucks() (*Fucks, error) {
	archFucksGivenURL := "http://tilde.town/~archangelic/fucks.json"
	resp, err := http.Get(archFucksGivenURL)
	if err != nil {
		return nil, fmt.Errorf("failed to GET ~archangelic's fucks json: %v", err.Error())
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var f Fucks

	err = json.Unmarshal(bodyBytes, &f)
	if err != nil {
		return nil, fmt.Errorf("failed to put unique fucks in struct: %v", err.Error())
	}

	return &f, nil
}

func getUpperCaseFucks(uf []string) (*[]string, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return nil, err
	}

	var ufc []string

	for _, v := range uf {
		processedString := reg.ReplaceAllString(v, "")
		runes := []rune(processedString)
		lenR := len(runes)
		u := 0

		for _, r := range runes {
			if unicode.IsUpper(r) {
				u++
			}
		}

		if u == lenR {
			ufc = append(ufc, v)
		}
	}

	return &ufc, nil
}
