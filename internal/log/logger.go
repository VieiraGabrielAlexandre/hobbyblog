package log

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle) (*zap.Logger, error) {
	// Implementation of logger initialization
	l, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error { return l.Sync() },
	})

	return l, nil
}
