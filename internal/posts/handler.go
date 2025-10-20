package posts

import (
	"net/http"
	"strings"
	"time"

	cryptoRand "crypto/rand"
	"github.com/gin-gonic/gin"
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
	if in.Title == "" || in.Slug == "" || in.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title, slug and content are required"})
		return
	}
	now := time.Now().UTC().Format(time.RFC3339)
	p := &Post{
		ID:        newULID(), // implemente util simples ou troque depois
		Title:     in.Title,
		Slug:      NormalizeSlug(in.Slug),
		Content:   in.Content,
		Tags:      dedupLower(in.Tags), // opcional
		Status:    StatusDraft,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := h.Repo.Create(c.Request.Context(), p); err != nil {
		switch err {
		case ErrConflict:
			c.JSON(http.StatusConflict, gin.H{"error": "conflict"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, p)
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
