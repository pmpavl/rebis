package main

import (
	"fmt"
	"log"

	"github.com/pmpavl/rebis"
)

const (
	confDir = "rebisConfig.yaml"
	confDef = "rebisDefaultConfig.yaml"
)

func main() {
	config, err := rebis.ConfigFrom(confDir)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%+v\n", config)

	cache, err := rebis.NewCache(config)
	cache.Set("foo", "bar", 0)

	fmt.Println(cache.Get("foo"))
}
