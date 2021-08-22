package rebis

import (
	"log"
	"os"
)

/*
	Logger for realization printf
*/
type Logger interface {
	Printf(format string, v ...interface{})
}

/*
	this is a safeguard, breaking on compile time in case
	log.Logger does not adhere to our Logger interface.
	see https://golang.org/doc/faq#guarantee_satisfies_interface
*/
var _ Logger = &log.Logger{}

/*
	DefaultLogger returns a `Logger` implementation
	backed by stdlib's log
*/
func DefaultLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}

/*
	ChangeLogger override the logger specified in the
	cahce structure if it matches the Logger interface
*/
func (c *cache) ChangeLogger(custom Logger) {
	if custom != nil {
		c.logger = custom
	}
}

func (c *cache) logIf(format string, v ...interface{}) {
	if c.logAll {
		c.logger.Printf(format, v...)
	}
}
