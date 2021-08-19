package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pmpavl/rebis"
)

const (
	confDir = "rebisConfig.yaml"
	confDef = "rebisDefaultConfig.yaml"
)

func main() {
	rebisConfig, err := rebis.ConfigFrom(confDir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rebisCache, err := rebis.NewCache(rebisConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	time.Sleep(1 * time.Second)


	rebisCache.LoadFile("123.gob")
	fmt.Println(rebisCache.Items())
	for true {

	}
}
