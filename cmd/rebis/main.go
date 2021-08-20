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

	rebisCache.Set("1", 10, time.Duration(time.Second*100))
	rebisCache.Set("2", 20, time.Duration(time.Second*0))
	rebisCache.Set("3", 30, time.Duration(time.Second*-1))
	rebisCache.SetDefault("4", 40)

	rebisCache.Add("5", 50, time.Duration(time.Second*-1))

	fmt.Println(rebisCache.Get("2"))
	fmt.Println(rebisCache.GetWithExpiration("1"))
	rebisCache.Replace("3", 300, time.Duration(time.Second*100))

	fmt.Println(rebisCache.Decrement("1", 10))
	fmt.Println(rebisCache.GetWithExpiration("1"))

	for true {

	}
}
