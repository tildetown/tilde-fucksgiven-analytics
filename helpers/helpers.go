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
