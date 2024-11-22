package bot

import (
	"time"

	"github.com/Tnze/go-mc/chat"
)

func (b *Bot) handleDeath() error {
	b.logger.Warn("died")
	if err := b.player.Respawn(); err != nil {
		b.logger.Error("failed respawning", "err", err)
		return err
	}
	b.logger.Info("respawned")
	return nil
}

func (b *Bot) handleDisconnect(reason chat.Message) error {
	b.logger.Warn("disconnected, rejoining after 5 seconds", "reason", reason)
	time.Sleep(5 * time.Second)
	b.joinServer()
	return nil
}
