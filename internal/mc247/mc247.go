package mc247

import (
	"log/slog"

	"github.com/interrrp/mc247/internal/bot"
	"github.com/interrrp/mc247/internal/config"
)

type Mc247 struct {
	logger *slog.Logger
	cfg    *config.Config
	bot    bot.Bot
}

func New(logger *slog.Logger) (*Mc247, error) {
	cfg, err := config.LoadFromEnvironmentVariables()
	if err != nil {
		return nil, err
	}
	logger.Info("loaded config", "username", cfg.Username, "address", cfg.Address)

	bot := bot.New(logger, cfg)

	return &Mc247{logger, cfg, bot}, nil
}

func (m *Mc247) Run() error {
	return m.bot.Run()
}
