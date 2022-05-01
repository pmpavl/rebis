package rebis

import (
	"os"
	"strconv"
	"time"
)

type backup struct {
	Path     string
	Interval time.Duration
	stop     chan bool
}

func runBackup(c *cache, bp string, bi time.Duration) {
	b := &backup{
		Path:     bp + "/backup" + strconv.Itoa(int(time.Now().Unix())) + ".json",
		Interval: bi,
		stop:     make(chan bool),
	}

	if _, err := os.Create(b.Path); err != nil {
		c.logger.Printf("can not open json file %s", err.Error())
	}

	c.backup = b
	go b.run(c)
}

func stopBackup(c *Cache) {
	c.backup.stop <- true
}

func (b *backup) run(c *cache) {
	ticker := time.NewTicker(b.Interval)
	c.logIf("start cache backup with file save in %s and interval %s", b.Path, b.Interval)

	for {
		select {
		case <-ticker.C:
			go c.BackupSave()
		case <-b.stop:
			c.logIf("stop backup")
			ticker.Stop()

			return
		}
	}
}
