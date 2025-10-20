package posts

type CreateInput struct {
	Title   string   `json:"title"`
	Slug    string   `json:"slug"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
