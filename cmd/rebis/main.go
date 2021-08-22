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
	// rebis.ConfigCreateDefault(confDef)
	rebisConfig, err := rebis.ConfigFrom(confDir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rebisCache, err := rebis.NewCache(rebisConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// start := time.Now()
	// for i := 0; i < 100; i++ {
	// 	rebisCache.Set(strconv.Itoa(i), i, 0)
	// }
	// end := time.Since(start)
	// time.Sleep(time.Second * 2)
	// fmt.Println(end)

	fmt.Println(rebisCache.Get("2"))
	// fmt.Println(rebisCache.GetWithExpiration("1"))
	// rebisCache.Replace("3", 300, time.Duration(time.Second*500))

	// fmt.Println(rebisCache.Decrement("1", 1000))
	// fmt.Println(rebisCache.GetWithExpiration("1"))

	// rebisCache.Set("config", &rebisConfig, time.Duration(time.Second*500))

	// rebisCache.BackupRecovery()
	// fmt.Println(rebisCache.Items())

	// err = rebisCache.BackupRecoveryFile("./backup/backup1629630459.json")
	// fmt.Println(err)

	// rebisCache.ItemCount()
	// rebisCache.Items()

	// time.Sleep(time.Second * 3)
	// fmt.Println(a)

	for true {

	}
}
