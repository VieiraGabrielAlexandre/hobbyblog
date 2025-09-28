package app

import (
	"github.com/VieiraGabrielAlexandre/hobbyblog/internal/config"
	"github.com/VieiraGabrielAlexandre/hobbyblog/internal/health"
	apphttp "github.com/VieiraGabrielAlexandre/hobbyblog/internal/http"
	"github.com/VieiraGabrielAlexandre/hobbyblog/internal/log"
	"github.com/VieiraGabrielAlexandre/hobbyblog/internal/server"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.New,
			log.NewLogger,
			apphttp.NewEngine,
		),
		health.Module,
		fx.Invoke(server.StartHttpServer),
	)
}
