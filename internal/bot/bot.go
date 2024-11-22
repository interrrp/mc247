package bot

import (
	"log/slog"
	"time"

	mcbot "github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/interrrp/mc247/internal/config"
)

type Bot struct {
	logger    *slog.Logger
	client    *mcbot.Client
	player    *basic.Player
	cfg       *config.Config
	connected bool
}

func New(logger *slog.Logger, cfg *config.Config) Bot {
	bot := Bot{logger: logger}

	bot.client = mcbot.NewClient()
	bot.client.Name = cfg.Username
	bot.client.Auth.Name = cfg.Username

	bot.player = basic.NewPlayer(bot.client, basic.DefaultSettings, basic.EventsListener{
		Death:      bot.handleDeath,
		Disconnect: bot.handleDisconnect,
	})

	bot.cfg = cfg

	return bot
}

func (b *Bot) Run() error {
	b.joinServer()
	b.startRejoinTask()

	for {
		if b.connected {
			b.handlePackets()
		}
	}
}

func (b *Bot) joinServer() {
	for err := b.client.JoinServer(b.cfg.Address); err != nil; {
		b.logger.Error("failed joining server, rejoining after 5 seconds", "err", err)
		time.Sleep(5 * time.Second)
	}
	b.logger.Info("joined server")
}

func (b *Bot) handlePackets() {
	if err := b.client.HandleGame(); err != nil {
		b.logger.Warn("failed handling packets", "err", err)
	}
}
