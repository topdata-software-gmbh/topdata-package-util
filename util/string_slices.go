package util

import (
	"log"
	"regexp"
)

// FilterStringSlice filters the given strings based on the given regex pattern and returns only those that match/match not the pattern.
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

func FilterStringSlicePositive(branches []string, regexPattern string) []string {
	return FilterStringSlice(branches, regexPattern, true)
}

func FilterStringSliceNegative(names []string, regex string) []string {
	return FilterStringSlice(names, regex, false)
}

// StringSliceContains checks if a string slice contains a string
func StringSliceContains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
