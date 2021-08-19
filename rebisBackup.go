package rebis

import (
	"os"
	"strconv"
	"time"
)

type backup struct {
	File     *os.File
	Path     string
	Interval time.Duration
	stop     chan bool
}

func runBackup(c *cache, bp string, bi time.Duration) {
	var err error
	b := &backup{
		Path:     bp + "/backup" + strconv.Itoa(int(time.Now().Unix())) + ".gob",
		Interval: bi,
		stop:     make(chan bool),
	}
	b.File, err = os.Create(b.Path)
	if err != nil {
		c.logger.Printf("can not open gob file %s", err.Error())
	}

	c.backup = b
	go b.run(c)
}

func stopBackup(c *Cache) {
	c.backup.stop <- true
}

func (b *backup) run(c *cache) {
	ticker := time.NewTicker(b.Interval)
	c.logger.Printf("start cache backup with file save in %s and interval %s", b.Path, b.Interval)
	for {
		select {
		case <-ticker.C:
			go c.BackupSave()
		case <-b.stop:
			c.logger.Printf("stop backup")
			ticker.Stop()
			return
		}
	}
}
