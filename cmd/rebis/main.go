package main

import (
	"log"

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

	// rebisCache.Set("1", 10, time.Duration(time.Second*500))
	// rebisCache.Set("2", 20, time.Duration(time.Second*500))
	// rebisCache.Set("3", 30, time.Duration(time.Second*500))
	// rebisCache.SetDefault("4", 500)

	// rebisCache.Add("5", 50, time.Duration(time.Second*600))

	// fmt.Println(rebisCache.Get("2"))
	// fmt.Println(rebisCache.GetWithExpiration("1"))
	// rebisCache.Replace("3", 300, time.Duration(time.Second*500))

	// fmt.Println(rebisCache.Decrement("1", 1000))
	// fmt.Println(rebisCache.GetWithExpiration("1"))

	// rebisCache.Set("config", &rebisConfig, time.Duration(time.Second*500))

	// rebisCache.BackupRecovery()
	// fmt.Println(rebisCache.Items())

	rebisCache.BackupRecoveryFile("./backup/backup1629531679.json")


	for true {

	}
}
