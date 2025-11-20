package posts

import (
	"context"
	"errors"
	"sync"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
)

type Post struct {
	ID string `json:"id"`
	// Categoria ampla da obra/conteúdo
	// exemplos: "movie","comic","manga","anime","series","game","escape_room","article","other"
	Type string `json:"type"`

	Title string `json:"title"`
	Slug  string `json:"slug"`

	// Conteúdo longo; pode conter HTML/Markdown dependendo de ContentFormat
	Content       string `json:"content"`
	ContentFormat string `json:"contentFormat"` // "html" | "markdown" | "plaintext"

	// Metadados úteis para card/listagem/SEO
	Excerpt    string   `json:"excerpt,omitempty"`
	CoverImage string   `json:"coverImage,omitempty"`
	Tags       []string `json:"tags,omitempty"`

	// Metadados opcionais de obra
	Rating      *float64       `json:"rating,omitempty"`      // 0..10 ou 0..5, você decide
	ReleaseDate *string        `json:"releaseDate,omitempty"` // RFC3339 ou "YYYY-MM"
	Creators    []string       `json:"creators,omitempty"`    // diretores, autores, etc.
	SourceURL   string         `json:"sourceUrl,omitempty"`   // link de referência
	Meta        map[string]any `json:"meta,omitempty"`        // flexível (ex.: duração, estúdio, cidade do escape…)

	Status      Status  `json:"status"`
	CreatedAt   string  `json:"createdAt"` // RFC3339
	UpdatedAt   string  `json:"updatedAt"` // RFC3339
	PublishedAt *string `json:"publishedAt,omitempty"`
}

type Repository interface {
	Create(ctx context.Context, p *Post) error
	GetByID(ctx context.Context, id string) (*Post, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
}

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
)

// ---------- Implementação em memória (para desenvolvimento/teste) ----------
type repoMem struct {
	mu    sync.RWMutex
	posts map[string]*Post
}

func NewRepoMem() Repository {
	return &repoMem{posts: make(map[string]*Post)}
}

func (r *repoMem) Create(ctx context.Context, p *Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// conflito por slug
	for _, existing := range r.posts {
		if existing.Slug == p.Slug {
			return ErrConflict
		}
	}
	r.posts[p.ID] = p
	return nil
}

func (r *repoMem) GetByID(ctx context.Context, id string) (*Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.posts[id]
	if !ok {
		return nil, ErrNotFound
	}
	return p, nil
}

func (r *repoMem) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.posts {
		if p.Slug == slug {
			return p, nil
		}
	}
	return nil, ErrNotFound
}
