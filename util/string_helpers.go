package util

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

// RenderString replaces placeholders in a string with values from a map.
// Placeholders in the string should be in the format {key}.
// Parameters:
// - s: The string containing placeholders.
// - m: A map where keys correspond to placeholders in the string and values are the replacements.
// Returns: The string with placeholders replaced by their corresponding values.
func RenderString(s string, m map[string]string) string {
	for k, v := range m {
		s = strings.ReplaceAll(s, "{"+k+"}", v)
	}
	return s
}

// MapToTable converts a map of key-value pairs to a formatted table string.
// Parameters:
// - keyValuePairs: A map containing the key-value pairs to be displayed in the table.
// Returns: A string representing the formatted table.
func MapToTable(keyValuePairs map[string]string) string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	for key, value := range keyValuePairs {
		t.AppendRow(table.Row{key, value})
	}

	return t.Render()
}

// FormatBool returns one of two strings based on the boolean value provided.
// Parameters:
// - store: The boolean value to evaluate.
// - trueText: The string to return if the boolean value is true.
// - falseText: The string to return if the boolean value is false.
// Returns: trueText if store is true, otherwise falseText.
func FormatBool(store bool, trueText string, falseText string) string {
	if store {
		return trueText
	} else {
		return falseText
	}
}

// Truncate shortens a string to a specified maximum length, adding an ellipsis if truncation occurs.
// Parameters:
// - s: The string to be truncated.
// - maxLength: The maximum allowed length of the string.
// Returns: The truncated string with an ellipsis if it exceeds maxLength, otherwise the original string.
func Truncate(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength] + "…"
	}
	return s
}
