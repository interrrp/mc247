package bot

import (
	"log/slog"
	"time"
)

func (b *Bot) startRejoinTask() {
	ticker := time.NewTicker(time.Duration(b.cfg.RejoinIntervalMins) * time.Minute)
	go func() {
		for range ticker.C {
			b.connected = false
			if err := b.client.Close(); err != nil {
				slog.Error("failed closing connection", "err", err)
				time.Sleep(3 * time.Second)
				continue
			}
			b.joinServer()
			b.connected = true
		}
	}()
}
