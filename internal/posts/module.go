package posts

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewHandler),
	fx.Provide(NewRepoMem), // trocar depois
	fx.Invoke(registerRoutes),
)

func registerRoutes(r *gin.Engine, h *Handler) {
	v1 := r.Group("/v1")
	posts := v1.Group("/posts")
	posts.POST("", h.Create)
	posts.GET("/:id", h.GetByID)
	posts.GET("/slug/:slug", h.GetBySlug)
}
