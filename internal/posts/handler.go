package posts

import (
	cryptoRand "crypto/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/oklog/ulid/v2"
)

type Handler struct {
	Repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{Repo: repo}
}

func (h *Handler) Create(c *gin.Context) {
	var in CreateInput
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}
	// validações mínimas
	if in.Title == "" || in.Slug == "" || in.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, slug and content are required"})
		return
	}
	// default para category e formato
	if in.Type == "" {
		in.Type = "other"
	}
	switch in.ContentFormat {
	case "html", "markdown", "plaintext":
	default:
		in.ContentFormat = "plaintext"
	}

	// normalizações
	slug := NormalizeSlug(in.Slug)
	tags := dedupLower(in.Tags)
	creators := dedupKeepCase(in.Creators) // não força lower aqui (nomes próprios)

	// sanitização se for HTML
	content := in.Content
	if in.ContentFormat == "html" {
		p := bluemonday.UGCPolicy()
		// se quiser permitir iframes de youtube/vimeo, você pode expandir a policy aqui
		content = p.Sanitize(content)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	pst := &Post{
		ID:            newULID(),
		Type:          in.Type,
		Title:         in.Title,
		Slug:          slug,
		Content:       content,
		ContentFormat: in.ContentFormat,
		Excerpt:       in.Excerpt,
		CoverImage:    in.CoverImage,
		Tags:          tags,
		Rating:        in.Rating,
		ReleaseDate:   in.ReleaseDate,
		Creators:      creators,
		SourceURL:     in.SourceURL,
		Meta:          in.Meta,

		Status:    StatusDraft,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.Repo.Create(c.Request.Context(), pst); err != nil {
		switch err {
		case ErrConflict:
			c.JSON(http.StatusConflict, gin.H{"error": "conflict"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, pst)
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}
	p, err := h.Repo.GetByID(c.Request.Context(), id)
	if err != nil {
		switch err {
		case ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *Handler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug required"})
		return
	}
	p, err := h.Repo.GetBySlug(c.Request.Context(), slug)
	if err != nil {
		switch err {
		case ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, p)
}

func newULID() string {
	return ulid.MustNew(ulid.Timestamp(time.Now().UTC()), cryptoRand.Reader).String()
}

func dedupLower(xs []string) []string {
	seen := make(map[string]struct{}, len(xs))
	out := make([]string, 0, len(xs))
	for _, t := range xs {
		t = strings.ToLower(strings.TrimSpace(t))
		if t == "" {
			continue
		}
		if _, ok := seen[t]; ok {
			continue
		}
		seen[t] = struct{}{}
		out = append(out, t)
	}
	return out
}

func dedupKeepCase(xs []string) []string {
	seen := make(map[string]struct{}, len(xs))
	out := make([]string, 0, len(xs))
	for _, t := range xs {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		key := strings.ToLower(t)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, t)
	}
	return out
}
