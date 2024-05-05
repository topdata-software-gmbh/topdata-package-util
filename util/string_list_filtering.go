package util

import (
	"log"
	"regexp"
)

func FilterStringArray(strings []string, regexPattern string, keepMatches bool) []string {
	filtered := make([]string, 0)
	for _, str := range strings {
		matched, err := regexp.MatchString(regexPattern, str)
		if err != nil {
			log.Printf("Error matching regex: %v", err)
			continue
		}
		if matched == keepMatches {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

// FilterStringArrayPositive filters the given branches and returns only those that are either "server" or start with "release-".
// It takes a slice of strings representing the branch names as input and returns a slice of strings containing the filtered branch names.
func FilterStringArrayPositive(branches []string, regexPattern string) []string {
	return FilterStringArray(branches, regexPattern, true)
}

func FilterStringArrayNegative(names []string, regex string) []string {
	return FilterStringArray(names, regex, false)
}
