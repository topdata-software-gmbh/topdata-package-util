package util

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

// RenderString replaces placeholders in a string with values from a map
func RenderString(s string, m map[string]string) string {
	for k, v := range m {
		s = strings.ReplaceAll(s, "{"+k+"}", v)
	}
	return s
}

func MapToTable(keyValuePairs map[string]string) string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	for key, value := range keyValuePairs {
		t.AppendRow(table.Row{key, value})
	}

	return t.Render()
}

func FormatBool(store bool, trueText string, falseText string) string {
	if store {
		return trueText
	} else {
		return falseText
	}
}

func Truncate(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength] + "…"
	}
	return s
}
