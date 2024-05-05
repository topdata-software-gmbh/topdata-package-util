package util

import "strings"

func RenderString(s string, m map[string]string) string {
	for k, v := range m {
		s = strings.ReplaceAll(s, "{"+k+"}", v)
	}
	return s
}
