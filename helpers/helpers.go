package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"unicode"

	"github.com/tildetown/tilde-fucksgiven-analytics/fucks"
)

// GetKindOfFucks takes in a command line argument for kind, and
// a slice of unique fucks. It returns output text to be displayed
// back to the caller. Otherwise, an error is returned.
func GetKindOfFucks(kind string, uf *[]string) (string, error) {
	var out string
	var err error

	switch kind {
	case "uppercase":
		out, err = UppercaseFucks(uf)
		if err != nil {
			return "", fmt.Errorf("failed to get Uppercase Fucks: %v", err.Error())
		}

		return out, nil
	case "lowercase":
		out, err = LowercaseFucks(uf)
		if err != nil {
			return "", fmt.Errorf("failed to get Lowercase Fucks: %v", err.Error())
		}

		return out, nil
	}

	return out, nil
}

// UppercaseFucks takes a slice of strings representing unique
// fucks, and returns a string showcasing a count of them, a count of
// uppercase fucks, and its percentage (uppercase / total unique).
// Otherwise, an error is returned.
func UppercaseFucks(uf *[]string) (string, error) {
	ucf, err := GetUpperCaseFucks(*uf)
	if err != nil {
		return "", fmt.Errorf("failed to get uppercase fucks: %v", err.Error())
	}

	lenUF := len(*uf)
	lenUCF := len(ucf)

	p := float32(lenUCF) / float32(lenUF)

	tf := fmt.Sprintf("# of unique fucks: %v; ", lenUF)
	cf := fmt.Sprintf("# of uppercase fucks: %v; ", lenUCF)
	pf := fmt.Sprintf("percentage: %v", p*100)

	return tf + cf + pf, nil
}

// LowercaseFucks takes a slice of strings representing unique
// fucks, and returns a string showcasing a count of them, a count of
// lowercase fucks, and its percentage (lowercase / total unique).
// Otherwise, an error is returned.
func LowercaseFucks(uf *[]string) (string, error) {
	lcf, err := GetLowerCaseFucks(*uf)
	if err != nil {
		return "", fmt.Errorf("failed to get lowercase fucks: %v", err.Error())
	}

	lenUF := len(*uf)
	lenLCF := len(lcf)

	p := float32(lenLCF) / float32(lenUF)

	tf := fmt.Sprintf("# of unique fucks: %v; ", lenUF)
	cf := fmt.Sprintf("# of lowercase fucks: %v; ", lenLCF)
	pf := fmt.Sprintf("percentage: %v", p*100)

	return tf + cf + pf, nil
}

// NumUppercaseRunes takes a slice of runes and returns
// an integer value for a count of it
func NumUppercaseRunes(runes *[]rune) *int {
	u := 0

	for _, r := range *runes {
		if unicode.IsUpper(r) {
			u++
		}
	}

	return &u
}

// NumLowercaseRunes takes a slice of runes and returns
// an integer value for a count of it
func NumLowercaseRunes(runes *[]rune) *int {
	u := 0

	for _, r := range *runes {
		if unicode.IsLower(r) {
			u++
		}
	}

	return &u
}

// GetUpperCaseFucks takes a slice of unique variations of strings
// with "fuck" in them, and returns a slice of only those strings
// that are uppercase, otherwise, an error is returned
func GetUpperCaseFucks(uf []string) ([]string, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return nil, err
	}

	var ufc []string

	for _, v := range uf {
		processedString := reg.ReplaceAllString(v, "")
		runes := []rune(processedString)
		lenR := len(runes)

		u := NumUppercaseRunes(&runes)
		if *u == lenR {
			ufc = append(ufc, v)
		}
	}

	return ufc, nil
}

// GetLowerCaseFucks takes a slice of unique variations of strings
// with "fuck" in them, and returns a slice of only those strings
// that are lowercase, otherwise, an error is returned
func GetLowerCaseFucks(uf []string) ([]string, error) {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return nil, err
	}

	var ufc []string

	for _, v := range uf {
		processedString := reg.ReplaceAllString(v, "")
		runes := []rune(processedString)
		lenR := len(runes)

		u := NumLowercaseRunes(&runes)
		if *u == lenR {
			ufc = append(ufc, v)
		}
	}

	return ufc, nil
}

// ParseFucks takes no arguments and utilizes the supplied fucks.json
// HTTP URL to parse its data into a Fucks struct type
func ParseFucks() (*fucks.Fucks, error) {
	archFucksGivenURL := "http://tilde.town/~archangelic/fucks.json"
	resp, err := http.Get(archFucksGivenURL)
	if err != nil {
		return nil, fmt.Errorf("failed to GET ~archangelic's fucks json: %v", err.Error())
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var f fucks.Fucks

	err = json.Unmarshal(bodyBytes, &f)
	if err != nil {
		return nil, fmt.Errorf("failed to put unique fucks in struct: %v", err.Error())
	}

	return &f, nil
}
