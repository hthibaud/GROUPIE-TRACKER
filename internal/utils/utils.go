package utils

import (
	"strings"
	"unicode"
)

// CleanLocations takes the map from the API and replaces dashes and underscores by blanks,
// and uppers the first letter
func CleanLocations(rawRelations map[string][]string) map[string][]string {
	cleanMap := make(map[string][]string)

	for location, dates := range rawRelations {

		r := strings.NewReplacer("-", " ", "_", " ")
		spaced := r.Replace(location)

		words := strings.Fields(spaced)
		for i, word := range words {
			runes := []rune(word)
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
				words[i] = string(runes)
			}
		}

		cleanName := strings.Join(words, " ")
		cleanMap[cleanName] = dates
	}

	return cleanMap
}
