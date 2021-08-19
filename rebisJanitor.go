package rebis

import (
	"time"

	"go.uber.org/zap"
)

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.run(c)
}

func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func (j *janitor) run(c *cache) {
	ticker := time.NewTicker(j.Interval)
	c.logger.Info(
		"START JANITOR",
		zap.String("interval", j.Interval.String()),
	)
	for {
		select {
		case <-ticker.C:
			c.logger.Debug("DELETE EXPIRED")
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}
