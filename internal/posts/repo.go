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
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Slug        string   `json:"slug"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags,omitempty"`
	Status      Status   `json:"status"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
	PublishedAt *string  `json:"publishedAt,omitempty"`
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

// ------------------------------------------------------------------
// Implementação em memória (thread-safe, para testes locais)
// ------------------------------------------------------------------

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

	// verificar conflito por slug
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
