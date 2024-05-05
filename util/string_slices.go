package util

import (
	"log"
	"regexp"
)

func FilterStringSlice(strings []string, regexPattern string, keepMatches bool) []string {
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

// FilterStringSlicePositive filters the given branches and returns only those that are either "server" or start with "release-".
// It takes a slice of strings representing the branch names as input and returns a slice of strings containing the filtered branch names.
func FilterStringSlicePositive(branches []string, regexPattern string) []string {
	return FilterStringSlice(branches, regexPattern, true)
}

func FilterStringSliceNegative(names []string, regex string) []string {
	return FilterStringSlice(names, regex, false)
}

// contains checks if a string slice contains a string
func StringSliceContains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
