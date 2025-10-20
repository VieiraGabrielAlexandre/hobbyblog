package posts

import "strings"

func NormalizeSlug(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	// se quiser: remover acentos / colapsar múltiplos '-'
	return s
}
