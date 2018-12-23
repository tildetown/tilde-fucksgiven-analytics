package fucks

// Fucks is a type for parsing a list of unique forms of the word
// "fuck" on tilde.town
type Fucks struct {
	UniqueFucks []string `json:"unique_fucks"`
}
