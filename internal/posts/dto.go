package posts

// Campos opcionais permitem abrir o leque para obras diferentes
type CreateInput struct {
	Type          string         `json:"type"` // "movie","comic","manga","escape_room","other"...
	Title         string         `json:"title"`
	Slug          string         `json:"slug"`
	Content       string         `json:"content"`
	ContentFormat string         `json:"contentFormat"` // "html" | "markdown" | "plaintext"
	Excerpt       string         `json:"excerpt,omitempty"`
	CoverImage    string         `json:"coverImage,omitempty"`
	Tags          []string       `json:"tags"`
	Rating        *float64       `json:"rating,omitempty"`
	ReleaseDate   *string        `json:"releaseDate,omitempty"` // RFC3339 ou "YYYY-MM"
	Creators      []string       `json:"creators,omitempty"`
	SourceURL     string         `json:"sourceUrl,omitempty"`
	Meta          map[string]any `json:"meta,omitempty"` // flexível
}
