package envs

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"go-hexagonal-template/internal/infrastructure/config"
)

func WithEnvs(ctx context.Context, cfg *config.AppConfig) context.Context {
	_ = godotenv.Load()
	err := envconfig.Process(ctx, cfg)
	if err != nil {
		return nil
	}
	ctx = context.WithValue(ctx, "envs", cfg)
	return ctx
}
