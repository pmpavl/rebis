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
	rebisConfig, err := rebis.ConfigFrom(confDir)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(rebisConfig)

	rebisCache, err := rebis.NewCache(rebisConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// rebisCache.Set("0", []int{1, 2}, 0)
	// rebisCache.Set("-1", 123, -1)
	// rebisCache.Set("5s", 123, time.Duration(time.Second*5))
	// rebisCache.SetDefault("default", 100)

	// fmt.Println(rebisCache.Get("0"))
	// fmt.Println(rebisCache.Get("1"))
	// fmt.Println(rebisCache.Get("-1"))
	// fmt.Println(rebisCache.Get("5s"))
	// fmt.Println(rebisCache.GetWithExpiration("5s"))
	// fmt.Println(rebisCache.ItemCount())
	// fmt.Println(rebisCache.Items())
	// fmt.Println(rebisCache.Replace("-1", 55555, time.Duration(time.Second*5)))
	// fmt.Println(rebisCache.Get("default"))
	// rebisCache.Increment("default", 100)
	// fmt.Println(rebisCache.Get("default"))

	// rebisCache.SaveFile("user.txt")

	rebisCache.LoadFile("user.txt")
	fmt.Println(rebisCache.Items())
	for true {
	}
}
